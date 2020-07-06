package server

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
)

func connMiniredis(t *testing.T) *miniredis.Miniredis {
	t.Helper()
	miniredis, err := miniredis.Run()
	if err != nil {
		t.Fatalf("miniredis failed to conn: %v", err)
	}

	return miniredis
}

type record struct {
	UserId         uint64 `json:"user_id,omitempty"`
	AgreementToken uint64 `json:"agreement_token,omitempty"`
}

func (vt *record) save(t *testing.T, redis *redis.Client, exp time.Duration) {
	t.Helper()
	err := redis.Set(context.Background(), userIdToKey(vt.UserId), vt.AgreementToken, exp).Err()
	if err != nil {
		t.Fatalf("redis.Set: %v", err)
	}

}

func redisCli(miniredis *miniredis.Miniredis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: miniredis.Addr(),
	})
}
