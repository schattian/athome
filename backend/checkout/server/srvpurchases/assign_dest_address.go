package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AssignDestAddress(ctx context.Context, in *pbcheckout.AssignDestAddressRequest) (*emptypb.Empty, error) {
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

	addrs, addrsCloser, err := pbutil.ConnAddresses(ctx)
	if err != nil {
		return nil, err
	}
	defer addrsCloser()

	return s.assignAddress(ctx, db, in, addrs, o)
}

func (s *Server) assignAddress(
	ctx context.Context,
	db *sqlx.DB,
	in *pbcheckout.AssignDestAddressRequest,
	addrs pbaddress.AddressesClient,
	o *order.Purchase,
) (*emptypb.Empty, error) {
	resp, err := addrs.RetrieveAddress(ctx, &pbaddress.RetrieveAddressRequest{AddressId: in.GetDestAddressId()})
	if err != nil {
		return nil, err
	}
	if resp.GetUserId() != o.UserId {
		return nil, status.Errorf(xerrors.PermissionDenied, "the address' user id mismatch order's user id")
	}
	o.DestAddressId = in.GetDestAddressId()
	err = storeql.UpdateIntoDB(ctx, db, o)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.UpdateIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}
