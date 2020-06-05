package srvcreator

import (
	"context"
	"io"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/athome/backend/products/pb/pbproducts"
	"github.com/athomecomar/athome/backend/products/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Second(srv pbproducts.Creator_SecondServer) error {
	ctx := srv.Context()

	db, err := server.ConnDB()
	if err != nil {
		return err
	}

	sem, semCloser, err := server.ConnSemantic(ctx)
	if err != nil {
		return err
	}
	defer semCloser()

	var resp *emptypb.Empty
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
			draft, err = server.FetchLatestDraft(ctx, db, in.GetAccessToken())
			if err != nil {
				return err
			}
		}

		if draft.Stage != stage.Second {
			return status.Errorf(xerrors.InvalidArgument, "stage expected %v, got %v", stage.Second, draft.Stage)
		}

		resp, err = s.second(ctx, db, sem, in, draft)
		if err != nil {
			return err
		}
	}
}

func (s *Server) second(ctx context.Context, db *sqlx.DB, sem pbsemantic.ProductsClient, in *pbproducts.SecondRequest, draft *ent.Draft) (*emptypb.Empty, error) {
	// sem.GetAttributesSchema(ctx, &pbsemantic.GetAttributesSchemaRequest{CategoryId: }, opts ...grpc.CallOption)
	return nil, nil
}
