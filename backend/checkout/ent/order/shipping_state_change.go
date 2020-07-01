package order

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/athomecomar/athome/backend/checkout/ent"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
)

type ShippingStateChange struct {
	Id         uint64       `json:"id,omitempty"`
	ShippingId uint64       `json:"shipping_id,omitempty"`
	Name       sm.StateName `json:"name,omitempty"`
	Stage      int64
	CreatedAt  ent.Time `json:"created_at,omitempty"`
}

func (o *ShippingStateChange) GetName() string         { return string(o.Name) }
func (o *ShippingStateChange) GetStage() int64         { return o.Stage }
func (o *ShippingStateChange) GetCreatedAt() time.Time { return o.CreatedAt.Time }
func (o *ShippingStateChange) SetCreatedAt(t time.Time) {
	o.CreatedAt = ent.Time{NullTime: sql.NullTime{Time: t}}
}

func (o *ShippingStateChange) GetState() *sm.State {
	return sm.ShippingStateMachine.StateByStage(o.Stage)
}

func NewShippingStateChange(ctx context.Context, sId uint64, stateName sm.StateName) (*ShippingStateChange, error) {
	state := sm.PurchaseStateMachine.StateByName(stateName)
	if state == nil {
		return nil, fmt.Errorf("state named %s doesn't exists", stateName)
	}
	stage := sm.PurchaseStateMachine.StageByName(stateName)
	p := &ShippingStateChange{
		ShippingId: sId,
		Stage:      stage,
		Name:       state.Name,
	}
	p.SetCreatedAt(time.Now())
	return p, nil
}
