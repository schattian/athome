package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/notifier/ent"
	"github.com/athomecomar/athome/pb/pbnotifier"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) Retrieve(ctx context.Context, in *pbnotifier.RetrieveRequest) (*pbnotifier.Notification, error) {
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

func (s *Server) retrieve(ctx context.Context, db *sqlx.DB, userId uint64, in *pbnotifier.RetrieveRequest) (*pbnotifier.Notification, error) {
	notif, err := ent.FindNotification(ctx, db, in.GetNotificationId(), userId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "FindNotification with id: %v", in.GetNotificationId())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindNotification: %v", err)
	}
	resp, err := notif.ToPb()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "notif.ToPb: %v", err)
	}
	return resp, nil
}
