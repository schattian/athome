package server

import (
	"context"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/pb/pbauth"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func GetUserFromAccessToken(ctx context.Context, db *sqlx.DB, conn *grpc.ClientConn, access string) (*ent.User, error) {
	c := pbauth.NewAuthClient(conn)
	userId, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access})
	if err != nil {
		return nil, err
	}

	user := &ent.User{}
	row := db.QueryRowxContext(ctx, `SELECT * FROM users WHERE id=$1`, userId.GetUserId())
	err = row.StructScan(user)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "oldUser row.StructScan: %v", err)
	}
	return user, nil
}