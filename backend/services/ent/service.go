package ent

import (
	"context"

	"github.com/athomecomar/athome/backend/services/pb/pbaddress"
	"github.com/athomecomar/athome/backend/services/pb/pbservices"
	"github.com/athomecomar/athome/backend/services/pb/pbusers"
	"github.com/athomecomar/currency"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Service struct {
	Id         uint64 `json:"id,omitempty"`
	UserId     uint64 `json:"user_id,omitempty"`
	AddressId  uint64 `json:"address_id,omitempty"`
	CalendarId uint64 `json:"calendar_id,omitempty"`

	Name              string       `json:"name,omitempty"`
	DurationInMinutes uint64       `json:"duration_in_minutes,omitempty"`
	PriceMin          currency.ARS `json:"price_min,omitempty"`
	PriceMax          currency.ARS `json:"price_max,omitempty"`
}

func FindService(ctx context.Context, db *sqlx.DB, id uint64) (*Service, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM registries WHERE id=$1`, id)
	svc := &Service{}
	err := row.StructScan(svc)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return svc, nil
}

func (s *Service) User(ctx context.Context, user pbusers.ViewerClient) (*pbservices.UserData, error) {
	resp, err := user.ViewUser(ctx, &pbusers.ViewUserRequest{UserId: s.UserId})
	if err != nil {
		return nil, errors.Wrap(err, "user.ViewUser")
	}
	return &pbservices.UserData{
		Name:    resp.GetName(),
		Surname: resp.GetSurname(),
	}, nil
}
func (s *Service) Calendar(ctx context.Context, db *sqlx.DB) (*Calendar, error) {
	c, err := FindCalendar(ctx, db, s.CalendarId)
	if err != nil {
		return nil, errors.Wrap(err, "FindCalendar")
	}
	return c, nil
}

func (s *Service) Address(ctx context.Context, addr pbaddress.AddressClient) (*pbservices.AddressData, error) {
	resp, err := addr.RetrieveAddress(ctx, &pbaddress.RetrieveAddressRequest{AddressId: s.AddressId})
	if err != nil {
		return nil, errors.Wrap(err, "addr.RetrieveAddress")
	}
	return &pbservices.AddressData{
		Zipcode:    resp.GetZipcode(),
		Street:     resp.GetStreet(),
		Number:     resp.GetNumber(),
		Floor:      resp.GetFloor(),
		Department: resp.GetDepartment(),
	}, nil
}
