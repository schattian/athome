package srvcreator

import (
	"context"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/athome/backend/products/pb/pbproducts"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) CloneDraftLine(ctx context.Context, in *pbproducts.CloneDraftLineRequest) (*pbproducts.CloneDraftLineResponse, error) {
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}

	err = in.Validate()
	if err != nil {
		return nil, err
	}

	draft, err := server.FetchLatestDraft(ctx, db, in.GetAccessToken())
	if err != nil {
		return nil, err
	}

	if draft.Stage != stage.Second {
		return nil, status.Errorf(xerrors.InvalidArgument, "stage expected %v, got %v", stage.Second, draft.Stage)
	}

	return s.cloneDraftLine(ctx, db, in, draft)
}

func (s *Server) cloneDraftLine(ctx context.Context, db *sqlx.DB, in *pbproducts.CloneDraftLineRequest, draft *ent.Draft) (*pbproducts.CloneDraftLineResponse, error) {
	ln, err := ent.LineById(ctx, db, in.GetDraftLineId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "LineById: %v", err)
	}

	return &pbproducts.CloneDraftLineResponse{DraftLine: draftLineToPbDraftLine(ln)}, nil
}

func draftLineToPbDraftLine(ln *ent.DraftLine) *pbproducts.DraftLine {
	return &pbproducts.DraftLine{
		DraftLineId: ln.Id,

		First: &pbproducts.DraftLineFirst{
			Title:      ln.Title,
			CategoryId: ln.CategoryId,
		},

		Second: &pbproducts.DraftLineSecond{
			Price:      ln.Price.Float64(),
			Stock:      ln.Stock,
			Attributes: nil,
		},
	}
}
