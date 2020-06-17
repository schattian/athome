package server

import (
	"context"

	"github.com/athomecomar/athome/backend/messager/ent"
	"github.com/athomecomar/athome/backend/messager/messagerconf"
	"github.com/athomecomar/athome/pb/pbmessager"
	"github.com/athomecomar/athome/pb/pbnotifier"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) Reply(ctx context.Context, in *pbmessager.ReplyRequest) (*pbmessager.CreateResponse, error) {
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
	noti, notiCloser, err := pbutil.ConnNotifier(ctx)
	if err != nil {
		return nil, err
	}
	defer notiCloser()

	return s.reply(ctx, db, noti, in, userId)
}

func (s *Server) reply(ctx context.Context, db *sqlx.DB, noti pbnotifier.NotificationsClient, in *pbmessager.ReplyRequest, uid uint64) (*pbmessager.CreateResponse, error) {
	repliedMsg, err := ent.FindMessageByReceiver(ctx, db, in.GetMessageId(), uid)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "ent.FindMessageByReceiver: %v", err)
	}
	msg := ent.NewMessage(uid, repliedMsg.SenderId, in.GetBody())

	err = storeql.InsertIntoDB(ctx, db, msg)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}

	_, err = pbutil.CreateNotification(
		ctx,
		msg,
		repliedMsg.ReceiverId,
		messagerconf.GetNOTIFICATION_JWT_SECRET,
		messagerconf.GetNOTIFICATION_BODY(),
		pbutil.High,
		noti,
	)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "msg.ToPb: %v", err)
	}

	pbMsg, err := msg.ToPb()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "msg.ToPb: %v", err)
	}
	return &pbmessager.CreateResponse{
		MessageId: msg.Id,
		Message:   pbMsg,
	}, nil
}
