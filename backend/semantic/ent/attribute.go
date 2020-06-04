package ent

import "github.com/athomecomar/athome/backend/semantic/ent/field"

type Attribute interface {
	GetIsMultivalued() bool

	GetValues() []interface{}
	SetValues(v ...interface{})

	GetValueType() field.ValueType
	SetValueType(v field.ValueType)

	GetName() string
	SetName(s string)

	GetCategoryId() uint64
	SetCategoryId(p uint64)
}
