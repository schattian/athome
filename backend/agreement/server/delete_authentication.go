package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/athomecomar/athome/backend/auth/authconf"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/xerrors"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteAuthentication(ctx context.Context, in *pbauth.DeleteAuthenticationRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	return s.deleteAuthentication(ctx, in)
}

func (s *Server) deleteAuthentication(ctx context.Context, in *pbauth.DeleteAuthenticationRequest) (*emptypb.Empty, error) {
	claims, err := claimJwt(in.GetAccessToken(), authconf.GetAUTH_JWT_SECRET)
	if err != nil {
		return nil, err
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "strconv.ParseUint: %v", err)
	}

	access, _, err := retrieveTokens(ctx, s.Redis, userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Unauthenticated, "retrieveTokens: %v", err)
	}
	if access != in.GetAccessToken() {
		return nil, status.Error(xerrors.Unauthenticated, "token given mismatch")
	}
	err = deleteTokens(ctx, s.Redis, userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "deleteTokens: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func deleteTokens(ctx context.Context, r *redis.Client, userId uint64) (err error) {
	deletedQt, err := r.HDel(ctx, userIdToKey(userId), accessKey, refreshKey).Result()
	if deletedQt != 2 {
		err = fmt.Errorf("invalid qt of field deleted. Expected 2, got: %v", deletedQt)
		return
	}
	if err != nil {
		err = errors.Wrap(err, "redis.HDel")
		return
	}
	return
}
