package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

type Server struct{}

const dispatchDelayMinutes = 30 // TODO: add by-user time

func mustPrevState(ctx context.Context, db *sqlx.DB, o *order.Purchase, desired sm.StateName) error {
	sc, err := o.State(ctx, db)
	if err != nil {
		return status.Errorf(xerrors.Internal, "State: %v", err)
	}
	s, err := sm.Next(o.StateMachine(), sc.GetState())
	if err != nil {
		return status.Errorf(xerrors.OutOfRange, "GetState: %v", err)
	}
	if s.Name != desired {
		return status.Errorf(xerrors.InvalidArgument, "desired state: %s got: %s", desired, sc.Name)
	}
	return nil
}
