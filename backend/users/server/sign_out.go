package server

import (
	"context"
	"time"

	"github.com/athomecomar/athome/backend/users/pb/pbauth"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Proxies request to auth svc
func (s *Server) SignOut(ctx context.Context, in *pbuser.SignOutRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(userconf.GetAUTH_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, userconf.GetAUTH_ADDR())
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	return s.signOut(ctx, conn, in)
}

func (s *Server) signOut(ctx context.Context, conn *grpc.ClientConn, in *pbuser.SignOutRequest) (*emptypb.Empty, error) {
	c := pbauth.NewAuthClient(conn)
	authResponse, err := c.DeleteAuthentication(ctx, &pbauth.DeleteAuthenticationRequest{AccessToken: in.GetAccessToken()})
	if err != nil {
		return nil, err
	}
	return authResponse, nil
}
