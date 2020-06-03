package srvproviders

import "github.com/athomecomar/athome/backend/users/pb/pbsemantic"

type Server struct {
	pbsemantic.UnimplementedServiceProvidersServer
}
