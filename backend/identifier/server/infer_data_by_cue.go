package server

import (
	"context"

	"github.com/athomecomar/athome/pb/pbidentifier"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func (s *Server) InferDataByCUE(ctx context.Context, in *pbidentifier.InferDataByCUERequest) (*pbidentifier.InferDataByCUEResponse, error) {
	return nil, status.Error(xerrors.Unimplemented, "not implemented yet")
}
