package server

import "github.com/athomecomar/athome/backend/users/pb/pbuser"

type Server struct {
	pbuser.UnimplementedUserServer
}
