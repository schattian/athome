package order

import (
	"context"
	"database/sql"
	"time"

	"github.com/athomecomar/athome/backend/checkout/ent"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/pkg/errors"
)

type PurchaseStateChange struct {
	Id        uint64 `json:"id,omitempty"`
	OrderId   uint64 `json:"order_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Stage     uint64
	CreatedAt ent.Time `json:"created_at,omitempty"`
}

func (o *PurchaseStateChange) GetName() string         { return o.Name }
func (o *PurchaseStateChange) GetStage() uint64        { return o.Stage }
func (o *PurchaseStateChange) GetCreatedAt() time.Time { return o.CreatedAt.Time }
func (o *PurchaseStateChange) SetCreatedAt(t time.Time) {
	o.CreatedAt = ent.Time{NullTime: sql.NullTime{Time: t}}
}

func NewPurchaseStateChange(ctx context.Context, oId uint64) (*PurchaseStateChange, error) {
	state := sm.PurchaseStateMachine.First()
	if state == nil {
		return nil, errors.New("stage doesn't exists")
	}
	p := &PurchaseStateChange{
		OrderId: oId,
		Stage:   0,
		Name:    state.Name,
	}
	p.SetCreatedAt(time.Now())
	return p, nil
}
