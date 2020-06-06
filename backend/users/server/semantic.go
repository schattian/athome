package server

import (
	"context"
	"fmt"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ConnSemantic(ctx context.Context, role field.Role) (CategoriesClient, func() error, error) {
	conn, err := grpc.Dial(userconf.GetSEMANTIC_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, userconf.GetSEMANTIC_ADDR())
	}

	var c CategoriesClient
	switch role {
	case field.Merchant:
		c = pbsemantic.NewMerchantsClient(conn)
	case field.ServiceProvider:
		c = pbsemantic.NewServiceProvidersClient(conn)
	default:
		return nil, nil, fmt.Errorf("invalid role given to use semclient: %v", role)
	}

	return c, conn.Close, nil
}

type CategoriesClient interface {
	RetrieveCategories(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pbsemantic.RetrieveCategoriesResponse, error)
}
