package srvpayments

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/orderable"
	"github.com/athomecomar/athome/backend/checkout/ent/payment"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) Retrieve(ctx context.Context, in *pbcheckout.RetrievePaymentRequest) (*pbcheckout.Payment, error) {
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

	return s.retrieve(ctx, db, in, uid)
}

func (s *Server) retrieve(
	ctx context.Context,
	db *sqlx.DB,
	in *pbcheckout.RetrievePaymentRequest,
	userId uint64,
) (*pbcheckout.Payment, error) {
	py, err := payment.FindPayment(ctx, db, in.GetPaymentId(), userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindPayment: %v", err)
	}
	order, err := orderable.FromOrderable(ctx, db, py)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "py.Order: %v", err)
	}
	isAuthorized, err := order.CanView(ctx, db, userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "CanView: %v", err)
	}
	if !isAuthorized {
		return nil, status.Errorf(xerrors.PermissionDenied, "you are unauthorized to access this payment")
	}
	pyPb, err := py.ToPb()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "py.ToPb: %v", err)
	}
	return pyPb, nil
}
