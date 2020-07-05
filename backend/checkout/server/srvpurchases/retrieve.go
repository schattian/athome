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

func (s *Server) Retrieve(ctx context.Context, in *pbcheckout.RetrieveOrderRequest) (*pbcheckout.Purchase, error) {
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
	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsCloser()
	return s.retrieve(ctx, db, in, prods, uid)
}

func (s *Server) retrieve(
	ctx context.Context,
	db *sqlx.DB,
	in *pbcheckout.RetrieveOrderRequest,
	prods pbproducts.ViewerClient,
	uid uint64,
) (*pbcheckout.Purchase, error) {
	p, err := purchase.FindPurchase(ctx, db, in.GetOrderId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindPurchase: %v", err)
	}
	allowed, err := p.CanView(ctx, db, uid)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "CanView: %v", err)
	}
	if !allowed {
		return nil, status.Error(xerrors.PermissionDenied, "You are not allowed to view this purchase")
	}
	o, err := p.ToPbWrapped(ctx, db, prods)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "o.ToPbWrapped: %v", err)
	}

	return o, nil
}
