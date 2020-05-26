package ent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/jmoiron/sqlx"
)

type Onboarding struct {
	Id      uint64        `json:"id,omitempty"`
	Email   field.Email   `json:"email,omitempty"`
	Role    field.Role    `json:"role,omitempty"`
	Stage   field.Stage   `json:"stage"`
	Name    field.Name    `json:"name,omitempty"`
	Surname field.Surname `json:"surname,omitempty"`
}

func (o *Onboarding) Next() *Onboarding {
	o.Stage = o.Stage.Next(o.Role)
	return o
}

func (o *Onboarding) String() string {
	s, _ := json.Marshal(o)
	return string(s)
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
