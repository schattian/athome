package srvcreator

import (
	"context"
	"database/sql"
	"errors"
	"io"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/athome/backend/products/pb/pbproducts"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) First(srv pbproducts.Creator_FirstServer) error {
	ctx := srv.Context()

	db, err := server.ConnDB()
	if err != nil {
		return err
	}

	var resp *pbproducts.FirstResponse
	var draft *ent.Draft
	for {
		select {
		case <-ctx.Done():
			err = srv.SendAndClose(resp)
			if err != nil {
				return err
			}
			return ctx.Err()
		default: // no-op
		}

		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = in.Validate()
		if err != nil {
			return err
		}

		if draft == nil { // first iteration
			draft, err = server.RetrieveLatestDraft(ctx, db, in.GetAccessToken())
			if err != nil {
				return err
			}
			continue // see oneof def
		}

		if draft.Stage != stage.First {
			return status.Errorf(xerrors.InvalidArgument, "stage expected %v, got %v", stage.First, draft.Stage)
		}

		resp, err = s.first(ctx, db, in.GetDraftLine(), draft)
		if err != nil {
			return err
		}
	}
}

func (s *Server) first(ctx context.Context, db *sqlx.DB, in *pbproducts.DraftLineFirst, draft *ent.Draft) (*pbproducts.FirstResponse, error) {
	ln, err := draft.LineByTitle(ctx, db, in.GetTitle())
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "draft.LineByTitle: %v", err)
	}
	if ln != nil {
		return nil, status.Errorf(xerrors.Internal, "couldnt store >1 draft lines with same title: %v", in.GetTitle())
	}

	ln = firstRequestToDraftLine(in)
	err = draft.ValidateLineByStage(ln)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "ValidateLineByStage: %v", err)
	}
	ln.DraftId = draft.Id

	err = storeql.InsertIntoDB(ctx, db, ln)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}

	return draftToFirstResponse(draft), nil
}

func firstRequestToDraftLine(in *pbproducts.DraftLineFirst) *ent.DraftLine {
	return &ent.DraftLine{
		Title:      in.GetTitle(),
		CategoryId: in.GetCategoryId(),
	}
}

func draftToFirstResponse(d *ent.Draft) *pbproducts.FirstResponse {
	return &pbproducts.FirstResponse{
		DraftId: d.Id,
	}
}
