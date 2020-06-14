package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func FindUser(ctx context.Context, db *sqlx.DB, uid uint64) (*ent.User, error) {
	user, err := ent.FindUser(ctx, db, uid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "cant find user by id: %v", uid)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindUser: %v", err)
	}
	return user, nil
}
