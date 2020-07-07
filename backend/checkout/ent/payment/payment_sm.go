package payment

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/jmoiron/sqlx"
)

func (p *Payment) GetMerchantId() uint64        { return 0 }
func (p *Payment) GetConsumerId() uint64        { return p.UserId }
func (p *Payment) GetServiceProviderId() uint64 { return 0 }

func (p *Payment) StateMachine() *sm.StateMachine {
	return sm.PaymentStateMachine
}

func (p *Payment) ValidateStateChange(ctx context.Context, db *sqlx.DB, newState *sm.State) (err error) {
	switch newState.Name {
	}
	return
}
