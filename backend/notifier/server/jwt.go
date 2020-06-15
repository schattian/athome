package server

import (
	"fmt"
	"strconv"

	"github.com/athomecomar/xerrors"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/status"
)

func claimJwt(token string, secretFn func() string) (jwt.MapClaims, error) {
	claimableToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing alg: %v", token.Header["alg"])
		}

		return []byte(secretFn()), nil
	})
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "jwt isnt parsable: %v", err)
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

func userIdFromClaim(claims jwt.MapClaims) (uint64, error) {
	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return 0, status.Errorf(xerrors.InvalidArgument, "strconv.ParseUint: %v", err)
	}
	return userId, nil
}
