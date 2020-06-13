package signsrv

import (
	"context"
	"time"

	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbusers"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Proxies request to auth svc
func (s *Server) SignOut(ctx context.Context, in *pbusers.SignOutRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	auth, authCloser, err := pbconf.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	return s.signOut(ctx, auth, in)
}

func (s *Server) signOut(ctx context.Context, c pbauth.AuthClient, in *pbusers.SignOutRequest) (*emptypb.Empty, error) {
	authResponse, err := c.DeleteAuthentication(ctx, &pbauth.DeleteAuthenticationRequest{AccessToken: in.GetAccessToken()})
	if err != nil {
		return nil, err
	}
	return authResponse, nil
}
