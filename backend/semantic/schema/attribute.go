package schema

import (
	"github.com/athomecomar/athome/backend/semantic/data"
	"github.com/athomecomar/athome/backend/semantic/data/value"
	"github.com/athomecomar/storeql"
)

type Attribute interface {
	storeql.Storable

	GetValueType() value.Type
	SetValueType(value.Type)

	GetName() string
	SetName(s string)

	GetCategoryId() uint64
	SetCategoryId(p uint64)

	NewData() (data.Attribute, error)
}
