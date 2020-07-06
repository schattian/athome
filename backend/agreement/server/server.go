package server

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"
	"time"

	"github.com/athomecomar/athome/backend/agreement/agreementconf"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type Server struct {
	Redis *redis.Client
}

type randFunc func(s int) (uint64, error)

func ttl(ctx context.Context, r *redis.Client, userId uint64) (time.Duration, error) {
	return r.TTL(ctx, userIdToKey(userId)).Result()
}

func retrieveOrCreateToken(ctx context.Context, r *redis.Client, userId uint64, randFn randFunc) (uint64, error) {
	val, err := retrieveToken(ctx, r, userId)
	if err != nil {
		return 0, errors.Wrap(err, "redis.Get")
	}
	if val == 0 {
		val, err = createToken(ctx, r, userId, randFn)
	}
	if err != nil {
		return 0, errors.Wrap(err, "redis.Get")
	}
	return val, nil
}

func parseToken(raw string) (uint64, error) {
	if raw == "" {
		return 0, nil
	}
	token, err := strconv.Atoi(raw)
	if err != nil {
		return 0, errors.Wrap(err, "strconv.Atoi")
	}
	return uint64(token), nil
}

func retrieveToken(ctx context.Context, r *redis.Client, userId uint64) (uint64, error) {
	val, err := r.Get(ctx, userIdToKey(userId)).Result()
	if errors.Is(err, redis.Nil) {
		val, err = "", nil
	}
	if err != nil {
		return 0, errors.Wrap(err, "redis.Get")
	}
	token, err := parseToken(val)
	if err != nil {
		return 0, errors.Wrap(err, "parseToken")
	}
	return uint64(token), nil
}

func createToken(ctx context.Context, r *redis.Client, userId uint64, randFn randFunc) (uint64, error) {
	token, err := randFn(6)
	if err != nil {
		return 0, errors.Wrap(err, "randString")
	}
	_, err = r.Set(ctx, userIdToKey(userId), token, agreementconf.GetAGREEMENT_TOKEN_EXP()).Result()
	if err != nil {
		return 0, errors.Wrap(err, "redis.Set")
	}
	return token, nil
}

func userIdToKey(uid uint64) string {
	return strconv.Itoa(int(uid))
}

func randInt(sz int) (uint64, error) {
	if sz <= 0 {
		return 0, errors.New("invalid len given (<=0)")
	}

	var strNum string
	for len(strNum) != sz {
		num, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			return 0, errors.Wrap(err, "rand.Int")
		}
		strNum += strconv.Itoa(int(num.Int64()) + 1)
	}
	token, err := parseToken(strNum)
	if err != nil {
		return 0, errors.Wrap(err, "parseToken")
	}
	return uint64(token), nil
}
