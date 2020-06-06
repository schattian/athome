package srvmerchants

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/semantic/schema"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) RetrieveCategories(ctx context.Context, _ *emptypb.Empty) (*pbsemantic.RetrieveCategoriesResponse, error) {
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.getCategories(ctx, db)
}

func (s *Server) getCategories(ctx context.Context, db *sqlx.DB) (*pbsemantic.RetrieveCategoriesResponse, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM merchant_categories ORDER BY parent_id`)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}

	var tree schema.CategoryTree
	for rows.Next() {
		cat := &schema.MerchantCategory{}
		err = rows.StructScan(cat)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "StructScan: %v", err)
		}

		tree, err = schema.AddCategoryToTree(tree, cat)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "AddCategoryToTree: %v", err)
		}
	}

	return server.CategoryTreeToRetrieveCategoriesResponse(tree), nil
}
