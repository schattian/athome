package ent

import (
	"context"

	"github.com/athomecomar/athome/backend/services/ent/stage"
	"github.com/athomecomar/currency"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Registry struct {
	Id     uint64      `json:"id,omitempty"`
	UserId uint64      `json:"user_id,omitempty"`
	Stage  stage.Stage `json:"stage,omitempty"`

	// First
	AddressId uint64 `json:"address_id,omitempty"`

	// Second
	Name              string       `json:"name,omitempty"`
	DurationInMinutes uint64       `json:"duration_in_minutes,omitempty"`
	PriceMin          currency.ARS `json:"price_min,omitempty"`
	PriceMax          currency.ARS `json:"price_max,omitempty"`

	// Third
	CalendarId uint64 `json:"calendar_id,omitempty"`
}

func (r *Registry) ToService() *Service {
	return &Service{
		UserId:            r.UserId,
		AddressId:         r.AddressId,
		Name:              r.Name,
		DurationInMinutes: r.DurationInMinutes,
		PriceMax:          r.PriceMax,
		PriceMin:          r.PriceMin,
		CalendarId:        r.CalendarId,
	}
}

func NewRegistry(userId uint64) *Registry {
	return &Registry{UserId: userId, Stage: stage.First}
}

func FindRegistry(ctx context.Context, db *sqlx.DB, id uint64) (*Registry, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM registries WHERE id=$1`, id)
	prod := &Registry{}
	err := row.StructScan(prod)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return prod, nil
}

func FindRegistryByUserId(ctx context.Context, db *sqlx.DB, uid uint64) (*Registry, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM registries WHERE user_id=$1`, uid)
	prod := &Registry{}
	err := row.StructScan(prod)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return prod, nil
}
