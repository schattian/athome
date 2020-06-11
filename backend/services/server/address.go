package server

import (
	"context"

	"github.com/athomecomar/athome/backend/services/serviceconf"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ConnAddress(ctx context.Context) (pbaddress.AddressesClient, func() error, error) {
	conn, err := grpc.Dial(serviceconf.GetADDRESS_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, serviceconf.GetADDRESS_ADDR())
	}
	c := pbaddress.NewAddressesClient(conn)
	return c, conn.Close, nil
}