package srvcreator

import (
	"context"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/athome/backend/products/pb/pbproducts"
	"github.com/athomecomar/athome/backend/products/pb/pbsemantic"
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
	sem, semCloser, err := server.ConnSemantic(ctx)
	if err != nil {
		return nil, err
	}
	defer semCloser()

	draft, err := server.FetchLatestDraft(ctx, db, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	return s.fetchDraft(ctx, db, sem, draft)
}

func (s *Server) fetchDraft(ctx context.Context, db *sqlx.DB, sem pbsemantic.ProductsClient, d *ent.Draft) (*pbproducts.FetchDraftResponse, error) {
	lns, err := d.Lines(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Lines: %v", err)
	}

	resp := draftToPbDraft(d)

	for _, ln := range lns {
		var atts []*pbproducts.AttributeData
		if d.Stage >= stage.Second {
			semResp, err := sem.GetAttributesData(ctx, &pbsemantic.GetAttributesDataRequest{EntityId: ln.GetId(), EntityTable: ln.SQLTable()})
			if err != nil {
				return nil, err
			}
			atts = server.PbSemanticGetAttributesDataToPbProductAttributes(semResp)
		}
		resp.Lines = append(resp.Lines, draftLineToPbDraftLine(ln, atts))
	}

	return resp, nil
}

func draftToPbDraft(d *ent.Draft) *pbproducts.FetchDraftResponse {
	return &pbproducts.FetchDraftResponse{
		DraftId: d.Id,
		Stage:   uint64(d.Stage),
	}
}
