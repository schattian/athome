package srvcreator

import (
	"context"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/athome/backend/products/pb/pbproducts"
	"github.com/athomecomar/athome/backend/products/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) CloneDraftLine(ctx context.Context, in *pbproducts.CloneDraftLineRequest) (*pbproducts.CloneDraftLineResponse, error) {
	err := in.Validate()
	if err != nil {
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

	draft, err := server.RetrieveLatestDraft(ctx, db, in.GetAccessToken())
	if err != nil {
		return nil, err
	}

	if draft.Stage != stage.Second {
		return nil, status.Errorf(xerrors.InvalidArgument, "stage expected < %v, got %v", stage.Second, draft.Stage)
	}

	return s.cloneDraftLine(ctx, db, sem, in, draft)
}

func (s *Server) cloneDraftLine(ctx context.Context, db *sqlx.DB, sem pbsemantic.ProductsClient, in *pbproducts.CloneDraftLineRequest, draft *ent.Draft) (*pbproducts.CloneDraftLineResponse, error) {
	ln, err := draft.LineById(ctx, db, in.GetDraftLineId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "LineById: %v", err)
	}
	cpLn, err := ln.Clone()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Clone: %v", err)
	}

	err = storeql.InsertIntoDB(ctx, db, cpLn)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}

	atts, err := cloneAttributes(ctx, sem, ln, cpLn, in.GetAccessToken())
	if err != nil {
		return nil, err
	}

	return &pbproducts.CloneDraftLineResponse{DraftLine: cpLn.ToPb(atts)}, nil
}

func cloneAttributes(ctx context.Context, c pbsemantic.ProductsClient, from, dest storeql.Storable, access string) ([]*pbproducts.AttributeData, error) {
	if from.SQLTable() != dest.SQLTable() {
		return nil, status.Error(xerrors.InvalidArgument, "couldnt clone attributes from different entity table")
	}
	req := &pbsemantic.CloneAttributeDatasRequest{
		AccessToken:  access,
		EntityTable:  from.SQLTable(),
		DestEntityId: dest.GetId(),
		FromEntityId: from.GetId(),
	}
	resp, err := c.CloneAttributeDatas(ctx, req)
	if err != nil {
		return nil, err
	}
	var atts []*pbproducts.AttributeData
	for _, respAtt := range resp.Attributes {
		atts = append(atts, server.PbSemanticToPbProductAttributes(respAtt))
	}
	return atts, nil
}
