package pbutil

import (
	"context"

	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/storeql"
)

func GetUserFromAccessToken(ctx context.Context, c pbauth.AuthClient, access string) (uint64, error) {
	resp, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access})
	if err != nil {
		return 0, err
	}

	return resp.GetUserId(), nil
}

func ToPbSemanticEntity(s storeql.Storable) *pbsemantic.Entity {
	return &pbsemantic.Entity{
		EntityId:    s.GetId(),
		EntityTable: s.SQLTable(),
	}
}
func ToPbImagesEntity(s storeql.Storable) *pbimages.Entity {
	return &pbimages.Entity{
		EntityId:    s.GetId(),
		EntityTable: s.SQLTable(),
	}
}
