package order

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/jmoiron/sqlx"
)

func (s *Shipping) StateMachine() *sm.StateMachine {
	return sm.ShippingStateMachine
}

func (s *Shipping) StateChange() sm.StateChange { return &ShippingStateChange{} }

func (s *Shipping) ValidateStateChange(ctx context.Context, db *sqlx.DB, newState *sm.State) (err error) {
	switch newState.Name {
	}
	return
}
