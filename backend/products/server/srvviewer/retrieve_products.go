package srvviewer

import (
	"context"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveProducts(ctx context.Context, in *pbproducts.RetrieveProductsRequest) (*pbproducts.RetrieveProductsResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return s.retrieveProducts(ctx, db, in)
}

func (s *Server) retrieveProducts(ctx context.Context, db *sqlx.DB,
	in *pbproducts.RetrieveProductsRequest,
) (*pbproducts.RetrieveProductsResponse, error) {
	prods, err := ent.FindProductsById(ctx, db, in.GetIds())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductsById: %v", err)
	}
	resp := &pbproducts.RetrieveProductsResponse{}
	resp.Products = make(map[uint64]*pbproducts.Product)
	for _, prod := range prods {
		resp.Products[prod.Id] = prod.ToPb()
	}
	return resp, nil
}
