package ent

import (
	"context"

	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/currency"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Service struct {
	Id         uint64 `json:"id,omitempty"`
	UserId     uint64 `json:"user_id,omitempty"`
	AddressId  uint64 `json:"address_id,omitempty"`
	CalendarId uint64 `json:"calendar_id,omitempty"`

	Title             string       `json:"title,omitempty"`
	DurationInMinutes uint64       `json:"duration_in_minutes,omitempty"`
	PriceMin          currency.ARS `json:"price_min,omitempty"`
	PriceMax          currency.ARS `json:"price_max,omitempty"`
}

func FindService(ctx context.Context, db *sqlx.DB, id uint64) (*Service, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM services WHERE id=$1`, id)
	svc := &Service{}
	err := row.StructScan(svc)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return svc, nil
}

func (s *Service) User(ctx context.Context, user pbusers.ViewerClient) (*pbservices.User, error) {
	resp, err := user.RetrieveUser(ctx, &pbusers.RetrieveUserRequest{UserId: s.UserId})
	if err != nil {
		return nil, errors.Wrap(err, "user.RetrieveUser")
	}
	return &pbservices.User{
		Name:    resp.GetUser().GetName(),
		Surname: resp.GetUser().GetSurname(),
	}, nil
}
func (s *Service) Calendar(ctx context.Context, db *sqlx.DB) (*Calendar, error) {
	c, err := FindCalendar(ctx, db, s.CalendarId)
	if err != nil {
		return nil, errors.Wrap(err, "FindCalendar")
	}
	return c, nil
}

func (s *Service) Address(ctx context.Context, addr pbaddress.AddressesClient) (*pbservices.Address, error) {
	resp, err := addr.RetrieveAddress(ctx, &pbaddress.RetrieveAddressRequest{AddressId: s.AddressId})
	if err != nil {
		return nil, errors.Wrap(err, "addr.RetrieveAddress")
	}
	return &pbservices.Address{
		Zipcode:    resp.GetZipcode(),
		Street:     resp.GetStreet(),
		Number:     resp.GetNumber(),
		Floor:      resp.GetFloor(),
		Department: resp.GetDepartment(),
	}, nil
}

func (s *Service) PbPrice() *pbservices.Price {
	return &pbservices.Price{
		Min: s.PriceMin.Float64(),
		Max: s.PriceMax.Float64(),
	}
}

func (s *Service) ToPbSearchResult(ctx context.Context, users pbusers.ViewerClient) (*pbservices.ServiceSearchResult, error) {
	user, err := s.User(ctx, users)
	if err != nil {
		return nil, errors.Wrap(err, "User")
	}
	return &pbservices.ServiceSearchResult{
		User: user,
		Service: &pbservices.ServiceSearchResult_Service{
			AddressId:  s.AddressId,
			CalendarId: s.CalendarId,
			Title:      s.Title,
			Price:      s.PbPrice(),
		},
	}, nil
}

func (s *Service) ToPb() *pbservices.Service {
	return &pbservices.Service{
		Title:      s.Title,
		UserId:     s.UserId,
		AddressId:  s.AddressId,
		CalendarId: s.CalendarId,

		DurationInMinutes: s.DurationInMinutes,
		Price:             s.PbPrice(),
	}
}
