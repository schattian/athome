package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/athomecomar/athome/backend/users/pbuser"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/xerrors"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/status"
)

func (s *Server) Sign(ctx context.Context, in *pbuser.SignRequest) (*pbuser.SignResponse, error) {
	claims, err := claimJwt(in.GetSignJwt(), userconf.GetSIGN_JWT_SECRET)
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

	return &pbuser.SignResponse{Jwt: token}, nil
}

func retrieveAuthToken(signToken string) (string, error) {
	conn, err := grpc.Dial(userconf.GetAUTH_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", errors.Wrapf(err, "grpc.Dial at: %v", userconf.GetAUTH_ADDR())
	}
	defer conn.Close()
	c := pbuser.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	authResponse, err := c.Sign(ctx, &pbuser.SignRequest{Jwt: signToken})
	if err != nil {
		return "", errors.Wrap(err, "pbuser.Sign")
	}
	return authResponse.GetJwt(), nil
}

func claimJwt(token string, secretFn func() string) (jwt.MapClaims, error) {
	claimableToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing alg: %v", token.Header["alg"])
		}

		return []byte(secretFn()), nil
	})
	if err != nil {
		return nil, status.Error(xerrors.Internal, "jwt isnt parsable")
	}

	claims, ok := claimableToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, status.Error(xerrors.FailedPrecondition, "jwt isnt claimable")
	}

	if !claimableToken.Valid {
		return nil, status.Errorf(xerrors.Unauthenticated, "claimable Jwt.Valid: %v", err)
	}

	if err := claims.Valid(); err != nil {
		return nil, status.Errorf(xerrors.Unauthenticated, "claimed jwt.Valid: %v", err)
	}

	return claims, nil
}
