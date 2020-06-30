package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/athomecomar/athome/backend/auth/authconf"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/xerrors"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"

	"github.com/dgrijalva/jwt-go"
)

func (s *Server) CreateAuthentication(ctx context.Context, in *pbauth.CreateAuthenticationRequest) (*pbauth.CreateAuthenticationResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	return s.createAuthentication(ctx, in)
}

func (s *Server) createAuthentication(ctx context.Context, in *pbauth.CreateAuthenticationRequest) (*pbauth.CreateAuthenticationResponse, error) {
	claims, err := claimJwt(in.GetSignToken(), authconf.GetSIGN_JWT_SECRET)
	if err != nil {
		return nil, err
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "strconv.ParseUint: %v", err)
	}

	accessToken, err := createAccessToken(userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "createAccessToken: %v", err)
	}

	refreshToken, err := createRefreshToken(userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "createRefreshToken: %v", err)
	}

	err = s.Redis.HMSet(ctx, userIdToKey(userId),
		map[string]interface{}{
			accessKey:  accessToken,
			refreshKey: refreshToken,
		}).Err()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "redis.HSet: %v", err)
	}
	return &pbauth.CreateAuthenticationResponse{
		AccessToken: accessToken, RefreshToken: refreshToken,
		AccessTokenExpNs:  uint64(authconf.GetAUTH_JWT_EXP().Nanoseconds()),
		RefreshTokenExpNs: uint64(authconf.GetAUTH_JWT_REFRESH_EXP().Nanoseconds()),
	}, nil
}

func createAccessToken(userId uint64) (string, error) {
	return createToken(userId, authconf.GetAUTH_JWT_EXP(), authconf.GetAUTH_JWT_SECRET)
}

func createRefreshToken(userId uint64) (string, error) {
	return createToken(userId, authconf.GetAUTH_JWT_REFRESH_EXP(), authconf.GetAUTH_JWT_SECRET)
}

func createToken(userId uint64, extraExp time.Duration, secretFn func() string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(extraExp).Unix(),
		"iat":     time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretFn()))
	if err != nil {
		return "", errors.Wrap(err, "jwt.SignedString")

	}
	return tokenString, nil
}
