package srvregister

import (
	"context"

	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/jmoiron/sqlx"
)

func (s *Server) RetrieveRegistry(ctx context.Context, in *pbservices.RetrieveRegistryRequest) (*pbservices.Registry, error) {
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
	defer authCloser()

	return s.retrieveRegistry(ctx, db, server.GetUserFromAccessToken(auth, in.GetAccessToken()))
}

func (s *Server) retrieveRegistry(ctx context.Context, db *sqlx.DB, authFn server.AuthFunc) (*pbservices.Registry, error) {
	reg, err := retrieveRegistryByUser(ctx, db, authFn)
	if err != nil {
		return nil, err
	}
	return reg.ToPb(), nil
}
