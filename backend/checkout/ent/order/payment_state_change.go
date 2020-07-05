package order

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/athomecomar/athome/backend/checkout/ent"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
)

type PaymentStateChange struct {
	Id        uint64 `json:"id,omitempty"`
	EntityId  uint64
	Name      sm.StateName `json:"name,omitempty"`
	Stage     int64
	CreatedAt ent.Time `json:"created_at,omitempty"`
}

func (o *PaymentStateChange) GetName() string         { return string(o.Name) }
func (o *PaymentStateChange) GetStage() int64         { return o.Stage }
func (o *PaymentStateChange) GetCreatedAt() time.Time { return o.CreatedAt.Time }
func (o *PaymentStateChange) SetCreatedAt(t time.Time) {
	o.CreatedAt = ent.Time{NullTime: sql.NullTime{Time: t}}
}

func (o *PaymentStateChange) GetState() *sm.State {
	return sm.ShippingStateMachine.StateByStage(o.Stage)
}

func NewPaymentStateChange(ctx context.Context, sId uint64, stateName sm.StateName) (*PaymentStateChange, error) {
	state := sm.PurchaseStateMachine.StateByName(stateName)
	if state == nil {
		return nil, fmt.Errorf("state named %s doesn't exists", stateName)
	}
	stage := sm.PurchaseStateMachine.StageByName(stateName)
	p := &PaymentStateChange{
		EntityId: sId,
		Stage:    stage,
		Name:     state.Name,
	}
	p.SetCreatedAt(time.Now())
	return p, nil
}
