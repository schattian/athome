package ent

import (
	"encoding/json"

	"github.com/athomecomar/athome/backend/users/ent/field"
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

func (i *Identification) String() string {
	s, _ := json.Marshal(i)
	return string(s)
}
