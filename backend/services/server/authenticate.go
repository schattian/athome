package server

import (
	"context"

	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbutil"
)

func GetUserFromAccessToken(c pbauth.AuthClient, access string) AuthFunc {
	return func(ctx context.Context) (uint64, error) {
		return pbutil.GetUserFromAccessToken(ctx, c, access)
	}
}

type AuthFunc func(ctx context.Context) (uint64, error)
