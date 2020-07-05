package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) ConfirmPurchase(ctx context.Context, in *pbcheckout.ConfirmPurchaseRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
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

	prodsViewer, prodsViewerCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsViewerCloser()

	prodsManager, prodsManagerCloser, err := pbutil.ConnProductsManager(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsManagerCloser()

	cals, calsCloser, err := pbutil.ConnServicesCalendars(ctx)
	if err != nil {
		return nil, err
	}
	defer calsCloser()

	return s.confirmPurchase(ctx, db, in, cals, prodsViewer, prodsManager, uid)
}

func (s *Server) confirmPurchase(ctx context.Context, db *sqlx.DB, in *pbcheckout.ConfirmPurchaseRequest,
	cals pbservices.CalendarsClient,
	prodsViewer pbproducts.ViewerClient,
	prodsManager pbproducts.ManagerClient,
	userId uint64,
) (*pbcheckout.RetrievePurchaseResponse, error) {
	p, err := purchase.FindPurchase(ctx, db, in.GetPurchaseId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindPurchase: %v", err)
	}
	if p.MerchantId != userId {
		return nil, status.Error(xerrors.PermissionDenied, "You are not allowed to confirm this order")
	}
	_, err = prodsManager.ConfirmReserveStock(ctx, &pbproducts.ReserveStockRequest{AccessToken: in.GetAccessToken(), Order: pbutil.ToPbEntity(p)})
	if err != nil {
		return nil, err
	}
	ship, err := p.Shipping(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Shipping: %v", err)
	}
	_, err = cals.ConfirmEvent(ctx, &pbservices.ConfirmEventRequest{AccessToken: in.GetAccessToken(), EventId: ship.EventId})
	if err != nil {
		return nil, err
	}
	resp, err := s.changeState(ctx, db, sm.Next, prodsViewer, p, userId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
