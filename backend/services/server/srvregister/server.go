package srvregister

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/ent/stage"
	"github.com/athomecomar/athome/backend/services/pb/pbauth"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

type Server struct{}

func mustStage(got stage.Stage, want stage.Stage) error {
	if got != want {
		return status.Errorf(xerrors.InvalidArgument, "got %v stage, while expecting %v", got, want)
	}
	return nil
}

func retrieveRegistryByUser(ctx context.Context, db *sqlx.DB, c pbauth.AuthClient, access string, authFn server.AuthFunc) (*ent.Registry, error) {
	userId, err := authFn(ctx, c, access)
	if err != nil {
		return nil, err
	}
	reg, err := ent.FindRegistryByUserId(ctx, db, userId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "registry for user with id %v wasnt found", userId)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindRegistryByUserId: %v", err)
	}
	return reg, nil
}
