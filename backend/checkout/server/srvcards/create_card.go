package srvcards

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCard(ctx context.Context, in *pbcheckout.CreateCardRequest) (*pbcheckout.CreateCardResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	uid, err := pbutil.GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	err = authCloser()
	if err != nil {
		return nil, err
	}

	return s.createCard(ctx, db, in, uid)
}

func (s *Server) createCard(ctx context.Context, db *sqlx.DB, in *pbcheckout.CreateCardRequest, userId uint64) (*pbcheckout.CreateCardResponse, error) {
	card, err := order.NewCardFromPb(ctx, in.GetCard(), userId)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "NewCardFromPb: %v", err)
	}
	err = storeql.InsertIntoDB(ctx, db, card)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
	}

	return &pbcheckout.CreateCardResponse{
		CardId: card.Id,
		Card:   card.ToPb(),
	}, nil
}
