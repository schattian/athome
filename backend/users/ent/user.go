package ent

import (
	"encoding/json"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/xerrors"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id uint64 `json:"id,omitempty"`

	Email        field.Email `json:"email,omitempty"`
	PasswordHash string      `json:"password_hash,omitempty"`

	Name    field.Name    `json:"name,omitempty"`
	Surname field.Surname `json:"surname,omitempty"`

	Role field.Role `json:"role,omitempty"`
}

func (u *User) AssignPassword(pwd string) error {
	hash, err := passwordHash(pwd)
	if err != nil {
		return errors.Wrap(err, "passwordHash")
	}
	u.PasswordHash = hash
	return nil
}

func (u *User) String() string {
	s, _ := json.Marshal(u)
	return string(s)
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

func passwordHash(pwd string) (string, error) {
	ph, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "bcrypt.GenerateFromPassword")
	}
	return string(ph), nil
}
