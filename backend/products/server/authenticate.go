package server

import (
	"context"

	"github.com/athomecomar/athome/backend/products/pb/pbauth"
	"github.com/athomecomar/athome/backend/products/productconf"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func GetUserFromAccessToken(ctx context.Context, db *sqlx.DB, c pbauth.AuthClient, access string) (uint64, error) {
	resp, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access})
	if err != nil {
		return 0, err
	}

	return resp.GetUserId(), nil
}

func ConnAuth(ctx context.Context) (pbauth.AuthClient, func() error, error) {
	conn, err := grpc.Dial(productconf.GetAUTH_ADDR(), grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, productconf.GetAUTH_ADDR())
	}
	defer conn.Close()
	c := pbauth.NewAuthClient(conn)
	return c, conn.Close, nil
}