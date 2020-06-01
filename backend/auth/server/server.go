package server

import (
	"github.com/go-redis/redis/v8"
)

type Server struct {
	Redis *redis.Client
}
