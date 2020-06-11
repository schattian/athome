package srvregister

import (
	"context"

	"github.com/athomecomar/athome/backend/services/pb/pbauth"
	"github.com/athomecomar/athome/backend/services/pb/pbservices"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/jmoiron/sqlx"
)

func (s *Server) RetrieveRegistry(ctx context.Context, in *pbservices.RetrieveRegistryRequest) (*pbservices.RegistryDetail, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	auth, authCloser, err := server.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	return s.retrieveRegistry(ctx, db, auth, server.GetUserFromAccessToken, in)
}

func (s *Server) retrieveRegistry(ctx context.Context, db *sqlx.DB, auth pbauth.AuthClient, authFn server.AuthFunc, in *pbservices.RetrieveRegistryRequest) (*pbservices.RegistryDetail, error) {
	reg, err := retrieveRegistryByUser(ctx, db, auth, in.GetAccessToken(), authFn)
	if err != nil {
		return nil, err
	}
	return &pbservices.RegistryDetail{
		RegistryId: reg.Id,
		Registry:   reg.ToPb(),
	}, nil
}
