package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/athomecomar/athome/backend/users/pb/pbauth"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/xerrors"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func (s *Server) Sign(ctx context.Context, in *pbuser.SignRequest) (*pbuser.SignResponse, error) {
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
	return s.sign(ctx, conn, in)
}

func (s *Server) sign(ctx context.Context, conn *grpc.ClientConn, in *pbuser.SignRequest) (*pbuser.SignResponse, error) {
	userId, err := handleJwt(in.GetSignToken(), userconf.GetSIGN_JWT_SECRET)
	if err != nil {
		return nil, err
	}

	tokens, err := createTokens(ctx, conn, userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "createAuthToken: %v", err)
	}

	return &pbuser.SignResponse{
		AccessToken:       tokens.GetAccessToken(),
		RefreshToken:      tokens.GetRefreshToken(),
		AccessTokenExpNs:  tokens.GetAccessTokenExpNs(),
		RefreshTokenExpNs: tokens.GetRefreshTokenExpNs(),
	}, nil
}

func createTokens(ctx context.Context, conn *grpc.ClientConn, userId uint64) (*pbauth.CreateAuthenticationResponse, error) {
	c := pbauth.NewAuthClient(conn)
	signJwt, err := createSignToken(userId)
	if err != nil {
		return nil, errors.Wrap(err, "createSignToken")
	}
	authResponse, err := c.CreateAuthentication(ctx, &pbauth.CreateAuthenticationRequest{SignToken: signJwt})
	if err != nil {
		return nil, errors.Wrap(err, "pbuser.Sign")
	}
	return authResponse, nil
}

func handleJwt(token string, secretFn func() string) (uint64, error) {
	claims, err := claimJwt(token, secretFn)
	if err != nil {
		return 0, err
	}
	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return 0, status.Errorf(xerrors.InvalidArgument, "strconv.ParseUint: %v", err)
	}
	return userId, nil
}

func claimJwt(token string, secretFn func() string) (jwt.MapClaims, error) {
	claimableToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing alg: %v", token.Header["alg"])
		}

		return []byte(secretFn()), nil
	})
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "jwt.Parse: %v", err)
	}

	claims, ok := claimableToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, status.Error(xerrors.FailedPrecondition, "jwt isnt claimable")
	}

	if !claimableToken.Valid {
		return nil, status.Errorf(xerrors.InvalidArgument, "claimable token is not valid")
	}

	if err := claims.Valid(); err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "claimed jwt.Valid: %v", err)
	}

	return claims, nil
}

func createSignToken(userId uint64) (string, error) {
	return createToken(userId, userconf.GetSIGN_JWT_SECRET, userconf.GetSIGN_JWT_EXP)
}

func createForgotToken(userId uint64) (string, error) {
	return createToken(userId, userconf.GetFORGOT_JWT_SECRET, userconf.GetFORGOT_JWT_EXP)
}

func createToken(userId uint64, secretFn func() string, expFn func() time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(expFn()).Unix(),
		"nbf":     time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretFn()))
	if err != nil {
		return "", errors.Wrap(err, "jwt.SignedString")
	}
	return tokenString, nil
}
