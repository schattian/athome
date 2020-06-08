package srvviewer

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/pb/pbimages"
	"github.com/athomecomar/athome/backend/products/pb/pbproducts"
	"github.com/athomecomar/athome/backend/products/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/products/pb/pbusers"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveProductDetail(ctx context.Context, in *pbproducts.RetrieveProductRequest) (*pbproducts.RetrieveProductDetailResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sem, semCloser, err := server.ConnSemantic(ctx)
	if err != nil {
		return nil, err
	}
	defer semCloser()

	img, imgCloser, err := server.ConnImages(ctx)
	if err != nil {
		return nil, err
	}
	defer imgCloser()

	users, usersCloser, err := server.ConnUsers(ctx)
	if err != nil {
		return nil, err
	}
	defer usersCloser()

	return s.retrieveProductDetail(ctx, db, users, sem, img, in)
}

func (s *Server) retrieveProductDetail(ctx context.Context, db *sqlx.DB,
	users pbusers.ViewerClient, sem pbsemantic.ProductsClient, img pbimages.ImagesClient,
	in *pbproducts.RetrieveProductRequest,
) (*pbproducts.RetrieveProductDetailResponse, error) {
	prod, err := ent.FindProduct(ctx, db, in.GetProductId())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "product with id %d wasnt found", in.GetProductId())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProduct: %v", err)
	}

	atts, err := prod.GetViewableAttributes(ctx, sem)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "GetViewableAttributes: %v", err)
	}
	imgs, err := prod.GetImages(ctx, img)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "GetImages: %v", err)
	}
	user, err := prod.GetUser(ctx, users)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "GetImages: %v", err)
	}
	return &pbproducts.RetrieveProductDetailResponse{
		Title:      prod.Title,
		CategoryId: prod.CategoryId,
		Price:      prod.Price.Float64(),
		Stock:      prod.Stock,
		Attributes: atts,
		ImageUris:  imgs,
		User:       user,
	}, nil
}
