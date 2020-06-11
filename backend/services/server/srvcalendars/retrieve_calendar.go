package srvcalendars

import (
	"context"

	"github.com/athomecomar/athome/backend/services/pb/pbservices"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/jmoiron/sqlx"
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
	return server.RetrieveCalendarDetail(ctx, db, in.GetCalendarId())
}
