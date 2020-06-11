package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func GetUserFromAccessToken(ctx context.Context, db *sqlx.DB, c pbauth.AuthClient, access string) (*ent.User, error) {
	resp, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access})
	if err != nil {
		return nil, err
	}

	user := &ent.User{}
	row := db.QueryRowxContext(ctx, `SELECT * FROM users WHERE id=$1`, resp.GetUserId())
	err = row.StructScan(user)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "oldUser row.StructScan: %v", err)
	}
	return user, nil
}

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
