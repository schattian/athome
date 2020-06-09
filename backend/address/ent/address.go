package ent

import (
	"context"

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
	Department uint64
	Latitude   uint64
	Longitude  uint64
	Alias      string
}

func FindAddress(ctx context.Context, db *sqlx.DB, id uint64) (*Address, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM calendars WHERE id=$1`, id)
	prod := &Address{}
	err := row.StructScan(prod)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return prod, nil
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
