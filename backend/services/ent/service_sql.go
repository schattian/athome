package ent

import "database/sql/driver"

func (u *Service) GetId() uint64 {
	return u.Id
}

func (u *Service) SetId(id uint64) {
	u.Id = id
}

func (u *Service) SQLTable() string {
	return "services"
}

func (u *Service) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":          u.Id,
		"user_id":     u.UserId,
		"calendar_id": u.CalendarId,
		"address_id":  u.AddressId,

		"title":               u.Title,
		"duration_in_minutes": u.DurationInMinutes,
		"price_min":           u.PriceMin,
		"price_max":           u.PriceMax,
	}
}
