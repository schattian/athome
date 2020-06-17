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

func (s *Server) RetrieveMany(ctx context.Context, in *pbnotifier.RetrieveManyRequest) (*pbnotifier.RetrieveManyResponse, error) {
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

func (s *Server) retrieveMany(ctx context.Context, db *sqlx.DB, userId uint64) (*pbnotifier.RetrieveManyResponse, error) {
	notifs, err := ent.FindNotificationsByUser(ctx, db, userId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "you dont have notifications")
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindNotification: %v", err)
	}

	resp := &pbnotifier.RetrieveManyResponse{}
	resp.Notifications = make(map[uint64]*pbnotifier.Notification)
	for _, notif := range notifs {
		notifPb, err := notif.ToPb()
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "notif.ToPb: %v", err)
		}
		resp.Notifications[notif.Id] = notifPb
	}
	return resp, nil
}
