package server

import "github.com/athomecomar/athome/backend/users/pbuser"

type Server struct {
	pbuser.UnimplementedUserServer
}
