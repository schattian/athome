package srvproducts

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/schema"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveCategoriesById(ctx context.Context, in *pbsemantic.RetrieveCategoriesByIdRequest) (*pbsemantic.RetrieveCategoriesResponse, error) {
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.retrieveCategoriesById(ctx, db, in)
}

func (s *Server) retrieveCategoriesById(ctx context.Context, db *sqlx.DB, in *pbsemantic.RetrieveCategoriesByIdRequest) (*pbsemantic.RetrieveCategoriesResponse, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM product_categories WHERE id IN($1) ORDER BY parent_id`, in.GetCategoryIds())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}

	var tree schema.CategoryTree
	for rows.Next() {
		cat := &schema.ProductCategory{}
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
