package ent

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Calendar struct {
	Id      uint64 `json:"id,omitempty"`
	UserId  uint64 `json:"user_id,omitempty"`
	GroupId uint64 `json:"group_id,omitempty"`

	Name string `json:"name,omitempty"`
}

func CalendarsByUserId(ctx context.Context, db *sqlx.DB, uid uint64) ([]*Calendar, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM calendars WHERE user_id=$1`, uid)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()
	var cs []*Calendar
	for rows.Next() {
		c := &Calendar{}
		err := rows.StructScan(c)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func FindCalendar(ctx context.Context, db *sqlx.DB, id uint64) (*Calendar, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM calendars WHERE id=$1`, id)
	prod := &Calendar{}
	err := row.StructScan(prod)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return prod, nil
}

func (c *Calendar) Availabilities(ctx context.Context, db *sqlx.DB) ([]*Availability, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM availabilities WHERE calendar_id=$1`, c.Id)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()
	var avs []*Availability
	for rows.Next() {
		av := &Availability{}
		err := rows.StructScan(av)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		avs = append(avs, av)
	}
	return avs, nil
}

func AvailabilitiesByUserGroup(ctx context.Context, db *sqlx.DB, uid, gid uint64) ([]*Availability, error) {
	rows, err := db.QueryxContext(ctx,
		`SELECT * FROM availabilities WHERE calendar_id IN (SELECT id FROM calendars WHERE user_id=$1 AND group_id=$2)`,
		uid, gid)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()
	var avs []*Availability
	for rows.Next() {
		av := &Availability{}
		err := rows.StructScan(av)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		avs = append(avs, av)
	}
	return avs, nil
}
