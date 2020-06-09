package ent

import "github.com/athomecomar/currency"

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
