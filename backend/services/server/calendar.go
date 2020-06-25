package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func RetrieveCalendar(
	ctx context.Context,
	db *sqlx.DB,
	id uint64,
) (*pbservices.Calendar, error) {
	c, err := ent.FindCalendar(ctx, db, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "can't find calendar with id: %v", id)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindCalendar: %v", err)
	}
	return c.ToPb(), nil
}
