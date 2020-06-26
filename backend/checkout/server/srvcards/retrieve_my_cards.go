package srvcards

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveMyCards(ctx context.Context, in *pbcheckout.RetrieveMyCardsRequest) (*pbcheckout.RetrieveMyCardsResponse, error) {
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

	return s.retrieveMyCards(ctx, db, uid)
}

func (s *Server) retrieveMyCards(ctx context.Context, db *sqlx.DB, uid uint64) (*pbcheckout.RetrieveMyCardsResponse, error) {
	cards, err := order.FindUserCards(ctx, db, uid)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindUserCards: %v", err)
	}
	resp := &pbcheckout.RetrieveMyCardsResponse{}
	resp.Cards = make(map[uint64]*pbcheckout.Card)
	for _, card := range cards {
		resp.Cards[card.Id] = card.ToPb()
	}
	return resp, nil
}
