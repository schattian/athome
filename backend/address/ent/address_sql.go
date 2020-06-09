package ent

import "database/sql/driver"

func (u *Address) GetId() uint64 {
	return u.Id
}

func (u *Address) SetId(id uint64) {
	u.Id = id
}

func (u *Address) SQLTable() string {
	return "addresses"
}

func (u *Address) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":         u.Id,
		"user_id":    u.UserId,
		"country":    u.Country,
		"province":   u.Province,
		"zipcode":    u.Zipcode,
		"street":     u.Street,
		"number":     u.Number,
		"floor":      u.Floor,
		"department": u.Department,
		"latitude":   u.Latitude,
		"longitude":  u.Longitude,
		"alias":      u.Alias,
	}
}
