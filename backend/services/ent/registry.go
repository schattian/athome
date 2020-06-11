package ent

import (
	"context"

	"github.com/athomecomar/athome/backend/services/ent/stage"
	"github.com/athomecomar/athome/backend/services/pb/pbservices"
	"github.com/athomecomar/currency"
	"github.com/athomecomar/storeql"
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
func (r *Registry) Calendar(ctx context.Context, db *sqlx.DB) (*Calendar, error) {
	c := &Calendar{}
	row := storeql.Where(ctx, db, c, `id=$1`, r.CalendarId)
	err := row.StructScan(c)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return nil, nil
}

func NewRegistry(userId uint64) *Registry {
	return &Registry{UserId: userId, Stage: stage.First}
}

func (r *Registry) ToPb() *pbservices.Registry {
	return &pbservices.Registry{
		Stage: uint64(r.Stage),

		First: &pbservices.FirstRequest_Body{
			AddressId: r.AddressId,
		},

		Second: &pbservices.SecondRequest_Body{
			Name:              r.Name,
			DurationInMinutes: r.DurationInMinutes,
			Price: &pbservices.Price{
				Max: r.PriceMax.Float64(),
				Min: r.PriceMin.Float64(),
			},
		},

		Third: &pbservices.ThirdRequest_Body{
			CalendarId: r.CalendarId,
		},
	}
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
