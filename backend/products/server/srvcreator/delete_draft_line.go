package srvcreator

import (
	"context"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteDraftLine(ctx context.Context, in *pbproducts.DeleteDraftLineRequest) (*emptypb.Empty, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sem, semCloser, err := pbutil.ConnSemanticProducts(ctx)
	if err != nil {
		return nil, err
	}
	defer semCloser()

	img, imgCloser, err := pbutil.ConnImages(ctx)
	if err != nil {
		return nil, err
	}
	defer imgCloser()

	draft, err := server.RetrieveLatestDraft(ctx, db, in.GetAccessToken())
	if err != nil {
		return nil, err
	}

	if draft.Stage != stage.Second {
		return nil, status.Errorf(xerrors.InvalidArgument, "stage expected < %v, got %v", stage.Second, draft.Stage)
	}

	return s.deleteDraftLine(ctx, db, sem, img, in, draft)
}

func (s *Server) deleteDraftLine(ctx context.Context, db *sqlx.DB, sem pbsemantic.ProductsClient, img pbimages.ImagesClient, in *pbproducts.DeleteDraftLineRequest, draft *ent.Draft) (*emptypb.Empty, error) {
	ln, err := draft.LineById(ctx, db, in.GetDraftLineId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "LineById: %v", err)
	}

	err = deleteAttributes(ctx, sem, ln, in.GetAccessToken())
	if err != nil {
		return nil, err
	}

	err = deleteImages(ctx, db, img, ln, in.GetAccessToken())
	if err != nil {
		return nil, err
	}

	err = storeql.DeleteFromDB(ctx, db, ln)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.DeleteFromDB: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func deleteAttributes(ctx context.Context, c pbsemantic.ProductsClient, from storeql.Storable, access string) error {
	req := &pbsemantic.DeleteAttributeDatasRequest{
		AccessToken: access,
		EntityTable: from.SQLTable(),
		EntityId:    from.GetId(),
	}
	_, err := c.DeleteAttributeDatas(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func deleteImages(ctx context.Context, db *sqlx.DB, c pbimages.ImagesClient, ln *ent.DraftLine, access string) error {
	req := &pbimages.DeleteImagesRequest{
		AccessToken: access,
		EntityId:    ln.Id,
		EntityTable: ln.SQLTable(),
	}
	_, err := c.DeleteImages(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
