package server

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/athomecomar/athome/backend/notifier/ent"
	"github.com/athomecomar/athome/pb/pbnotifier"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveStream(in *pbnotifier.RetrieveStreamRequest, srv pbnotifier.Notifications_RetrieveStreamServer) error {
	if err := in.Validate(); err != nil {
		return err
	}
	db, err := ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()
	auth, authCloser, err := pbutil.ConnAuth(srv.Context())
	if err != nil {
		return err
	}

	userId, err := pbutil.GetUserFromAccessToken(srv.Context(), auth, in.GetAccessToken())
	if err != nil {
		return err
	}
	authCloser()
	ticker := time.Duration(in.GetTickerMs()) * time.Millisecond

	var offset time.Time
	for {
		select {
		case <-srv.Context().Done():
			return srv.Context().Err()
		default:
			time.Sleep(ticker)
			resp, err := s.retrieveStream(srv.Context(), db, userId, offset)
			if err != nil {
				return err
			}
			for _, notif := range resp.GetNotifications() {
				err = srv.Send(notif)
				if err != nil {
					return err
				}
			}
			offset = time.Now()
		}
	}

}

func (s *Server) retrieveStream(ctx context.Context, db *sqlx.DB, userId uint64, offset time.Time) (*pbnotifier.RetrieveManyResponse, error) {
	notifs, err := ent.FindNotificationsByUserWithOffset(ctx, db, userId, offset)
	if errors.Is(err, sql.ErrNoRows) {
		return &pbnotifier.RetrieveManyResponse{}, nil
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
