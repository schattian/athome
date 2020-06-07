package server

import (
	"context"

	"github.com/athomecomar/athome/backend/images/imageconf"
	"github.com/athomecomar/athome/backend/images/pb/pbauth"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func GetUserFromAccessToken(ctx context.Context, c pbauth.AuthClient, access string) (uint64, error) {
	resp, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access})
	if err != nil {
		return 0, err
	}

	return resp.GetUserId(), nil
}

func ConnAuth(ctx context.Context) (pbauth.AuthClient, func() error, error) {
	conn, err := grpc.Dial(imageconf.GetAUTH_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, imageconf.GetAUTH_ADDR())
	}
	c := pbauth.NewAuthClient(conn)
	return c, conn.Close, nil
}