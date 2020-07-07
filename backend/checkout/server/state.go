package server

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func MustPrevState(ctx context.Context, db *sqlx.DB, stateful sm.Stateful, desired sm.StateName, uid uint64) error {
	sc, err := sm.LatestStateChange(ctx, db, stateful)
	if err != nil {
		return status.Errorf(xerrors.Internal, "State: %v", err)
	}
	s, err := sm.Next(stateful.StateMachine(), sc.GetState(stateful.StateMachine()), stateful, uid)
	if err != nil {
		return status.Errorf(xerrors.OutOfRange, "GetState: %v", err)
	}
	if s.Name != desired {
		return status.Errorf(xerrors.InvalidArgument, "desired state: %s got: %s", desired, sc.GetName())
	}
	return nil
}
