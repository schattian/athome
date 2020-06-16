package pbutil

import (
	"context"
	"time"

	"github.com/athomecomar/athome/pb/pbnotifier"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/status"
)

type Priority uint64

const (
	Low Priority = iota + 1
	Mid
	High
	Max
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "low"
	case Mid:
		return "mid"
	case High:
		return "high"
	case Max:
		return "max"
	}
	return ""
}

func CreateNotification(
	ctx context.Context,
	entity storeql.Storable,
	userId uint64,
	secretFn func() string,
	p Priority,
	c pbnotifier.NotificationsClient,
) (*pbnotifier.CreateResponse, error) {
	var exp = func() time.Duration {
		return 10 * time.Minute
	}
	token, err := CreateJWT(userId, secretFn, exp)
	if err != nil {
		return nil, err
	}
	return c.Create(ctx, &pbnotifier.CreateRequest{
		NotificationToken: token,
		Notification: &pbnotifier.Notification{
			UserId:   userId,
			Priority: p.String(),
			Entity:   ToPbNotifierEntity(entity),
		},
	})
}

func CreateJWT(userId uint64, secretFn func() string, expFn func() time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(expFn()).Unix(),
		"nbf":     time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretFn()))
	if err != nil {
		return "", status.Errorf(xerrors.Internal, "jwt.SignedString: %v", err)
	}

	return tokenString, nil

}
