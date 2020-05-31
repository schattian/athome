package ent

import "database/sql/driver"

func (u *User) GetId() uint64 {
	return u.Id
}

func (u *User) SetId(id uint64) {
	u.Id = id
}

func (u *User) SQLTable() string {
	return "users"
}

func (u *User) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":            u.Id,
		"password_hash": u.PasswordHash,
		"email":         u.Email,
		"category":      u.Category,
		"role":          u.Role,
		"name":          u.Name,
		"surname":       u.Surname,
	}
}
