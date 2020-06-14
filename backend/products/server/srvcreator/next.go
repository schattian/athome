package srvcreator

import (
	"context"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) Next(ctx context.Context, in *pbproducts.StageChangeRequest) (*pbproducts.StageChangeResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	sem, semCloser, err := pbconf.ConnSemanticProducts(ctx)
	if err != nil {
		return nil, err
	}
	defer semCloser()

	imgs, imgsCloser, err := pbconf.ConnImages(ctx)
	if err != nil {
		return nil, err
	}
	defer imgsCloser()

	draft, err := server.RetrieveLatestDraft(ctx, db, in.GetAccessToken())
	if err != nil {
		return nil, err
	}

	return s.next(ctx, db, imgs, sem, in, draft)
}

func (s *Server) next(ctx context.Context, db *sqlx.DB, imgs pbimages.ImagesClient, sem pbsemantic.ProductsClient, in *pbproducts.StageChangeRequest, d *ent.Draft) (*pbproducts.StageChangeResponse, error) {
	qt, err := d.Next(ctx, db, imgs, sem, in.GetAccessToken())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "draft.Next: %v", err)
	}

	return &pbproducts.StageChangeResponse{Size: uint64(qt)}, nil
}
