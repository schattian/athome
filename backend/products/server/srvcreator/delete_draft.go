package srvcreator

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteDraft(ctx context.Context, in *pbproducts.DeleteDraftRequest) (*emptypb.Empty, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	draft, err := server.RetrieveLatestDraft(ctx, db, in.GetAccessToken())
	if err != nil {
		return nil, err
	}

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

	return s.deleteDraft(ctx, db, sem, img, in, draft)
}

func (s *Server) deleteDraft(
	ctx context.Context,
	db *sqlx.DB,
	sem pbsemantic.ProductsClient,
	img pbimages.ImagesClient,
	in *pbproducts.DeleteDraftRequest,
	draft *ent.Draft,
) (*emptypb.Empty, error) {
	lns, err := draft.Lines(ctx, db)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "draft.Lines: %v", err)
	}
	for _, ln := range lns {
		_, err = s.deleteDraftLine(ctx, db, sem, img, &pbproducts.DeleteDraftLineRequest{AccessToken: in.GetAccessToken(), DraftLineId: ln.Id}, draft)
		if err != nil {
			return nil, err
		}
	}

	err = storeql.DeleteFromDB(ctx, db, draft)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.DeleteFromDB: %v", err)
	}

	return &emptypb.Empty{}, nil
}
