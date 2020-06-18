package order

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/athomecomar/athome/backend/checkout/ent"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Purchase struct {
	Id        uint64            `json:"id,omitempty"`
	UserId    uint64            `json:"user_id,omitempty"`
	CreatedAt ent.Time          `json:"created_at,omitempty"`
	UpdatedAt ent.Time          `json:"updated_at,omitempty"`
	Items     map[uint64]uint64 `json:"items,omitempty"`
}

func (o *Purchase) GetCreatedAt() time.Time  { return o.CreatedAt.Time }
func (o *Purchase) GetUpdatedAt() time.Time  { return o.UpdatedAt.Time }
func (o *Purchase) SetCreatedAt(t time.Time) { o.CreatedAt = ent.Time{NullTime: sql.NullTime{Time: t}} }
func (o *Purchase) SetUpdatedAt(t time.Time) { o.UpdatedAt = ent.Time{NullTime: sql.NullTime{Time: t}} }

func FindPurchase(ctx context.Context, db *sqlx.DB, id uint64) (*Purchase, error) {
	order := &Purchase{}
	row := storeql.Where(ctx, db, order, "id=$1", id)
	err := row.StructScan(order)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return order, nil
}

func FindLatestPurchase(ctx context.Context, db *sqlx.DB, userId uint64) (*Purchase, error) {
	row := db.QueryRowxContext(ctx,
		`SELECT * FROM purchases WHERE id=(
            SELECT order_id FROM purchase_state_changes WHERE order_id IN (
                SELECT id FROM purchases WHERE user_id = $1
            )
             ORDER BY stage ASC, created_at DESC
        )`,
		userId)
	order := &Purchase{}
	err := row.StructScan(order)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return order, nil
}

// Amount               float64                 `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
func (o *Purchase) OrderClass() class {
	return Purchases
}

func (o *Purchase) Products(ctx context.Context, c pbproducts.ViewerClient) (map[uint64]*pbproducts.Product, error) {
	prods, err := c.RetrieveProducts(ctx, &pbproducts.RetrieveProductsRequest{Ids: o.ProductIds()})
	if err != nil {
		return nil, errors.Wrap(err, "RetrieveProducts")
	}
	return prods.Products, nil
}

func (o *Purchase) ProductIds() (ids []uint64) {
	for pid := range o.Items {
		ids = append(ids, pid)
	}
	return
}

func (o *Purchase) Amount(ctx context.Context, c pbproducts.ViewerClient) (float64, error) {
	prods, err := o.Products(ctx, c)
	if err != nil {
		return 0, errors.Wrap(err, "Products")
	}
	return o.AmountFromProducts(ctx, prods), nil
}

func (o *Purchase) AmountFromProducts(ctx context.Context, prods map[uint64]*pbproducts.Product) float64 {
	var amount float64
	for id, prod := range prods {
		amount += float64(o.Items[id]) * prod.GetPrice()
	}
	return amount
}

func (o *Purchase) ToPbWrapped(ctx context.Context, db *sqlx.DB, prods pbproducts.ViewerClient) (*pbcheckout.Purchase, error) {
	scs, err := StateChanges(ctx, db, o)
	if err != nil {
		return nil, errors.Wrap(err, "StateChanges")
	}
	amount, err := o.Amount(ctx, prods)
	if err != nil {
		return nil, errors.Wrap(err, "Amount")
	}
	pb, err := o.ToPb(scs, amount)
	if err != nil {
		return nil, errors.Wrap(err, "ToPb")
	}
	return pb, nil
}

func (o *Purchase) ValidateStock(ctx context.Context, prods map[uint64]*pbproducts.Product) error {
	for id, prod := range prods {
		if prod.GetStock() < o.Items[id] {
			return fmt.Errorf("product with id %v has %d of stock while asked for %d", id, prod.GetStock(), o.Items[id])
		}
	}
	return nil
}

func NewPurchase(ctx context.Context, items map[uint64]uint64, uid uint64) *Purchase {
	p := &Purchase{UserId: uid, Items: items}
	p.SetCreatedAt(time.Now())
	return p
}

func (o *Purchase) StateMachine() *sm.StateMachine {
	return sm.PurchaseStateMachine
}

func (o *Purchase) ToPb(scs []StateChange, amount float64) (*pbcheckout.Purchase, error) {
	ts, err := ent.GetTimestamp(o)
	if err != nil {
		return nil, errors.Wrap(err, "GetTimestamp")
	}

	pbScs := make(map[uint64]*pbcheckout.StateChange)
	for _, sc := range scs {
		pbSc, err := StateChangeToPb(sc)
		if err != nil {
			return nil, errors.Wrap(err, "StateChangeToPb")
		}
		pbScs[sc.GetId()] = pbSc
	}

	return &pbcheckout.Purchase{
		UserId:       o.UserId,
		Items:        o.Items,
		Amount:       amount,
		Timestamp:    ts,
		StateChanges: pbScs,
	}, nil
}