package server

import (
	"context"

	"github.com/athomecomar/athome/backend/services/serviceconf"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ConnUsers(ctx context.Context) (pbusers.ViewerClient, func() error, error) {
	conn, err := grpc.Dial(serviceconf.GetUSERS_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, serviceconf.GetUSERS_ADDR())
	}
	c := pbusers.NewViewerClient(conn)
	return c, conn.Close, nil
}
