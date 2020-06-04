package server

import (
	"github.com/athomecomar/athome/backend/semantic/ent"
	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
)

func AttributeToPbsemanticAttribute(c ent.Attribute) *pbsemantic.Attribute {
	return &pbsemantic.Attribute{
		Name:       c.GetName(),
		Id:         c.GetId(),
		ValueType:  string(c.GetValueType()),
		CategoryId: c.GetCategoryId(),
	}
}
