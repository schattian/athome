package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/currency"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) assignShippingMethod(
	ctx context.Context,
	db *sqlx.DB,
	in *pbcheckout.AssignShippingMethodRequest,
	svcs pbservices.ViewerClient,
	cals pbservices.CalendarsClient,
	p *order.Purchase,
) (*emptypb.Empty, error) {
	eventResp, err := cals.CreateShippingEvent(ctx, &pbservices.CreateShippingEventRequest{
		AccessToken: in.GetAccessToken(),
		ServiceId:   in.GetServiceId(),
		Dow:         in.GetDow(),
		End:         in.GetTime(),
	})
	if err != nil {
		return nil, err
	}
	svcResp, err := svcs.RetrieveServiceDetail(ctx, &pbservices.RetrieveServiceDetailRequest{ServiceId: in.GetServiceId()})
	if err != nil {
		return nil, err
	}
	ppkm, err := order.CalculateShippingPricePerKilometer(ctx, db, svcResp.GetService().UserId, svcResp.GetService().GetPrice())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "CalculateShippingPricePerKilometer: %v", err)
	}
	price := ppkm.Float64() * p.DistanceInKilometers
	duration := pbutil.DiffTimeOfDay(eventResp.GetEvent().GetStart(), eventResp.GetEvent().GetEnd())
	ship := order.NewShipping(ctx, db, p, eventResp.GetEventId(), svcResp.Service.GetUserId(), currency.ToARS(price), uint64(duration))
	err = storeql.InsertIntoDB(ctx, db, ship)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}
	p.ShippingId = ship.GetId()
	err = storeql.UpdateIntoDB(ctx, db, p)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.UpdateIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) AssignShippingMethod(ctx context.Context, in *pbcheckout.AssignShippingMethodRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}

	uid, err := pbutil.GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	err = authCloser()
	if err != nil {
		return nil, err
	}

	o, err := server.FindLatestPurchase(ctx, db, uid)
	if err != nil {
		return nil, err
	}
	err = mustPrevState(ctx, db, o, sm.PurchaseShippingMethod)
	if err != nil {
		return nil, err
	}
	svcs, svcsCloser, err := pbutil.ConnServicesViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer svcsCloser()

	cals, calsCloser, err := pbutil.ConnServicesCalendars(ctx)
	if err != nil {
		return nil, err
	}
	defer calsCloser()
	return s.assignShippingMethod(ctx, db, in, svcs, cals, o)
}
