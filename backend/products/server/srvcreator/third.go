package srvcreator

import (
	"context"
	"io"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Third(srv pbproducts.Creator_ThirdServer) error {
	ctx := srv.Context()

	db, err := server.ConnDB()
	if err != nil {
		return err
	}

	var resp *emptypb.Empty
	var draft *ent.Draft
	var access string
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
			return srv.SendAndClose(resp)
		}
		if err != nil {
			return err
		}
		err = in.Validate()
		if err != nil {
			return err
		}

		if draft == nil { // first iteration
			access = in.GetAccessToken()
			draft, err = server.RetrieveLatestDraft(ctx, db, access)
			if err != nil {
				return err
			}
			continue // see oneof def
		}

		if draft.Stage != stage.Third {
			return status.Errorf(xerrors.InvalidArgument, "stage expected %v, got %v", stage.Third, draft.Stage)
		}

		resp, err = s.third(ctx, db, in.GetBody(), draft)
		if err != nil {
			return err
		}
	}
}

func (s *Server) third(ctx context.Context, db *sqlx.DB, in *pbproducts.ThirdRequest_Body, draft *ent.Draft) (*emptypb.Empty, error) {
	ln, err := draft.LineById(ctx, db, in.GetDraftLineId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "LineById: %v", err)
	}

	ln = applyThirdRequestToDraftLine(in.GetDraftLine(), ln)
	err = storeql.UpdateIntoDB(ctx, db, ln)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "UpdateIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func applyThirdRequestToDraftLine(in *pbproducts.DraftLineThird, ln *ent.DraftLine) *ent.DraftLine {
	return ln
}
