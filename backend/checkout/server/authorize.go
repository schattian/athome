package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func FindPurchase(ctx context.Context, db *sqlx.DB, oId, uId uint64) (*order.Purchase, error) {
	order, err := order.FindPurchase(ctx, db, oId, uId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "FindPurchase with order_id: %v", oId)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindPurchase")
	}
	return order, nil
}

func FindLatestPurchase(ctx context.Context, db *sqlx.DB, uId uint64) (*order.Purchase, error) {
	order, err := order.FindLatestPurchase(ctx, db, uId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(xerrors.NotFound, "FindCurrentPurchase")
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindPurchase")
	}
	return order, nil
}
