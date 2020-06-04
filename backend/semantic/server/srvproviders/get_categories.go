package srvproviders

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/ent"
	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetCategories(ctx context.Context, _ *emptypb.Empty) (*pbsemantic.GetCategoriesResponse, error) {
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.getCategories(ctx, db)
}

func (s *Server) getCategories(ctx context.Context, db *sqlx.DB) (*pbsemantic.GetCategoriesResponse, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM service_provider_categories ORDER BY parent_id`)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}

	var tree ent.CategoryTree
	for rows.Next() {
		cat := &ent.ServiceProviderCategory{}
		err = rows.StructScan(cat)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "StructScan: %v", err)
		}

		tree, err = ent.AddCategoryToTree(tree, cat)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "AddCategoryToTree: %v", err)
		}
	}

	return server.CategoryTreeToGetCategoriesResponse(tree), nil
}
