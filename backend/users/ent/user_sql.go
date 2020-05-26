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
		"email":         string(u.Email),
		"password_hash": u.PasswordHash,
		"role":          string(u.Role),
		"name":          string(u.Name),
		"surname":       string(u.Surname),
	}
}
