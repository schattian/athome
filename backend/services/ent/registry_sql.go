package ent

import "database/sql/driver"

func (u *Registry) GetId() uint64 {
	return u.Id
}

func (u *Registry) SetId(id uint64) {
	u.Id = id
}

func (u *Registry) SQLTable() string {
	return "registries"
}

func (u *Registry) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":      u.Id,
		"user_id": u.UserId,
		"stage":   u.Stage,

		"address_id": u.AddressId,

		"price_min":           u.PriceMin,
		"price_max":           u.PriceMax,
		"duration_in_minutes": u.DurationInMinutes,
		"title":               u.Title,

		"calendar_id": u.CalendarId,
	}
}
