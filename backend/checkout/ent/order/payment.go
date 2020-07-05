package order

import (
	"context"
	"database/sql"
	"time"

	"github.com/athomecomar/athome/backend/checkout/ent"
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
	EntityTable     Class        `json:"entity_table,omitempty"`
	Amount          currency.ARS `json:"amount,omitempty"`
	CreatedAt       ent.Time     `json:"created_at,omitempty"`
	UpdatedAt       ent.Time     `json:"updated_at,omitempty"`
	Installments    uint64       `json:"installments,omitempty"`
}

func (o *Payment) GetCreatedAt() time.Time { return o.CreatedAt.Time }
func (o *Payment) GetUpdatedAt() time.Time { return o.UpdatedAt.Time }
func (o *Payment) SetCreatedAt(t time.Time) {
	o.CreatedAt = ent.Time{NullTime: sql.NullTime{Time: t}}
	o.SetUpdatedAt(t)
}
func (o *Payment) SetUpdatedAt(t time.Time) { o.UpdatedAt = ent.Time{NullTime: sql.NullTime{Time: t}} }

func (p *Payment) Card(ctx context.Context, db *sqlx.DB) (*Card, error) {
	card, err := FindCard(ctx, db, p.CardId, p.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "FindCard")
	}
	return card, nil
}

func (p *Payment) Order(ctx context.Context, db *sqlx.DB) (o Order, err error) {
	switch p.EntityTable {
	case Purchases:
		o, err = p.Purchase(ctx, db)
	case "reservations":
	case "bookings":
	}
	return
}

func (p *Payment) Purchase(ctx context.Context, db *sqlx.DB) (*Purchase, error) {
	if p.EntityTable != Purchases {
		return nil, errors.New("payment's entity isn't a purchase")
	}
	ord, err := FindPurchase(ctx, db, p.EntityId)
	if err != nil {
		return nil, errors.Wrap(err, "FindPurchase")
	}
	return ord, nil
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
