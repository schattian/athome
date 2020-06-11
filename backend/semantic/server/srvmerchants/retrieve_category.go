package srvmerchants

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/semantic/schema"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveCategory(ctx context.Context, in *pbsemantic.RetrieveCategoryRequest) (*pbsemantic.Category, error) {
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.retrieveCategory(ctx, db, in)
}

func (s *Server) retrieveCategory(ctx context.Context, db *sqlx.DB, in *pbsemantic.RetrieveCategoryRequest) (*pbsemantic.Category, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM merchant_categories WHERE id=$1`, in.GetCategoryId())
	cat := &schema.MerchantCategory{}
	err := row.StructScan(cat)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "couldn't find category with id: %v", err)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "StructScan: %v", err)
	}

	return schema.CategoryToPb(cat), nil
}
