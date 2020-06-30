package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/messenger/ent"
	"github.com/athomecomar/athome/pb/pbmessenger"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) Retrieve(ctx context.Context, in *pbmessenger.RetrieveRequest) (*pbmessenger.Message, error) {
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

	return s.retrieve(ctx, db, userId, in)
}

func (s *Server) retrieve(ctx context.Context, db *sqlx.DB, userId uint64, in *pbmessenger.RetrieveRequest) (*pbmessenger.Message, error) {
	notif, err := ent.FindMessage(ctx, db, in.GetMessageId(), userId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "FindMessage with id: %v", in.GetMessageId())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindMessage: %v", err)
	}
	resp, err := notif.ToPb()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "notif.ToPb: %v", err)
	}
	return resp, nil
}
