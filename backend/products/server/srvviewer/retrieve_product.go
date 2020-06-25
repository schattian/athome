package srvviewer

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveProduct(ctx context.Context, in *pbproducts.RetrieveProductRequest) (*pbproducts.Product, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return s.retrieveProduct(ctx, db, in)
}

func (s *Server) retrieveProduct(ctx context.Context, db *sqlx.DB,
	in *pbproducts.RetrieveProductRequest,
) (*pbproducts.Product, error) {
	prod, err := ent.FindProduct(ctx, db, in.GetProductId())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "product with id %d wasnt found", in.GetProductId())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProduct: %v", err)
	}
	return prod.ToPb(), nil
}
