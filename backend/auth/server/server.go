package server

import (
	"github.com/athomecomar/athome/backend/auth/pb/pbauth"
	"github.com/go-redis/redis/v8"
)

type Server struct {
	Redis *redis.Client

	pbauth.UnimplementedAuthServer
}
