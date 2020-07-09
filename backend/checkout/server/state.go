package server

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/storeql"
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

func ChangeState(
	ctx context.Context,
	db *sqlx.DB,
	stateChanger sm.StateChanger,
	s sm.Stateful,
	uid uint64,
) error {
	sc, err := sm.LatestStateChange(ctx, db, s)
	if err != nil {
		return status.Errorf(xerrors.Internal, "LatestStateChange")
	}
	state, err := stateChanger(s.StateMachine(), sc.GetState(s.StateMachine()), s, uid)
	if err != nil {
		return status.Errorf(xerrors.InvalidArgument, "sm stateChanger: %v", err)
	}
	err = s.ValidateStateChange(ctx, db, state)
	if err != nil {
		return status.Errorf(xerrors.InvalidArgument, "ValidateStateChange: %v", err)
	}
	sc, err = sm.NewStateChange(ctx, s.GetId(), state.Name, s)
	if err != nil {
		return status.Errorf(xerrors.Internal, "NewShippingStateChange: %v", err)
	}
	err = storeql.InsertIntoDB(ctx, db, sc)
	if err != nil {
		return status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}
	return nil
}
