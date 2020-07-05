package order

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
	Id               uint64
	UserId           uint64
	EventId          uint64
	ShippingMethodId uint64

	SrcAddressId      uint64
	DestAddressId     uint64
	ManhattanDistance float64

	OrderPrice             currency.ARS
	OrderDurationInMinutes uint64

	RealPrice             currency.ARS
	RealDurationInMinutes uint64
}

func NewShipping(ctx context.Context, db *sqlx.DB,
	p *Purchase,
	eventId uint64,
	providerId uint64,
	shippingMethodId uint64,
	orderPrice currency.ARS,
	orderDuration uint64,
) *Shipping {
	return &Shipping{
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

func (s *Shipping) Purchase(ctx context.Context, db *sqlx.DB) (p *Purchase, err error) {
	row := storeql.Where(ctx, db, p, "shipping_id=$1", s.Id)
	err = row.StructScan(p)
	if err != nil {
		err = errors.Wrap(err, "storeql.Where")
		return
	}
	return
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
