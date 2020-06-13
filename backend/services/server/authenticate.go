package server

import (
	"context"

	"github.com/athomecomar/athome/pb/pbauth"
)

func GetUserFromAccessToken(c pbauth.AuthClient, access string) AuthFunc {
	return func(ctx context.Context) (uint64, error) {
		resp, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access})
		if err != nil {
			return 0, err
		}

		return resp.GetUserId(), nil
	}
}

type AuthFunc func(ctx context.Context) (uint64, error)
