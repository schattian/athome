package srvproviders

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/schema"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveShippingCategories(ctx context.Context, in *pbsemantic.RetrieveShippingCategoriesRequest) (*pbsemantic.RetrieveCategoriesResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.retrieveShippingCategories(ctx, db, in)
}

func (s *Server) retrieveShippingCategories(ctx context.Context, db *sqlx.DB, in *pbsemantic.RetrieveShippingCategoriesRequest) (*pbsemantic.RetrieveCategoriesResponse, error) {
	rows, err := db.QueryxContext(
		ctx,
		`SELECT * FROM service_provider_categories WHERE max_vol_weight >= $1 ORDER BY parent_id`,
		in.GetMaxVolWeight(),
	)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}
	var tree schema.CategoryTree
	for rows.Next() {
		cat := &schema.ServiceProviderCategory{}
		err = rows.StructScan(cat)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "StructScan: %v", err)
		}

		tree, err = schema.AddCategoryToTree(tree, cat)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "AddCategoryToTree: %v", err)
		}
	}

	return &pbsemantic.RetrieveCategoriesResponse{Categories: tree.ToPb()}, nil
}
