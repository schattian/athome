package signsrv

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"

	"github.com/athomecomar/athome/backend/users/internal/userjwt"

	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/xerrors"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) Sign(ctx context.Context, in *pbusers.SignRequest) (*pbusers.SignResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return s.sign(ctx, auth, in)
}

func (s *Server) sign(ctx context.Context, c pbauth.AuthClient, in *pbusers.SignRequest) (*pbusers.SignResponse, error) {
	userId, err := handleJwt(in.GetSignToken(), userconf.GetSIGN_JWT_SECRET)
	if err != nil {
		return nil, err
	}

	tokens, err := createAuthTokens(ctx, c, userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "createAuthToken: %v", err)
	}

	return &pbusers.SignResponse{
		AccessToken:       tokens.GetAccessToken(),
		RefreshToken:      tokens.GetRefreshToken(),
		AccessTokenExpNs:  tokens.GetAccessTokenExpNs(),
		RefreshTokenExpNs: tokens.GetRefreshTokenExpNs(),
	}, nil
}

func handleJwt(token string, secretFn func() string) (uint64, error) {
	claims, err := userjwt.ClaimJwt(token, secretFn)
	if err != nil {
		return 0, err
	}
	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return 0, status.Errorf(xerrors.InvalidArgument, "strconv.ParseUint: %v", err)
	}
	return userId, nil
}
func createAuthTokens(ctx context.Context, c pbauth.AuthClient, userId uint64) (*pbauth.CreateAuthenticationResponse, error) {
	signJwt, err := userjwt.CreateSignToken(userId)
	if err != nil {
		return nil, errors.Wrap(err, "createSignToken")
	}
	authResponse, err := c.CreateAuthentication(ctx, &pbauth.CreateAuthenticationRequest{SignToken: signJwt})
	if err != nil {
		return nil, errors.Wrap(err, "pbusers.Sign")
	}
	return authResponse, nil
}
