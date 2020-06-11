package server

import (
	"context"
	"fmt"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/xpbsemantic"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/athome/pb/pbidentifier"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ConnCategories(ctx context.Context, role field.Role) (xpbsemantic.CategoriesClient, func() error, error) {
	conn, connCloser, err := ConnSemantic(ctx)
	if err != nil {
		return nil, nil, err
	}

	var c xpbsemantic.CategoriesClient
	switch role {
	case field.Merchant:
		c = pbsemantic.NewMerchantsClient(conn)
	case field.ServiceProvider:
		c = pbsemantic.NewServiceProvidersClient(conn)
	default:
		return nil, nil, fmt.Errorf("invalid role given to use semclient: %v", role)
	}

	return c, connCloser, nil
}

func ConnSemantic(ctx context.Context) (*grpc.ClientConn, func() error, error) {
	conn, err := grpc.Dial(userconf.GetSEMANTIC_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, userconf.GetSEMANTIC_ADDR())
	}
	return conn, conn.Close, nil
}

func ConnIdentifier(ctx context.Context) (pbidentifier.IdentifierClient, func() error, error) {
	conn, err := grpc.Dial(userconf.GetIDENTIFIER_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, userconf.GetIDENTIFIER_ADDR())
	}
	c := pbidentifier.NewIdentifierClient(conn)
	return c, conn.Close, nil
}

func PbSemanticCategoryToPbUserCategory(c *pbsemantic.Category) *pbusers.Category {
	childs := make(map[uint64]*pbusers.Category)
	for id, child := range c.GetChilds() {
		childs[id] = PbSemanticCategoryToPbUserCategory(child)
	}
	return &pbusers.Category{
		Name:                   c.GetName(),
		Childs:                 childs,
		ParentId:               c.GetParentId(),
		IdentificationTemplate: c.GetIdentificationTemplate(),
	}
}
