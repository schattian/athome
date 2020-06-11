package srvregister

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/ent/stage"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) First(ctx context.Context, in *pbservices.FirstRequest) (*pbservices.FirstResponse, error) {
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

	return s.first(ctx, db, server.GetUserFromAccessToken(auth, in.GetAccessToken()), in)
}

func (s *Server) first(ctx context.Context, db *sqlx.DB, authFn server.AuthFunc, in *pbservices.FirstRequest) (*pbservices.FirstResponse, error) {
	userId, err := authFn(ctx)
	if err != nil {
		return nil, err
	}
	reg, err := ent.FindRegistryByUserId(ctx, db, userId)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
		reg = ent.NewRegistry(userId)
	}
	if err != nil {
		return nil, err
	}
	err = mustStage(reg.Stage, stage.First)
	if err != nil {
		return nil, err
	}
	reg = applyFirstRequestToRegistry(in.GetBody(), reg)
	err = storeql.UpsertIntoDB(ctx, db, reg)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "UpsertIntoDB: %v", err)
	}

	return &pbservices.FirstResponse{RegistryId: reg.Id}, nil
}

func applyFirstRequestToRegistry(f *pbservices.FirstRequest_Body, r *ent.Registry) *ent.Registry {
	r.AddressId = f.GetAddressId()
	return r
}
