package server

import (
	"context"

	"github.com/athomecomar/athome/backend/address/ent"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAddress(ctx context.Context, in *pbaddress.CreateAddressRequest) (*pbaddress.CreateAddressResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()
	return s.createAddress(ctx, db, auth, in)
}

func (s *Server) createAddress(ctx context.Context, db *sqlx.DB, auth pbauth.AuthClient, in *pbaddress.CreateAddressRequest) (*pbaddress.CreateAddressResponse, error) {
	userId, err := pbutil.GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	addr := ent.AddressFromPb(in.GetBody())
	addr.UserId = userId

	err = storeql.InsertIntoDB(ctx, db, addr)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
	}
	resp := &pbaddress.CreateAddressResponse{Metadata: addr.ToPb(), AddressId: addr.Id}
	return resp, nil
}
