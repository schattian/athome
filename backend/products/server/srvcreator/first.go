package srvcreator

import (
	"context"
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
	auth, authCloser, err := server.ConnAuth(ctx)
	if err != nil {
		return err
	}

	var userId uint64
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

		if userId == 0 {
			userId, err = server.GetUserFromAccessToken(ctx, db, auth, in.GetAccessToken())
			if err != nil {
				return err
			}
			err = authCloser()
			if err != nil {
				return status.Errorf(xerrors.Internal, "authConn.Close: %v", err)
			}

		}

		resp, err := s.first(ctx, db, in, userId)
		if err != nil {
			return err
		}

		err = srv.Send(resp)
		if err != nil {
			return err
		}
	}
}

func (s *Server) first(ctx context.Context, db *sqlx.DB, in *pbproducts.FirstRequest, userId uint64) (*pbproducts.FirstResponse, error) {
	draft := firstRequestToDraft(in)
	err := draft.ValidateByStage()
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "ValidateByStage: %v", err)
	}

	err = storeql.InsertIntoDB(ctx, db, draft)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}

	return draftTofirstResponse(draft), nil
}

func firstRequestToDraft(in *pbproducts.FirstRequest) *ent.Draft {
	return &ent.Draft{
		Title:      in.GetTitle(),
		CategoryId: in.GetCategoryId(),
		Stage:      stage.First,
	}
}

func draftTofirstResponse(d *ent.Draft) *pbproducts.FirstResponse {
	return &pbproducts.FirstResponse{
		DraftId: d.Id,
	}
}
