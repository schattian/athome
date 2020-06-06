package data

import (
	"github.com/athomecomar/athome/backend/semantic/data/value"
	"github.com/athomecomar/storeql"
)

type Attribute interface {
	storeql.Storable

	GetSchemaId() uint64
	SetSchemaId(uint64)

	GetValue() value.Value
	SetValue(interface{}) error
}
