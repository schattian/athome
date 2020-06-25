package order

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/athomecomar/athome/backend/checkout/ent"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Purchase struct {
	Id            uint64 `json:"id,omitempty"`
	UserId        uint64 `json:"user_id,omitempty"`
	DestAddressId uint64 `json:"dest_address_id,omitempty"`

	DistanceInKilometers float64           `json:"distance_in_kilometers,omitempty"`
	SrcAddressId         uint64            `json:"src_address_id,omitempty"`
	CreatedAt            ent.Time          `json:"created_at,omitempty"`
	MerchantId           uint64            `json:"merchant_id,omitempty"`
	UpdatedAt            ent.Time          `json:"updated_at,omitempty"`
	Items                map[uint64]uint64 `json:"items,omitempty"`
	ShippingId           uint64            `json:"shipping_id,omitempty"`
}

func (o *Purchase) GetCreatedAt() time.Time { return o.CreatedAt.Time }
func (o *Purchase) GetUpdatedAt() time.Time { return o.UpdatedAt.Time }
func (o *Purchase) SetCreatedAt(t time.Time) {
	o.CreatedAt = ent.Time{NullTime: sql.NullTime{Time: t}}
	o.SetUpdatedAt(t)
}
func (o *Purchase) SetUpdatedAt(t time.Time) { o.UpdatedAt = ent.Time{NullTime: sql.NullTime{Time: t}} }

func (o *Purchase) State(ctx context.Context, db *sqlx.DB) (*PurchaseStateChange, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM purchase_state_changes WHERE order_id IN (
                SELECT id FROM purchases WHERE id = $1
            )
             ORDER BY stage ASC, created_at DESC`,
		o.Id,
	)
	sc := &PurchaseStateChange{}
	err := row.StructScan(sc)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return sc, nil
}

func FindPurchase(ctx context.Context, db *sqlx.DB, oId uint64, userId uint64) (*Purchase, error) {
	order := &Purchase{}
	row := storeql.Where(ctx, db, order, `id=$1 AND user_id=$2`, oId, userId)
	err := row.StructScan(order)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return order, nil
}

func FindPurchasesByUser(ctx context.Context, db *sqlx.DB, uid uint64) ([]*Purchase, error) {
	rows, err := storeql.WhereMany(ctx, db, &Purchase{}, `user_id=$1`, uid)
	if err != nil {
		return nil, errors.Wrap(err, "WhereMany")
	}
	var orders []*Purchase
	for rows.Next() {
		order := &Purchase{}
		err = rows.StructScan(order)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (p *Purchase) Shipping(ctx context.Context, db *sqlx.DB) (*Shipping, error) {
	return FindShipping(ctx, db, p.ShippingId)
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

func (o *Purchase) Merchant(ctx context.Context, c pbusers.ViewerClient) (*pbusers.User, error) {
	return c.RetrieveUser(ctx, &pbusers.RetrieveUserRequest{UserId: o.MerchantId})
}

func (o *Purchase) Products(ctx context.Context, c pbproducts.ViewerClient) (map[uint64]*pbproducts.Product, error) {
	prods, err := c.RetrieveProducts(ctx, &pbproducts.RetrieveProductsRequest{Ids: o.ProductIds()})
	if err != nil {
		return nil, errors.Wrap(err, "products.RetrieveProducts")
	}
	return prods.Products, nil
}

func (o *Purchase) MaxVolWeight(ctx context.Context, c pbproducts.ViewerClient) (float64, error) {
	resp, err := c.RetrieveProductsMaxVolWeight(ctx, &pbproducts.RetrieveProductsRequest{Ids: o.ProductIds()})
	if err != nil {
		return 0, errors.Wrap(err, "products.RetrieveProductsMaxColWeight")
	}
	var totalMaxVolWeight float64
	for prodId, maxVolWeight := range resp.GetMaxVolWeights() {
		totalMaxVolWeight += float64(o.Items[prodId]) * maxVolWeight
	}
	return totalMaxVolWeight, nil
}

func (o *Purchase) ProductIds() (ids []uint64) {
	for pid := range o.Items {
		ids = append(ids, pid)
	}
	return
}

func (o *Purchase) Amount(ctx context.Context, db *sqlx.DB, c pbproducts.ViewerClient) (float64, error) {
	prods, err := o.Products(ctx, c)
	if err != nil {
		return 0, errors.Wrap(err, "Products")
	}
	amount, err := o.AmountFromProducts(ctx, prods)
	if err != nil {
		return 0, errors.Wrap(err, "AmountFromProducts")
	}

	if o.ShippingId == 0 {
		return amount, nil
	}

	ship, err := o.Shipping(ctx, db)
	if err != nil {
		return 0, errors.Wrap(err, "Shipping")
	}
	return amount + ship.OrderPrice.Float64(), nil
}

func (o *Purchase) AmountFromProducts(ctx context.Context, prods map[uint64]*pbproducts.Product) (float64, error) {
	var amount float64
	for id, prod := range prods {
		desiredStock, gotStock := o.Items[id], prod.GetStock()
		if desiredStock > gotStock {
			return 0, fmt.Errorf("product with id %v has %d of stock while asked for %d", id, gotStock, desiredStock)
		}
		plus := float64(desiredStock) * prod.GetPrice()
		if plus <= 0 {
			return 0, fmt.Errorf("invalid amount calculated. Given quantity %v with prod price %v", desiredStock, prod.GetPrice())
		}
		amount += plus
	}
	return amount, nil
}

func (o *Purchase) ToPbWrapped(ctx context.Context, db *sqlx.DB, prods pbproducts.ViewerClient) (*pbcheckout.Purchase, error) {
	scs, err := StateChanges(ctx, db, o)
	if err != nil {
		return nil, errors.Wrap(err, "StateChanges")
	}
	amount, err := o.Amount(ctx, db, prods)
	if err != nil {
		return nil, errors.Wrap(err, "Amount")
	}
	pb, err := o.ToPb(scs, amount)
	if err != nil {
		return nil, errors.Wrap(err, "ToPb")
	}
	return pb, nil
}

func (o *Purchase) ValidateStateChange(ctx context.Context, db *sqlx.DB, newState *sm.State) (err error) {
	switch newState.Name {
	case sm.PurchaseAddress:
		err = o.validateStateChangeAddress(ctx, db)
	}
	return
}

func (o *Purchase) validateStateChangeAddress(ctx context.Context, db *sqlx.DB) error {
	if o.DestAddressId == 0 {
		return errors.New("nil dest address id")
	}
	return nil
}

func (o *Purchase) AssignSrcAddress(ctx context.Context, users pbusers.ViewerClient, addr pbaddress.AddressesClient) error {
	if o.MerchantId == 0 {
		return errors.New("no merchant assigned")
	}
	u, err := users.RetrieveUser(ctx, &pbusers.RetrieveUserRequest{UserId: o.MerchantId})
	if err != nil {
		return errors.Wrap(err, "users.RetrieveUser")
	}
	resp, err := addr.MeasureDistance(ctx, &pbaddress.MeasureDistanceRequest{BAddressId: o.SrcAddressId, AAddressId: u.GetAddressId()})
	if err != nil {
		return errors.Wrap(err, "addr.MeasureDistance")
	}
	o.SrcAddressId = u.GetAddressId()
	o.DistanceInKilometers = resp.ManhattanInKilometers
	return nil
}

func (o *Purchase) AssignMerchant(ctx context.Context, prods map[uint64]*pbproducts.Product) error {
	for id, prod := range prods {
		if o.MerchantId == 0 {
			o.MerchantId = prod.GetUserId()
		}
		if prod.GetUserId() != o.MerchantId {
			return fmt.Errorf("product with id %v mixes merchant id on order", id)
		}
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
		UserId:        o.UserId,
		Items:         o.Items,
		DestAddressId: o.DestAddressId,
		SrcAddressId:  o.SrcAddressId,
		Amount:        amount,
		Timestamp:     ts,
		MerchantId:    o.MerchantId,
		ShippingId:    o.ShippingId,
	}, nil
}
