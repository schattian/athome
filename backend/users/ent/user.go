package ent

import (
	"context"
	"encoding/json"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/xpbsemantic"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id uint64 `json:"id,omitempty"`

	Email        field.Email `json:"email,omitempty"`
	PasswordHash string      `json:"password_hash,omitempty"`

	CategoryId uint64        `json:"category_id,omitempty"`
	Name       field.Name    `json:"name,omitempty"`
	Surname    field.Surname `json:"surname,omitempty"`

	Role field.Role `json:"role,omitempty"`
}

func FindUser(ctx context.Context, db *sqlx.DB, uid uint64) (*User, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM users WHERE id=$1`, uid)
	u := &User{}
	err := row.StructScan(u)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return u, nil
}

func (u *User) ToPb() *pbusers.User {
	return &pbusers.User{
		Email:      string(u.Email),
		Role:       string(u.Role),
		Name:       string(u.Name),
		Surname:    string(u.Surname),
		CategoryId: u.CategoryId,
	}
}

func (u *User) Category(ctx context.Context, sem xpbsemantic.CategoriesClient) (*pbsemantic.Category, error) {
	cat, err := sem.RetrieveCategory(ctx, &pbsemantic.RetrieveCategoryRequest{CategoryId: u.CategoryId})
	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (u *User) Identification(ctx context.Context, db *sqlx.DB) (*Identification, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM identifications WHERE user_id=?`, u.Id)
	err := row.Err()
	if err != nil {
		return nil, errors.Wrap(err, "row.Err")
	}
	i := &Identification{}
	err = row.StructScan(i)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return i, nil
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
