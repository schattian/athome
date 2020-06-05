package server

import (
	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/semantic/schema"
)

func AttributeToPbsemanticAttribute(c schema.Attribute) *pbsemantic.AttributeSchema {
	return &pbsemantic.AttributeSchema{
		Name:       c.GetName(),
		Id:         c.GetId(),
		ValueType:  string(c.GetValueType()),
		CategoryId: c.GetCategoryId(),
	}
}
