package srvcalendars

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveMyCalendars(ctx context.Context, in *pbservices.RetrieveMyCalendarsRequest) (*pbservices.RetrieveMyCalendarsResponse, error) {
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

	return s.retrieveMyCalendars(ctx, db, server.GetUserFromAccessToken(auth, in.GetAccessToken()))
}

func (s *Server) retrieveMyCalendars(
	ctx context.Context,
	db *sqlx.DB,
	authFn server.AuthFunc,
) (*pbservices.RetrieveMyCalendarsResponse, error) {
	userId, err := authFn(ctx)
	if err != nil {
		return nil, err
	}
	cs, err := ent.CalendarsByUserId(ctx, db, userId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "can't find calendars for user with id: %v", userId)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindCalendar: %v", err)
	}
	resp := &pbservices.RetrieveMyCalendarsResponse{}
	resp.Calendars = make(map[uint64]*pbservices.Calendar)
	for _, c := range cs {
		resp.Calendars[c.Id] = c.ToPb()
	}
	return resp, nil
}
