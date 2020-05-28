package server

import (
	"github.com/athomecomar/athome/backend/identifier/pb/pbidentifier"
)

type Server struct {
	pbidentifier.UnimplementedIdentifierServer
}
