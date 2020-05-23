package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/athomecomar/athome/backend/auth/authconf"
	"github.com/athomecomar/athome/backend/auth/pbauth"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func (s *Server) Authenticate(ctx context.Context, in *pbauth.AuthenticateRequest) (*pbauth.AuthenticateResponse, error) {
	claims, err := claimJwt(in.GetJwt(), authconf.GetAUTHENTICATE_JWT_SECRET)
	if err != nil {
		return nil, err
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "strconv.ParseUint: %v", err)
	}

	return &pbauth.AuthenticateResponse{UserId: userId}, nil
}
