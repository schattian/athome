package srvcreator

import (
	"context"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/pb/pbproducts"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) FetchDraft(ctx context.Context, in *pbproducts.FetchDraftRequest) (*pbproducts.FetchDraftResponse, error) {
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
	return s.fetchDraft(ctx, db, draft)
}

func (s *Server) fetchDraft(ctx context.Context, db *sqlx.DB, d *ent.Draft) (*pbproducts.FetchDraftResponse, error) {
	lns, err := d.Lines(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Lines: %v", err)
	}
	var pbLns []*pbproducts.DraftLine
	for _, ln := range lns {
		pbLns = append(pbLns, draftLineToPbDraftLine(ln))
	}
	resp := draftToPbDraft(d)
	resp.Lines = pbLns
	return resp, nil
}

func draftToPbDraft(d *ent.Draft) *pbproducts.FetchDraftResponse {
	return &pbproducts.FetchDraftResponse{
		DraftId: d.Id,
		Stage:   uint64(d.Stage),
	}
}
