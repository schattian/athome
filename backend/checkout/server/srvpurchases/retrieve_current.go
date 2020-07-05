package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveCurrent(ctx context.Context, in *pbcheckout.RetrieveCurrentRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
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

	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsCloser()
	return s.retrievePurchase(ctx, db, prods, o)
}

func (s *Server) retrievePurchase(
	ctx context.Context,
	db *sqlx.DB,
	prods pbproducts.ViewerClient,
	order *purchase.Purchase,
) (*pbcheckout.RetrievePurchaseResponse, error) {
	o, err := order.ToPbWrapped(ctx, db, prods)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "o.ToPbWrapped")
	}

	return &pbcheckout.RetrievePurchaseResponse{PurchaseId: order.Id, Purchase: o}, nil
}
