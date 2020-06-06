package ent

import "database/sql/driver"

func (u *Draft) GetId() uint64 {
	return u.Id
}

func (u *Draft) SetId(id uint64) {
	u.Id = id
}

func (u *Draft) SQLTable() string {
	return "drafts"
}

func (u *Draft) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":      u.Id,
		"user_id": u.UserId,
		"stage":   u.Stage,
	}
}
