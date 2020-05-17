package server

import "github.com/athomecomar/athome/users/pb/pbuser"

type Server struct {
	pbuser.UnimplementedUserServer
}
