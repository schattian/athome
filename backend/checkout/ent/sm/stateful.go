package sm

import (
	"context"

	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
)

type Stateful interface {
	storeql.Storable

	StateMachine() *StateMachine

	GetConsumerId() uint64
	GetMerchantId() uint64
	GetServiceProviderId() uint64

	ValidateStateChange(ctx context.Context, db *sqlx.DB, state *State) error
}
