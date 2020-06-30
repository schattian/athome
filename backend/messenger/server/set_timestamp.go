package server

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/athomecomar/athome/backend/messenger/ent"
	"github.com/athomecomar/athome/pb/pbmessenger"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SetReceived(ctx context.Context, in *pbmessenger.UpdateStatusRequest) (*emptypb.Empty, error) {
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

	return s.setTimestamp(ctx, db, userId, in, setReceived)
}

func (s *Server) SetSeen(ctx context.Context, in *pbmessenger.UpdateStatusRequest) (*emptypb.Empty, error) {
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

	return s.setTimestamp(ctx, db, userId, in, setSeen)
}

func (s *Server) setTimestamp(ctx context.Context, db *sqlx.DB, userId uint64, in *pbmessenger.UpdateStatusRequest, setterFunc func(*ent.Message) *ent.Message) (*emptypb.Empty, error) {
	msg, err := ent.FindMessageByReceiver(ctx, db, in.GetMessageId(), userId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "FindMessageByReceiver with id: %v", in.GetMessageId())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindMessage: %v", err)
	}
	msg = setterFunc(msg)
	err = storeql.UpdateIntoDB(ctx, db, msg)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.UpdateIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func setReceived(n *ent.Message) *ent.Message {
	n.ReceivedAt = ent.Time{NullTime: sql.NullTime{Time: time.Now()}}
	return n
}

func setSeen(n *ent.Message) *ent.Message {
	n.SeenAt = ent.Time{NullTime: sql.NullTime{Time: time.Now()}}
	return n
}
