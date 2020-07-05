package order

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/jmoiron/sqlx"
)

type Order interface {
	sm.Stateful

	OrderClass() Class
	CanView(ctx context.Context, db *sqlx.DB, userId uint64) (bool, error)
}
