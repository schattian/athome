package ent

import (
	"encoding/json"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
)

type Identification struct {
	Id     uint64 `json:"id,omitempty"`
	UserId uint64 `json:"user_id,omitempty"`

	DNI      field.DNI `json:"dni,omitempty"`
	Verified bool      `json:"verified,omitempty"`

	Name    field.Name    `json:"name,omitempty"`
	Surname field.Surname `json:"surname,omitempty"`

	License uint64 `json:"license,omitempty"`
	Tome    uint64 `json:"tome,omitempty"`
	Folio   uint64 `json:"folio,omitempty"`

	CUE uint64 `json:"cue,omitempty"`
}

func (i *Identification) ToPb() *pbusers.Identification {
	return &pbusers.Identification{
		Verified: i.Verified,
		UserId:   i.UserId,
		Dni:      uint64(i.DNI),
		Name:     string(i.Name),
		Surname:  string(i.Surname),
		License:  i.License,
		Tome:     i.Tome,
		Folio:    i.Folio,
		Cue:      i.CUE,
	}
}

func (i *Identification) String() string {
	s, _ := json.Marshal(i)
	return string(s)
}
