package srvshippings

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/shipping"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AgreeShippingDispatch(ctx context.Context, in *pbcheckout.AgreeShippingDispatchRequest) (*emptypb.Empty, error) {
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
	sh, err := shipping.FindShipping(ctx, db, in.GetShippingId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindShipping: %v", err)
	}

	return s.agreeShippingDispatch(ctx, db, sh, uid)
}

func (s *Server) agreeShippingDispatch(
	ctx context.Context,
	db *sqlx.DB,
	sh *shipping.Shipping,
	uid uint64,
	id uint64,
) (*emptypb.Empty, error) {
	return nil, nil
	// return &pbcheckout.RetrieveShippingResponse{}, nil
}
