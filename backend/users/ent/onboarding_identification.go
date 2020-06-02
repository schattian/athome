package ent

import (
	"context"
	"encoding/json"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type OnboardingIdentification struct {
	Id           uint64 `json:"id,omitempty"`
	OnboardingId uint64 `json:"onboarding_id,omitempty"`

	DNI field.DNI `json:"dni,omitempty"`

	Name    field.Name    `json:"name,omitempty"`
	Surname field.Surname `json:"surname,omitempty"`

	License uint64 `json:"license,omitempty"`
	Tome    uint64 `json:"tome,omitempty"`
	Folio   uint64 `json:"folio,omitempty"`

	CUE uint64 `json:"cue,omitempty"`
}

func (oi *OnboardingIdentification) Close(ctx context.Context, db *sqlx.DB, userId uint64) (*Identification, error) {
	id := oi.ToIdentification()
	id.UserId = userId
	err := storeql.InsertIntoDB(ctx, db, id)
	if err != nil {
		return nil, errors.Wrap(err, "identification InsertIntoDB")
	}
	err = storeql.DeleteFromDB(ctx, db, oi)
	if err != nil {
		return nil, errors.Wrap(err, "onboarding identification DeleteFromDB")
	}
	return id, nil
}

func (oi *OnboardingIdentification) String() string {
	s, _ := json.Marshal(oi)
	return string(s)
}

func (o *OnboardingIdentification) ToIdentification() *Identification {
	return &Identification{
		DNI: o.DNI,

		Name:    o.Name,
		Surname: o.Surname,

		License: o.License,
		Tome:    o.Tome,
		Folio:   o.Folio,
		CUE:     o.CUE,

		Verified: false,
	}
}
