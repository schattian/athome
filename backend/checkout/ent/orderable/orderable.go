package orderable

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/jmoiron/sqlx"
)

type Orderable interface {
	GetOrderId() uint64
	GetOrderClass() order.Class
}

func FromOrderable(ctx context.Context, db *sqlx.DB, orderable Orderable) (o order.Order, err error) {
	switch orderable.GetOrderClass() {
	case order.Purchases:
		o, err = purchase.FindPurchase(ctx, db, orderable.GetOrderId())
	case order.Reservations:
	case order.Bookings:
	}
	return
}
