package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/athomecomar/athome/backend/auth/authconf"
	"github.com/athomecomar/athome/backend/auth/pbauth"
	"github.com/athomecomar/xerrors"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveAuthentication(ctx context.Context, in *pbauth.RetrieveAuthenticationRequest) (*pbauth.RetrieveAuthenticationResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	return s.retrieveAuthentication(ctx, in)
}

func (s *Server) retrieveAuthentication(ctx context.Context, in *pbauth.RetrieveAuthenticationRequest) (*pbauth.RetrieveAuthenticationResponse, error) {
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
		return nil, status.Errorf(xerrors.Internal, "retrieveTokens: %v", err)
	}
	if access != in.GetAccessToken() {
		return nil, status.Error(xerrors.Unauthenticated, "token given mismatch")
	}
	return &pbauth.RetrieveAuthenticationResponse{UserId: userId}, nil
}

func retrieveTokens(ctx context.Context, r *redis.Client, userId uint64) (access string, refresh string, err error) {
	vals, err := r.HMGet(ctx, userIdToKey(userId), accessKey, refreshKey).Result()
	if err != nil {
		err = errors.Wrap(err, "redis.HMGet")
		return
	}
	if vlen := len(vals); vlen != 2 {
		err = fmt.Errorf("len of hmget mismatch. Expected 2, got: %d", vlen)
		return
	}

	access, ok := vals[0].(string)
	if !ok {
		err = fmt.Errorf("invalid value type of value key: %s", accessKey)
		return
	}

	refresh, ok = vals[1].(string)
	if !ok {
		err = fmt.Errorf("invalid value type of value key: %s", refreshKey)
		return
	}

	return
}
