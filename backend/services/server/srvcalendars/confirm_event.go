package srvcalendars

import (
	"context"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ConfirmEvent(ctx context.Context, in *pbservices.ConfirmEventRequest) (*emptypb.Empty, error) {
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
	defer authCloser()
	return s.confirmEvent(ctx, db, nil, server.GetUserFromAccessToken(auth, in.GetAccessToken()), in)
}

func (s *Server) confirmEvent(
	ctx context.Context,
	db *sqlx.DB,
	checkout pbcheckout.BookingsClient,
	authFn server.AuthFunc,
	in *pbservices.ConfirmEventRequest,
) (*emptypb.Empty, error) {
	claimantId, err := authFn(ctx)
	if err != nil {
		return nil, err
	}
	ev, err := ent.FindEvent(ctx, db, in.GetEventId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindEvent: %v", err)
	}
	if ev.ClaimantId != claimantId {
		return nil, status.Errorf(xerrors.PermissionDenied, "this event does not belong to you")
	}
	ev.IsConfirmed = true
	err = storeql.UpdateIntoDB(ctx, db, ev)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "UpdateIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}
