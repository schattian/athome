package srvcreator

import (
	"context"
	"io"

	"github.com/athomecomar/athome/backend/products/ent"
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
	auth, authCloser, err := server.ConnAuth(ctx)
	if err != nil {
		return err
	}

	var draft *ent.Draft
	for {
		select {
		case <-ctx.Done():
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
			draft, err = server.FetchLatestDraft(ctx, db, auth, authCloser, in.GetAccessToken())
			if err != nil {
				return err
			}
		}

		resp, err := s.first(ctx, db, in, draft)
		if err != nil {
			return err
		}

		err = srv.Send(resp)
		if err != nil {
			return err
		}
	}
}

func (s *Server) first(ctx context.Context, db *sqlx.DB, in *pbproducts.FirstRequest, draft *ent.Draft) (f *pbproducts.FirstResponse, err error) {
	switch in.IsDeletion {
	case true:
		f, err = s.firstDeletion(ctx, db, in, draft)
	case false:
		f, err = s.firstAddition(ctx, db, in, draft)
	}
	return
}

func (s *Server) firstAddition(ctx context.Context, db *sqlx.DB, in *pbproducts.FirstRequest, draft *ent.Draft) (*pbproducts.FirstResponse, error) {
	ln := firstRequestToDraftLine(in)
	err := draft.ValidateLineByStage(ln)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "ValidateLineByStage: %v", err)
	}
	ln.Id = draft.Id

	err = storeql.InsertIntoDB(ctx, db, ln)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}

	return draftToFirstResponse(draft), nil
}

func (s *Server) firstDeletion(ctx context.Context, db *sqlx.DB, in *pbproducts.FirstRequest, draft *ent.Draft) (*pbproducts.FirstResponse, error) {
	lns, err := draft.Lines(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Lines: %v", err)
	}
	var ln *ent.DraftLine
	for _, ln = range lns {
		if ln.Title == in.GetTitle() {
			break
		}
	}
	err = storeql.DeleteFromDB(ctx, db, ln)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.DeleteFromDB: %v", err)
	}

	return draftToFirstResponse(draft), nil
}

func firstRequestToDraftLine(in *pbproducts.FirstRequest) *ent.DraftLine {
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
