package srvcreator

import (
	"context"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveDraft(ctx context.Context, in *pbproducts.RetrieveDraftRequest) (*pbproducts.RetrieveDraftResponse, error) {
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

	return s.retrieveDraft(ctx, draft)
}

func (s *Server) retrieveDraft(ctx context.Context, d *ent.Draft) (*pbproducts.RetrieveDraftResponse, error) {
	return &pbproducts.RetrieveDraftResponse{DraftId: d.Id, Draft: d.ToPb()}, nil
}

func (s *Server) RetrieveDraftLines(ctx context.Context, in *pbproducts.RetrieveDraftLinesRequest) (*pbproducts.RetrieveDraftLinesResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}

	draft, err := server.RetrieveLatestDraft(ctx, db, in.GetAccessToken())
	if err != nil {
		return nil, err
	}

	sem, semCloser, err := pbutil.ConnSemanticProducts(ctx)
	if err != nil {
		return nil, err
	}
	defer semCloser()

	return s.retrieveDraftLines(ctx, db, sem, draft)
}

func (s *Server) retrieveDraftLines(ctx context.Context, db *sqlx.DB, sem pbsemantic.ProductsClient, d *ent.Draft) (*pbproducts.RetrieveDraftLinesResponse, error) {
	lns, err := d.Lines(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Lines: %v", err)
	}

	resp := &pbproducts.RetrieveDraftLinesResponse{}
	resp.DraftLines = make(map[uint64]*pbproducts.DraftLine)
	for _, ln := range lns {
		var atts []*pbproducts.AttributeData
		if d.Stage >= stage.Second {
			semResp, err := sem.RetrieveAttributeDatas(ctx,
				&pbsemantic.RetrieveAttributeDatasRequest{Entity: pbutil.ToPbEntity(ln)},
			)
			if err != nil {
				return nil, err
			}
			atts = server.PbSemanticRetrieveAttributesDataToPbProductAttributes(semResp)
		}
		resp.DraftLines[ln.Id] = ln.ToPb(atts)
	}

	return resp, nil
}
