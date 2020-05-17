package ent

import (
	"github.com/athomecomar/athome/users/ent/field"
	"github.com/athomecomar/xerrors"
)

type User struct {
	Id int64 `json:"id,omitempty"`

	Email        field.Email `json:"email,omitempty"`
	PasswordHash string      `json:"password_hash,omitempty"`

	Name    field.Name    `json:"name,omitempty"`
	Surname field.Surname `json:"surname,omitempty"`

	Role field.Role `json:"role,omitempty"`
}

func (u *User) GetId() int64 {
	return u.Id
}

func (u *User) SetId(id int64) {
	u.Id = id
}

func (u *User) SQLTable() string {
	return "users"
}

func (u *User) SQLColumns() []string {
	return []string{
		"id",
		"email",
		"password_hash",
		"role",
		"name",
		"surname",
	}
}

func (u *User) Validate() *xerrors.Errors {
	errs := xerrors.NewErrors()
	for _, field := range []Validable{u.Email, u.Role, u.Name, u.Surname, u.Role} {
		if err := field.Validate(); err != nil {
			errs.Add(field.Name(), err.Error())
		}
	}
	if errs.Count() > 0 {
		return errs
	}
	return nil
}
