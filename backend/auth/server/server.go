package server

import (
	"github.com/athomecomar/athome/backend/auth/pbauth"
	"github.com/go-redis/redis/v8"
)

type Server struct {
	Redis *redis.Client

	pbauth.UnimplementedAuthServer
}
