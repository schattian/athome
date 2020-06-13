package server

import (
	"context"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/xpbsemantic"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func ConnCategories(ctx context.Context, role field.Role) (c xpbsemantic.CategoriesClient, closer func() error, err error) {
	switch role {
	case field.Merchant:
		c, closer, err = pbconf.ConnSemanticMerchants(ctx)
	case field.ServiceProvider:
		c, closer, err = pbconf.ConnSemanticServiceProviders(ctx)
	default:
		err = status.Errorf(xerrors.InvalidArgument, "invalid role given to use semclient: %v", role)
	}
	return
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
