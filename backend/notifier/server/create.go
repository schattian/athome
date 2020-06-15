package server

import (
	"context"
	"time"

	"github.com/athomecomar/athome/backend/notifier/ent"
	"github.com/athomecomar/athome/backend/notifier/notifierconf"
	"github.com/athomecomar/athome/pb/pbnotifier"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) Create(ctx context.Context, in *pbnotifier.CreateRequest) (*pbnotifier.CreateResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return s.create(ctx, db, in)
}

func (s *Server) create(ctx context.Context, db *sqlx.DB, in *pbnotifier.CreateRequest) (*pbnotifier.CreateResponse, error) {
	claims, err := claimJwt(in.GetNotificationToken(), notifierconf.GetNOTIFICATION_JWT_SECRET)
	if err != nil {
		return nil, err
	}
	userId, err := userIdFromClaim(claims)
	if err != nil {
		return nil, err
	}
	notif, err := ent.NotificationFromPb(in.GetNotification())
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "NotificationFromPb: %v", err)
	}
	notif.CreatedAt = time.Now()
	notif.UserId = userId
	err = storeql.InsertIntoDB(ctx, db, notif)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}
	pbNotif, err := notif.ToPb()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "notif.ToPb: %v", err)
	}

	return &pbnotifier.CreateResponse{
			NotificationId: notif.Id,
			Notification:   pbNotif,
		},
		nil
}
