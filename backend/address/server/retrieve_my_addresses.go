package server

import (
	"context"

	"github.com/athomecomar/athome/backend/address/ent"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveMyAddresses(
	ctx context.Context,
	in *pbaddress.RetrieveMyAddressesRequest,
) (*pbaddress.RetrieveMyAddressesResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	auth, authCloser, err := ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	return s.retrieveMyAddresses(ctx, db, auth, in)
}
func (s *Server) retrieveMyAddresses(
	ctx context.Context,
	db *sqlx.DB,
	auth pbauth.AuthClient,
	in *pbaddress.RetrieveMyAddressesRequest,
) (*pbaddress.RetrieveMyAddressesResponse, error) {
	userId, err := GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	addrs, err := ent.AddressesByUser(ctx, db, userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "AddressesByUser: %v", err)
	}

	resp := &pbaddress.RetrieveMyAddressesResponse{}
	resp.Addresses = make(map[uint64]*pbaddress.Address)
	for _, addr := range addrs {
		resp.Addresses[addr.Id] = addr.ToPb()
	}
	return resp, nil
}
