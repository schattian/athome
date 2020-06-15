package server

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/athomecomar/athome/backend/notifier/ent"
	"github.com/athomecomar/athome/pb/pbnotifier"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SetReceived(ctx context.Context, in *pbnotifier.UpdateStatusRequest) (*emptypb.Empty, error) {
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
	authCloser()

	return s.setTimestamp(ctx, db, userId, in, setReceived)
}

func (s *Server) SetSeen(ctx context.Context, in *pbnotifier.UpdateStatusRequest) (*emptypb.Empty, error) {
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
	authCloser()

	return s.setTimestamp(ctx, db, userId, in, setSeen)
}

func (s *Server) setTimestamp(ctx context.Context, db *sqlx.DB, userId uint64, in *pbnotifier.UpdateStatusRequest, setterFunc func(*ent.Notification) *ent.Notification) (*emptypb.Empty, error) {
	notif, err := ent.FindNotification(ctx, db, in.GetNotificationId(), userId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "FindNotification with id: %v", in.GetNotificationId())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindNotification: %v", err)
	}
	notif = setterFunc(notif)
	err = storeql.UpdateIntoDB(ctx, db, notif)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.UpdateIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func setReceived(n *ent.Notification) *ent.Notification {
	n.ReceivedAt = ent.Time{NullTime: sql.NullTime{Time: time.Now()}}
	return n
}

func setSeen(n *ent.Notification) *ent.Notification {
	n.SeenAt = ent.Time{NullTime: sql.NullTime{Time: time.Now()}}
	return n
}
