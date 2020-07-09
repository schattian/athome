package payment

import (
	"context"
	"database/sql"
	"time"

	"github.com/athomecomar/athome/backend/checkout/ent"
	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/currency"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Payment struct {
	Id              uint64       `json:"id,omitempty"`
	UserId          uint64       `json:"user_id,omitempty"`
	PaymentMethodId uint64       `json:"payment_method_id,omitempty"`
	CardId          uint64       `json:"card_id,omitempty"`
	EntityId        uint64       `json:"entity_id,omitempty"`
	EntityTable     order.Class  `json:"entity_table,omitempty"`
	Amount          currency.ARS `json:"amount,omitempty"`
	CreatedAt       ent.Time     `json:"created_at,omitempty"`
	UpdatedAt       ent.Time     `json:"updated_at,omitempty"`
	Installments    uint64       `json:"installments,omitempty"`
}

func (py *Payment) GetOrderId() uint64         { return py.EntityId }
func (py *Payment) GetOrderClass() order.Class { return py.EntityTable }

func (o *Payment) GetCreatedAt() time.Time { return o.CreatedAt.Time }
func (o *Payment) GetUpdatedAt() time.Time { return o.UpdatedAt.Time }
func (o *Payment) SetCreatedAt(t time.Time) {
	o.CreatedAt = ent.Time{NullTime: sql.NullTime{Time: t}}
	o.SetUpdatedAt(t)
}
func (o *Payment) SetUpdatedAt(t time.Time) { o.UpdatedAt = ent.Time{NullTime: sql.NullTime{Time: t}} }

func (p *Payment) IsFinished(ctx context.Context, db *sqlx.DB) (bool, error) {
	sc, err := sm.LatestStateChange(ctx, db, p)
	if err != nil {
		return false, errors.Wrap(err, "LatestStateChange")
	}
	return sc.GetName() == sm.PaymentFinished, nil
}

func (p *Payment) Card(ctx context.Context, db *sqlx.DB) (*Card, error) {
	card, err := FindCard(ctx, db, p.CardId, p.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "FindCard")
	}
	return card, nil
}

func FindPayment(ctx context.Context, db *sqlx.DB, oId uint64, userId uint64) (*Payment, error) {
	order := &Payment{}
	row := storeql.Where(ctx, db, order, `id=$1 AND user_id=$2`, oId, userId)
	err := row.StructScan(order)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return order, nil
}

func PaymentFromPb(in *pbcheckout.PaymentInput) *Payment {
	return &Payment{
		PaymentMethodId: in.GetPaymentMethodId(),
		Installments:    in.GetInstallments(),
		CardId:          in.GetCardId(),
		EntityId:        in.Order.GetEntityId(),
		EntityTable:     order.Class(in.Order.GetEntityTable()),
	}
}

func (p *Payment) ToPb() (*pbcheckout.Payment, error) {
	ts, err := ent.GetTimestamp(p)
	if err != nil {
		return nil, errors.Wrap(err, "GetTimestamp")
	}
	return &pbcheckout.Payment{
		UserId:          p.UserId,
		CardId:          p.CardId,
		Amount:          p.Amount.Float64(),
		Timestamp:       ts,
		PaymentMethodId: p.PaymentMethodId,
		Installments:    p.Installments,
	}, nil
}
