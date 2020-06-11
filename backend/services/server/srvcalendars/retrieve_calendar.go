package srvcalendars

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/pb/pbservices"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveCalendar(ctx context.Context, in *pbservices.RetrieveCalendarRequest) (*pbservices.CalendarDetail, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.retrieveCalendar(ctx, db, in)
}

func (s *Server) retrieveCalendar(
	ctx context.Context,
	db *sqlx.DB,
	in *pbservices.RetrieveCalendarRequest,
) (*pbservices.CalendarDetail, error) {
	c, err := ent.FindCalendar(ctx, db, in.GetCalendarId())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "can't find calendar with id: %v", in.GetCalendarId())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindCalendar: %v", err)
	}
	avs, err := c.Availabilities(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Availabilities: %v", err)
	}
	resp := &pbservices.CalendarDetail{Calendar: c.ToPb()}
	resp.Availabilities = make(map[uint64]*pbservices.Availability)
	for _, av := range avs {
		resp.Availabilities[av.Id] = av.ToPb()
	}
	return resp, nil
}
