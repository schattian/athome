package server

import (
	"context"

	"github.com/athomecomar/athome/backend/services/pb/pbauth"
	"github.com/athomecomar/athome/backend/services/serviceconf"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func GetUserFromAccessToken(c pbauth.AuthClient, access string) AuthFunc {
	return func(ctx context.Context) (uint64, error) {
		resp, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access})
		if err != nil {
			return 0, err
		}

		return resp.GetUserId(), nil
	}
}

type AuthFunc func(ctx context.Context) (uint64, error)

func ConnAuth(ctx context.Context) (pbauth.AuthClient, func() error, error) {
	conn, err := grpc.Dial(serviceconf.GetAUTH_ADDR(), grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, serviceconf.GetAUTH_ADDR())
	}
	c := pbauth.NewAuthClient(conn)
	return c, conn.Close, nil
}
