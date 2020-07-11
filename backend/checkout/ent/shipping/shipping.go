package shipping

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/currency"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Shipping struct {
	Id               uint64 `json:"id,omitempty"`
	UserId           uint64 `json:"user_id,omitempty"`
	EventId          uint64 `json:"event_id,omitempty"`
	ShippingMethodId uint64 `json:"shipping_method_id,omitempty"`

	SrcAddressId      uint64  `json:"src_address_id,omitempty"`
	DestAddressId     uint64  `json:"dest_address_id,omitempty"`
	ManhattanDistance float64 `json:"manhattan_distance,omitempty"`

	OrderPrice             currency.ARS `json:"order_price,omitempty"`
	OrderDurationInMinutes uint64       `json:"order_duration_in_minutes,omitempty"`

	RealPrice             currency.ARS `json:"real_price,omitempty"`
	RealDurationInMinutes uint64       `json:"real_duration_in_minutes,omitempty"`
}

func (s *Shipping) ToPb() *pbcheckout.Shipping {
	return &pbcheckout.Shipping{
		UserId:               s.UserId,
		DurationInMinutes:    s.OrderDurationInMinutes,
		Amount:               s.OrderPrice.Float64(),
		EventId:              s.EventId,
		SrcAddressId:         s.SrcAddressId,
		DestAddressId:        s.DestAddressId,
		DistanceInKilometers: s.ManhattanDistance,
	}
}

func (s *Shipping) DiffDurationInMinutes() uint64 {
	return s.OrderDurationInMinutes - s.RealDurationInMinutes
}

func (s *Shipping) DiffPricePerKilometer() float64 {
	return s.OrderPricePerKilometer() - s.RealPricePerKilometer()
}

func (s *Shipping) DiffPrice() float64 {
	return (s.OrderPrice - s.RealPrice).Float64()
}

func FindShipping(ctx context.Context, db *sqlx.DB, id uint64) (*Shipping, error) {
	ship := &Shipping{}
	row := storeql.Where(ctx, db, ship, "id=$1", id)
	err := row.StructScan(ship)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return ship, nil
}

func (s *Shipping) RealPricePerKilometer() float64 {
	return s.pricePerKilometer(s.RealPrice)
}

func (s *Shipping) OrderPricePerKilometer() float64 {
	return s.pricePerKilometer(s.OrderPrice)
}

func (s *Shipping) pricePerKilometer(p currency.ARS) float64 {
	return p.Float64() / s.ManhattanDistance
}

func CalculateShippingPricePerKilometer(ctx context.Context, db *sqlx.DB, userId uint64, price *pbservices.Price) (currency.ARS, error) {
	ships, err := FindShippingsByUser(ctx, db, userId)
	if errors.Is(err, sql.ErrNoRows) {
		return currency.ToARS((price.Max - price.Min)) / 2, nil
	}

	if err != nil {
		return 0, errors.Wrap(err, "FindShippingsByUser")
	}
	var shippingPrice float64
	for _, ship := range ships {
		shippingPrice = (ship.RealPricePerKilometer() + shippingPrice) / 2
	}
	return currency.ToARS(shippingPrice), nil
}

func FindShippingsByUser(ctx context.Context, db *sqlx.DB, uid uint64) ([]*Shipping, error) {
	rows, err := storeql.WhereMany(ctx, db, &Shipping{}, "user_id=$1", uid)
	if err != nil {
		return nil, errors.Wrap(err, "WhereMany")
	}
	var ships []*Shipping
	for rows.Next() {
		ship := &Shipping{}
		err = rows.StructScan(ship)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		ships = append(ships, ship)
	}
	return ships, nil
}
