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

func (s *Server) Prev(ctx context.Context, in *pbproducts.StageChangeRequest) (*pbproducts.StageChangeResponse, error) {
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

	return s.prev(ctx, db, in, draft)
}

func (s *Server) prev(ctx context.Context, db *sqlx.DB, in *pbproducts.StageChangeRequest, d *ent.Draft) (*pbproducts.StageChangeResponse, error) {
	err := d.Prev(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "draft.Prev: %v", err)
	}

	return nil, nil
}
