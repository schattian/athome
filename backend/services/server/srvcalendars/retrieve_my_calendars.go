package srvcalendars

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/pb/pbauth"
	"github.com/athomecomar/athome/backend/services/pb/pbservices"
	"github.com/athomecomar/athome/backend/services/server"
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
	auth, authCloser, err := server.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	return s.retrieveMyCalendars(ctx, db, auth, server.GetUserFromAccessToken, in)
}

func (s *Server) retrieveMyCalendars(
	ctx context.Context,
	db *sqlx.DB,
	auth pbauth.AuthClient,
	authFn server.AuthFunc,
	in *pbservices.RetrieveMyCalendarsRequest,
) (*pbservices.RetrieveMyCalendarsResponse, error) {
	userId, err := authFn(ctx, auth, in.GetAccessToken())
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
	resp.Calendars = make(map[uint64]*pbservices.CalendarDetail)
	for _, c := range cs {
		avs, err := c.Availabilities(ctx, db)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "Availabilities: %v", err)
		}
		calDetail := &pbservices.CalendarDetail{Calendar: c.ToPb()}
		calDetail.Availabilities = make(map[uint64]*pbservices.Availability)
		for _, av := range avs {
			calDetail.Availabilities[av.Id] = av.ToPb()
		}
		resp.Calendars[c.Id] = calDetail
	}
	return resp, nil
}
