package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/messager/ent"
	"github.com/athomecomar/athome/pb/pbmessager"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveMany(ctx context.Context, in *pbmessager.RetrieveManyRequest) (*pbmessager.RetrieveManyResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	userId, err := pbutil.GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	err = authCloser()
	if err != nil {
		return nil, err
	}

	return s.retrieveMany(ctx, db, userId)
}

func (s *Server) retrieveMany(ctx context.Context, db *sqlx.DB, userId uint64) (*pbmessager.RetrieveManyResponse, error) {
	notifs, err := ent.FindMessagesByUser(ctx, db, userId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "you dont have messagcations")
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindMessage: %v", err)
	}

	resp := &pbmessager.RetrieveManyResponse{}
	resp.Messages = make(map[uint64]*pbmessager.Message)
	for _, notif := range notifs {
		notifPb, err := notif.ToPb()
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "notif.ToPb: %v", err)
		}
		resp.Messages[notif.Id] = notifPb
	}
	return resp, nil
}
