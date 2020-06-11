package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/pb/pbservices"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func RetrieveCalendarDetail(
	ctx context.Context,
	db *sqlx.DB,
	id uint64,
) (*pbservices.CalendarDetail, error) {
	c, err := ent.FindCalendar(ctx, db, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "can't find calendar with id: %v", id)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindCalendar: %v", err)
	}
	detail, err := c.Detail(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Detail: %v", err)
	}
	return detail, nil
}
