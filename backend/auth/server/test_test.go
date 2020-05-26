package server

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/athomecomar/xtest/xload"
	"github.com/go-redis/redis/v8"
)

var (
	gTokens goldenTokens
)

func init() {
	xload.DecodeJsonnet("tokens", &gTokens)
}

type goldenTokens struct {
	Valid          *variadicTokens `json:"valid,omitempty"`
	ExpiredAccess  *variadicTokens `json:"expired_access,omitempty"`
	ExpiredRefresh *variadicTokens `json:"expired_refresh,omitempty"`
	ExpiredSign    *variadicTokens `json:"expired_sign,omitempty"`
	Expired        *variadicTokens `json:"expired,omitempty"`
}

type variadicTokens struct {
	UserId       uint64 `json:"user_id,omitempty"`
	SignToken    string `json:"sign_token,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (vt *variadicTokens) save(t *testing.T, redis *redis.Client) {
	t.Helper()
	err := redis.HMSet(context.Background(), userIdToKey(vt.UserId), vt.toMap()).Err()
	if err != nil {
		t.Fatalf("redis.HSet: %v", err)
	}

}

func (vt *variadicTokens) toMap() map[string]interface{} {
	return map[string]interface{}{
		accessKey:  vt.AccessToken,
		refreshKey: vt.RefreshToken,
	}
}

func redisCli(miniredis *miniredis.Miniredis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: miniredis.Addr(),
	})
}
