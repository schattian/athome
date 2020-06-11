package schema

import (
	"github.com/athomecomar/athome/backend/semantic/data"
	"github.com/athomecomar/athome/backend/semantic/data/value"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/storeql"
)

type Attribute interface {
	storeql.Storable

	GetValueType() value.Type
	SetValueType(value.Type)

	GetName() string
	SetName(s string)

	RetrieveCategoryId() uint64
	SetCategoryId(p uint64)

	NewData() (data.Attribute, error)
}

func AttributeToPb(c Attribute) *pbsemantic.AttributeSchema {
	return &pbsemantic.AttributeSchema{
		Name:       c.GetName(),
		ValueType:  string(c.GetValueType()),
		CategoryId: c.RetrieveCategoryId(),
	}
}
