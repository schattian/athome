package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/athomecomar/athome/backend/auth/authconf"
	"github.com/athomecomar/xerrors"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"

	"github.com/athomecomar/athome/backend/auth/pbauth"
	"github.com/dgrijalva/jwt-go"
)

func (s *Server) Sign(ctx context.Context, in *pbauth.SignRequest) (*pbauth.SignResponse, error) {
	claims, err := claimJwt(in.GetSignJwt(), authconf.GetSIGN_JWT_SECRET)
	if err != nil {
		return nil, err
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "strconv.ParseUint: %v", err)
	}

	token, err := createAuthToken(userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "createAuthToken: %v", err)
	}

	return &pbauth.SignResponse{Jwt: token}, nil
}

func createAuthToken(userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(authconf.GetAUTHENTICATE_JWT_EXP()).Unix(),
		"nbf":     time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(authconf.GetAUTHENTICATE_JWT_SECRET()))
	if err != nil {
		return "", errors.Wrap(err, "jwt.SignedString")
	}

	return tokenString, nil
}
