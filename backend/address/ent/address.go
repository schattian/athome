package ent

import (
	"context"

	"github.com/athomecomar/athome/backend/address/ent/distance"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Address struct {
	Id     uint64 `json:"id,omitempty"`
	UserId uint64 `json:"user_id,omitempty"`

	Country    string
	Province   string
	Zipcode    string
	Street     string
	Number     uint64
	Floor      uint64
	Department string
	Latitude   float64
	Longitude  float64
	Alias      string
}

func (a *Address) coord() distance.Coord {
	return distance.Coord{
		Lat: a.Latitude,
		Lon: a.Longitude,
	}
}

func (a *Address) DistanceHaversine(b *Address) float64 {
	return distance.Haversine(a.coord(), b.coord())
}

func (a *Address) DistanceManhattan(b *Address) float64 {
	return distance.Manhattan(a.coord(), b.coord())
}

func FindAddress(ctx context.Context, db *sqlx.DB, id uint64) (*Address, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM addresses WHERE id=$1`, id)
	prod := &Address{}
	err := row.StructScan(prod)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return prod, nil
}

func FindAddresses(ctx context.Context, db *sqlx.DB, ids ...uint64) ([]*Address, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM addresses WHERE id IN($1)`, ids)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()
	var addrs []*Address
	for rows.Next() {
		addr := &Address{}
		err := rows.StructScan(addr)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		addrs = append(addrs, addr)
	}
	return addrs, nil
}

func AddressFromPb(d *pbaddress.Address) *Address {
	return &Address{
		Country:   d.GetCountry(),
		Province:  d.GetProvince(),
		Zipcode:   d.GetZipcode(),
		Street:    d.GetStreet(),
		Number:    d.GetNumber(),
		Floor:     d.GetFloor(),
		Latitude:  d.GetLatitude(),
		Longitude: d.GetLongitude(),

		Alias: d.GetAlias(),
	}
}

func (addr *Address) ToPb() *pbaddress.Address {
	return &pbaddress.Address{
		Country:    addr.Country,
		Province:   addr.Province,
		Zipcode:    addr.Zipcode,
		Street:     addr.Street,
		Number:     addr.Number,
		Department: addr.Department,
		Floor:      addr.Floor,
		Latitude:   addr.Latitude,
		Longitude:  addr.Longitude,
		Alias:      addr.Alias,
	}
}

func AddressesByUser(ctx context.Context, db *sqlx.DB, uid uint64) ([]*Address, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM addresses WHERE user_id=$1`, uid)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()
	var avs []*Address
	for rows.Next() {
		av := &Address{}
		err := rows.StructScan(av)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		avs = append(avs, av)
	}
	return avs, nil
}
