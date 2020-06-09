package ent

import "database/sql/driver"

func (u *Calendar) GetId() uint64 {
	return u.Id
}

func (u *Calendar) SetId(id uint64) {
	u.Id = id
}

func (u *Calendar) SQLTable() string {
	return "calendars"
}

func (u *Calendar) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":       u.Id,
		"user_id":  u.UserId,
		"name":     u.Name,
		"group_id": u.GroupId,
	}
}
