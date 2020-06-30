package server

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"strconv"

	"github.com/athomecomar/athome/backend/agreement/agreementconf"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type Server struct {
	Redis *redis.Client
}

func retrieveToken(ctx context.Context, r *redis.Client, userId uint64) (string, error) {
	val, err := r.Get(ctx, userIdToKey(userId)).Result()
	if errors.Is(err, redis.Nil) {
		val, err = createToken(ctx, r, userId, randString)
	}
	if err != nil {
		return "", errors.Wrap(err, "redis.Get")
	}
	return val, nil
}

func createToken(ctx context.Context, r *redis.Client, userId uint64, randFunc func(s int) (string, error)) (string, error) {
	token, err := randString(32)
	if err != nil {
		return "", errors.Wrap(err, "randString")
	}
	token, err = r.Set(ctx, userIdToKey(userId), token, agreementconf.GetAGREEMENT_TOKEN_EXP()).Result()
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
