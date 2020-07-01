package server

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"strconv"
	"time"

	"github.com/athomecomar/athome/backend/agreement/agreementconf"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type Server struct {
	Redis *redis.Client
}

func ttl(ctx context.Context, r *redis.Client, userId uint64) (time.Duration, error) {
	return r.TTL(ctx, userIdToKey(userId)).Result()
}

func retrieveOrCreateToken(ctx context.Context, r *redis.Client, userId uint64, randFn randFunc) (string, error) {
	val, err := retrieveToken(ctx, r, userId)
	if val == "" {
		val, err = createToken(ctx, r, userId, randFn)
	}
	if err != nil {
		return "", errors.Wrap(err, "redis.Get")
	}
	return val, nil
}

func retrieveToken(ctx context.Context, r *redis.Client, userId uint64) (string, error) {
	val, err := r.Get(ctx, userIdToKey(userId)).Result()
	if errors.Is(err, redis.Nil) {
		val, err = "", nil
	}
	if err != nil {
		return "", errors.Wrap(err, "redis.Get")
	}
	return val, nil
}

type randFunc func(s int) (string, error)

func createToken(ctx context.Context, r *redis.Client, userId uint64, randFn randFunc) (string, error) {
	token, err := randFn(32)
	if err != nil {
		return "", errors.Wrap(err, "randString")
	}
	_, err = r.Set(ctx, userIdToKey(userId), token, agreementconf.GetAGREEMENT_TOKEN_EXP()).Result()
	if err != nil {
		return "", errors.Wrap(err, "redis.Set")
	}
	return token, nil
}

func userIdToKey(uid uint64) string {
	return strconv.Itoa(int(uid))
}

func randString(s int) (string, error) {
	b, err := randBytes(s)
	if err != nil {
		return "", errors.Wrap(err, "randBytes")
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func randBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, errors.Wrap(err, "rand.Read")
	}
	return b, nil
}
