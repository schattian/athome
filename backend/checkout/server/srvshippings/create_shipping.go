package srvshippings

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/ent/shipping"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/backend/checkout/server/srvpurchases"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/currency"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateShipping(ctx context.Context, in *pbcheckout.CreateShippingRequest) (*pbcheckout.CreateShippingResponse, error) {
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
	err = srvpurchases.MustPrevState(ctx, db, o, sm.PurchaseShippingMethodSelected, uid)
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
	return s.createShipping(ctx, db, in, svcs, cals, o)
}

func (s *Server) createShipping(
	ctx context.Context,
	db *sqlx.DB,
	in *pbcheckout.CreateShippingRequest,
	svcs pbservices.ViewerClient,
	cals pbservices.CalendarsClient,
	p *purchase.Purchase,
) (*pbcheckout.CreateShippingResponse, error) {
	eventResp, err := cals.CreateShippingEvent(ctx, &pbservices.CreateShippingEventRequest{
		AccessToken: in.GetAccessToken(),
		ServiceId:   in.GetShipping().GetShippingMethodId(),
		Dow:         in.GetShipping().GetDow(),
		End:         in.GetShipping().GetTime(),
	})
	if err != nil {
		return nil, err
	}
	svcResp, err := svcs.RetrieveService(ctx, &pbservices.RetrieveServiceRequest{ServiceId: in.GetShipping().GetShippingMethodId()})
	if err != nil {
		return nil, err
	}
	ppkm, err := shipping.CalculateShippingPricePerKilometer(ctx, db, svcResp.GetUserId(), svcResp.GetPrice())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "CalculateShippingPricePerKilometer: %v", err)
	}
	price := ppkm.Float64() * p.DistanceInKilometers
	duration := pbutil.DiffTimeOfDay(eventResp.GetEvent().GetStart(), eventResp.GetEvent().GetEnd())
	ship := p.NewShipping(
		ctx, db,
		eventResp.GetEventId(),
		svcResp.GetUserId(),
		in.Shipping.GetShippingMethodId(),
		currency.ToARS(price),
		uint64(duration),
	)
	err = storeql.InsertIntoDB(ctx, db, ship)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}

	sc, err := shipping.NewShippingStateChange(ctx, ship.Id, sm.ShippingCreated)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "NewShippingStateChange")
	}
	err = storeql.InsertIntoDB(ctx, db, sc)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "sc InsertIntoDB")
	}

	p.ShippingId = ship.Id
	err = storeql.UpdateIntoDB(ctx, db, p)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.UpdateIntoDB: %v", err)
	}
	return &pbcheckout.CreateShippingResponse{
		ShippingId: ship.Id,
		Shipping:   ship.ToPb(),
	}, nil
}
