package ent

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/semantic/semerr"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
)

type Onboarding struct {
	Id       uint64        `json:"id,omitempty"`
	Email    field.Email   `json:"email,omitempty"`
	Role     field.Role    `json:"role,omitempty"`
	Stage    field.Stage   `json:"stage"`
	Name     field.Name    `json:"name,omitempty"`
	Surname  field.Surname `json:"surname,omitempty"`
	Category string        `json:"category,omitempty"`
}

func (o *Onboarding) Identification(ctx context.Context, db *sqlx.DB) (*OnboardingIdentification, error) {
	if o.Category == "" { // TODO: Add "if onboarding should passed the step X" <field.Identificaton>
		return nil, nil
	}
	row := db.QueryRowxContext(ctx, `SELECT * FROM onboarding_identifications WHERE onboarding_id=?`, o.Id)
	err := row.Err()
	if err != nil {
		return nil, errors.Wrap(err, "row.Err")
	}
	oi := &OnboardingIdentification{}
	err = row.StructScan(oi)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return oi, nil
}

func (o *Onboarding) Next() *Onboarding {
	o.Stage = o.Stage.Next(o.Role)
	return o
}

func (o *Onboarding) String() string {
	s, _ := json.Marshal(o)
	return string(s)
}

func (o *Onboarding) SetCategory(categoryName string) (err error) {
	switch o.Role {
	case field.Merchant:
		err = o.setCategoryMerchant(categoryName)
	case field.ServiceProvider:
		err = o.setCategoryServiceProvider(categoryName)
	default:
		err = fmt.Errorf("invalid role (not classifiable): %v", o.Role)
	}
	return
}

func (o *Onboarding) setCategoryMerchant(_ string) error {
	return errors.New("not implemented")
}

func (o *Onboarding) setCategoryServiceProvider(categoryName string) error {
	cat := semprov.Loc(categoryName)
	if cat == nil {
		return semerr.ErrProviderCategoryNotFound
	}
	if cat.Childs != nil {
		return fmt.Errorf("invalid category: %s. It got %d childs", cat.Name, len(cat.Childs))
	}
	return nil
}

func (o *Onboarding) Close(ctx context.Context, db *sqlx.DB, pwd string) (*User, *Identification, error) {
	user := o.ToUser()

	err := user.AssignPassword(pwd)
	if err != nil {
		return nil, nil, errors.Wrap(err, "AssignPassword")
	}

	err = storeql.InsertIntoDB(ctx, db, user)
	if err != nil {
		return nil, nil, errors.Wrap(err, "user InsertIntoDB")
	}

	oid, err := o.Identification(ctx, db)
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		return nil, nil, errors.Wrap(err, "Identification")
	}

	var id *Identification
	if oid != nil {
		id, err = oid.Close(ctx, db, user.GetId())
		if err != nil {
			return nil, nil, errors.Wrap(err, "oid.Close")
		}
	}

	err = storeql.DeleteFromDB(ctx, db, o)
	if err != nil {
		return nil, nil, errors.Wrap(err, "onboarding DeleteFromDB")
	}

	return user, id, nil
}

func (o *Onboarding) ToUser() *User {
	return &User{Email: o.Email, Role: o.Role, Name: o.Name, Surname: o.Surname}
}

func (o *Onboarding) ValidateByStage(ctx context.Context, db *sqlx.DB) (code codes.Code, err error) {
	switch o.Stage {
	case field.Shared:
		code, err = o.ValidateShared(ctx, db)
	}
	return
}

func (o *Onboarding) MustStage(s field.Stage) (code codes.Code, err error) {
	if o.Stage != s {
		return codes.OutOfRange, fmt.Errorf("invalid stage %v", o.Stage)
	}
	return
}

func (o *Onboarding) ValidateShared(ctx context.Context, db *sqlx.DB) (code codes.Code, err error) {
	rows, err := db.QueryxContext(ctx, `SELECT COUNT(id) FROM users WHERE email=$1 AND role=$2 LIMIT 1`, o.Email, o.Role)
	if err != nil {
		return codes.Internal, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return codes.Internal, errors.Wrap(err, "row.Scan")
		}
	}
	if count > 0 {
		return codes.AlreadyExists, errors.New("Invalid count")
	}
	return codes.OK, nil
}
