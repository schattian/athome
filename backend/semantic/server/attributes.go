package server

import (
	"github.com/athomecomar/athome/backend/semantic/data"
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

func DataAttributeToPbAttributeData(d data.Attribute) *pbsemantic.AttributeData {
	return &pbsemantic.AttributeData{
		SchemaId: d.GetSchemaId(),
		Values:   d.GetValue().Strings(),
	}
}
