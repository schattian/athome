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

func userIdToKey(uid uint64) string {
	return string(strconv.Itoa(int(uid)))
}

const accessKey = "access"
const refreshKey = "refresh"
