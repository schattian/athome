package order

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/athomecomar/athome/backend/checkout/ent"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
)

type PurchaseStateChange struct {
	Id        uint64       `json:"id,omitempty"`
	OrderId   uint64       `json:"order_id,omitempty"`
	Name      sm.StateName `json:"name,omitempty"`
	Stage     int64
	CreatedAt ent.Time `json:"created_at,omitempty"`
}

func (o *PurchaseStateChange) GetName() string         { return string(o.Name) }
func (o *PurchaseStateChange) GetStage() int64         { return o.Stage }
func (o *PurchaseStateChange) GetCreatedAt() time.Time { return o.CreatedAt.Time }
func (o *PurchaseStateChange) SetCreatedAt(t time.Time) {
	o.CreatedAt = ent.Time{NullTime: sql.NullTime{Time: t}}
}

func (o *PurchaseStateChange) GetState() *sm.State {
	return sm.PurchaseStateMachine.StateByStage(o.Stage)
}

func NewPurchaseStateChange(ctx context.Context, oId uint64, stateName sm.StateName) (*PurchaseStateChange, error) {
	state := sm.PurchaseStateMachine.StateByName(stateName)
	if state == nil {
		return nil, fmt.Errorf("state named %s doesn't exists", stateName)
	}
	p := &PurchaseStateChange{
		OrderId: oId,
		Stage:   0,
		Name:    state.Name,
	}
	p.SetCreatedAt(time.Now())
	return p, nil
}
