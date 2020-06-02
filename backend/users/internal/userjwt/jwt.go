package userjwt

import (
	"fmt"
	"time"

	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/xerrors"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func ClaimJwt(token string, secretFn func() string) (jwt.MapClaims, error) {
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

func CreateSignToken(userId uint64) (string, error) {
	return createToken(userId, userconf.GetSIGN_JWT_SECRET, userconf.GetSIGN_JWT_EXP)
}

func CreateForgotToken(userId uint64) (string, error) {
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
