package srvpayments

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/ent/payment"
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

func (s *Server) CreatePayment(ctx context.Context, in *pbcheckout.CreatePaymentRequest) (r *pbcheckout.CreatePaymentResponse, err error) {
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

	switch order.Class(in.GetPayment().Order.GetEntityTable()) {
	case order.Purchases:
		r, err = s.CreatePurchasePayment(ctx, db, in, uid)
	}
	return
}

func (s *Server) CreatePurchasePayment(ctx context.Context, db *sqlx.DB, in *pbcheckout.CreatePaymentRequest, userId uint64) (*pbcheckout.CreatePaymentResponse, error) {
	o, err := server.FindLatestPurchase(ctx, db, userId)
	if err != nil {
		return nil, err
	}
	err = server.MustPrevState(ctx, db, o, sm.PurchasePaid, userId)
	if err != nil {
		return nil, err
	}

	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsCloser()
	return s.createPurchasePayment(ctx, db, in, prods, o)
}

func (s *Server) createPurchasePayment(
	ctx context.Context,
	db *sqlx.DB,
	in *pbcheckout.CreatePaymentRequest,
	prods pbproducts.ViewerClient,
	p *purchase.Purchase,
) (*pbcheckout.CreatePaymentResponse, error) {
	py := payment.PaymentFromPb(in.GetPayment())
	py.UserId = p.UserId
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
		return nil, status.Errorf(xerrors.InvalidArgument, "py InsertIntoDB: %v", err)
	}

	sc, err := sm.NewStateChange(ctx, py.Id, sm.PaymentCreated, py)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "NewPaymentStateChange")
	}
	err = storeql.InsertIntoDB(ctx, db, sc)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "sc InsertIntoDB: %v", err)
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
