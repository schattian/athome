package data

import "github.com/athomecomar/athome/backend/semantic/data/value"

type Attribute interface {
	GetSchemaId() uint64
	SetSchemaId(uint64)

	GetValue() value.Value
	SetValue(value.Value) error
}
