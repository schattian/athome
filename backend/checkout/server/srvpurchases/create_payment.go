package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/currency"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePayment(ctx context.Context, in *pbcheckout.CreatePaymentRequest) (*pbcheckout.CreatePaymentResponse, error) {
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
	err = mustPrevState(ctx, db, o, sm.PurchasePayment)
	if err != nil {
		return nil, err
	}
	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsCloser()
	return s.createPayment(ctx, db, in, prods, o)
}

func (s *Server) createPayment(
	ctx context.Context,
	db *sqlx.DB,
	in *pbcheckout.CreatePaymentRequest,
	prods pbproducts.ViewerClient,
	p *order.Purchase,
) (*pbcheckout.CreatePaymentResponse, error) {
	py := order.PaymentFromPb(in.GetPayment())
	amount, err := p.Amount(ctx, db, prods)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Amount: %v", err)
	}
	py.Amount = currency.ToARS(amount)
	_, err = py.Card(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "Card: %v", err)
	}
	err = storeql.InsertIntoDB(ctx, db, py)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "Card: %v", err)
	}
	pyPb, err := py.ToPb()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "ToPb: %v", err)
	}
	return &pbcheckout.CreatePaymentResponse{
		PaymentId: py.Id,
		Payment:   pyPb,
	}, nil
}
