package purchase

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/athomecomar/athome/backend/checkout/ent"
	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/ent/payment"
	"github.com/athomecomar/athome/backend/checkout/ent/shipping"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/currency"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Items map[uint64]uint64

func (i Items) Scan(val interface{}) (err error) {
	switch v := val.(type) {
	case []byte:
		err = json.Unmarshal(v, &i)
	case string:
		err = json.Unmarshal([]byte(v), &i)
	default:
		err = errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
	return
}
func (i Items) Value() (driver.Value, error) {
	return json.Marshal(&i)
}

type Purchase struct {
	Id            uint64 `json:"id,omitempty"`
	UserId        uint64 `json:"user_id,omitempty"`
	DestAddressId uint64 `json:"dest_address_id,omitempty"`
	SrcAddressId  uint64 `json:"src_address_id,omitempty"`

	DistanceInKilometers float64  `json:"distance_in_kilometers,omitempty"`
	CreatedAt            ent.Time `json:"created_at,omitempty"`
	MerchantId           uint64   `json:"merchant_id,omitempty"`
	UpdatedAt            ent.Time `json:"updated_at,omitempty"`
	Items                Items    `json:"items,omitempty"`
	ShippingId           uint64   `json:"shipping_id,omitempty"`
}

func (o *Purchase) GetCreatedAt() time.Time { return o.CreatedAt.Time }
func (o *Purchase) GetUpdatedAt() time.Time { return o.UpdatedAt.Time }
func (o *Purchase) SetCreatedAt(t time.Time) {
	o.CreatedAt = ent.Time{NullTime: sql.NullTime{Time: t}}
	o.SetUpdatedAt(t)
}
func (o *Purchase) SetUpdatedAt(t time.Time) { o.UpdatedAt = ent.Time{NullTime: sql.NullTime{Time: t}} }

func (p *Purchase) NewShipping(ctx context.Context, db *sqlx.DB,
	eventId uint64,
	providerId uint64,
	shippingMethodId uint64,
	orderPrice currency.ARS,
	orderDuration uint64,
) *shipping.Shipping {
	return &shipping.Shipping{
		EventId:                eventId,
		OrderPrice:             orderPrice,
		OrderDurationInMinutes: orderDuration,
		SrcAddressId:           p.SrcAddressId,
		DestAddressId:          p.DestAddressId,
		ManhattanDistance:      p.DistanceInKilometers,
		UserId:                 providerId,
		ShippingMethodId:       shippingMethodId,
	}
}

func FindPurchase(ctx context.Context, db *sqlx.DB, oId uint64) (*Purchase, error) {
	order := &Purchase{}
	row := storeql.Where(ctx, db, order, `id=$1`, oId)
	err := row.StructScan(order)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return order, nil
}

func FindPurchaseByShipping(ctx context.Context, db *sqlx.DB, oId uint64) (*Purchase, error) {
	order := &Purchase{}
	row := storeql.Where(ctx, db, order, `shipping_id=$1`, oId)
	err := row.StructScan(order)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return order, nil
}

func FindPurchaseUserScoped(ctx context.Context, db *sqlx.DB, oId uint64, userId uint64) (*Purchase, error) {
	order := &Purchase{}
	row := storeql.Where(ctx, db, order, `id=$1 AND user_id=$2`, oId, userId)
	err := row.StructScan(order)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return order, nil
}

func (p *Purchase) AmountPaid(ctx context.Context, db *sqlx.DB) (currency.ARS, error) {
	pys, err := p.Payments(ctx, db)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	if err != nil {
		return 0, errors.Wrap(err, "Payments")
	}
	var total currency.ARS
	for _, py := range pys {
		total += py.Amount
	}
	return total, nil
}

func (p *Purchase) CanView(ctx context.Context, db *sqlx.DB, userId uint64) (bool, error) {
	if p.MerchantId == userId {
		return true, nil
	}
	if p.UserId == userId {
		return true, nil
	}
	ship, err := p.Shipping(ctx, db)
	if err != nil {
		return false, errors.Wrap(err, "Shipping")
	}
	if ship.UserId == userId {
		return true, nil
	}
	return false, nil
}

func (p *Purchase) Payments(ctx context.Context, db *sqlx.DB) ([]*payment.Payment, error) {
	rows, err := storeql.WhereMany(ctx, db, &payment.Payment{}, `entity_id=$1 AND entity_table=$2`, p.Id, p.SQLTable())
	if err != nil {
		return nil, errors.Wrap(err, "WhereMany")
	}
	var orders []*payment.Payment
	for rows.Next() {
		order := &payment.Payment{}
		err = rows.StructScan(order)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		orders = append(orders, order)
	}
	return orders, nil
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

func (o *Purchase) OrderClass() order.Class {
	return order.Purchases
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
	var intIds []int
	for pid := range o.Items {
		intIds = append(intIds, int(pid))
	}
	sort.Ints(intIds)
	for _, intId := range intIds {
		ids = append(ids, uint64(intId))
	}
	return
}

func (o *Purchase) Shipping(ctx context.Context, db *sqlx.DB) (*shipping.Shipping, error) {
	return shipping.FindShipping(ctx, db, o.ShippingId)
}

func (o *Purchase) TotalPaid(ctx context.Context, db *sqlx.DB) (amount currency.ARS, err error) {
	pys, err := o.FinishedPayments(ctx, db)
	if err != nil {
		err = errors.Wrap(err, "FinishedPayments")
		return
	}
	for _, py := range pys {
		amount += py.Amount
	}
	return
}

func (o *Purchase) FinishedPayments(ctx context.Context, db *sqlx.DB) (pys []*payment.Payment, err error) {
	var allPys []*payment.Payment
	allPys, err = o.Payments(ctx, db)
	if err != nil {
		err = errors.Wrap(err, "Payments")
		return
	}
	var isFinished bool
	for _, py := range allPys {
		isFinished, err = py.IsFinished(ctx, db)
		if err != nil {
			err = errors.Wrap(err, "IsFinished")
			return
		}
		if isFinished {
			pys = append(pys, py)
		}
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
	amount, err := o.Amount(ctx, db, prods)
	if err != nil {
		return nil, errors.Wrap(err, "Amount")
	}
	pb, err := o.ToPb(amount)
	if err != nil {
		return nil, errors.Wrap(err, "ToPb")
	}
	return pb, nil
}

func (o *Purchase) AssignSrcAddress(ctx context.Context, users pbusers.ViewerClient) error {
	if o.MerchantId == 0 {
		return errors.New("no merchant assigned")
	}
	u, err := users.RetrieveUser(ctx, &pbusers.RetrieveUserRequest{UserId: o.MerchantId})
	if err != nil {
		return errors.Wrap(err, "users.RetrieveUser")
	}
	o.SrcAddressId = u.GetAddressId()
	return nil
}

func (o *Purchase) AssignDistance(ctx context.Context, addr pbaddress.AddressesClient) error {
	resp, err := addr.MeasureDistance(ctx, &pbaddress.MeasureDistanceRequest{
		AAddressId: o.SrcAddressId,
		BAddressId: o.DestAddressId,
	})
	if err != nil {
		return errors.Wrap(err, "addr.MeasureDistance")
	}
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

func NewPurchase(ctx context.Context, items Items, uid uint64) *Purchase {
	p := &Purchase{UserId: uid, Items: items}
	p.SetCreatedAt(time.Now())
	return p
}

func (o *Purchase) ToPb(amount float64) (*pbcheckout.Purchase, error) {
	ts, err := ent.GetTimestamp(o)
	if err != nil {
		return nil, errors.Wrap(err, "GetTimestamp")
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
