package srvcalendars

import (
	"context"

	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/jmoiron/sqlx"
)

func (s *Server) RetrieveCalendar(ctx context.Context, in *pbservices.RetrieveCalendarRequest) (*pbservices.Calendar, error) {
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
) (*pbservices.Calendar, error) {
	return server.RetrieveCalendar(ctx, db, in.GetCalendarId())
}
