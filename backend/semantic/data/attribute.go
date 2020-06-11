package data

import (
	"github.com/athomecomar/athome/backend/semantic/data/value"
	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/storeql"
	"github.com/pkg/errors"
)

type Attribute interface {
	storeql.Storable

	GetSchemaId() uint64
	SetSchemaId(uint64)

	GetUserId() uint64
	SetUserId(uint64)

	GetValue() value.Value
	SetValue(interface{}) error
}

func MustUserId(d Attribute, uid uint64) error {
	userId := d.GetUserId()
	if userId != 0 && uid != userId {
		return errors.New("unauthorised rsc")
	}
	if uid == userId {
		return nil
	}
	d.SetUserId(uid)
	return nil
}

func AttributeToPb(d Attribute) *pbsemantic.AttributeData {
	return &pbsemantic.AttributeData{
		SchemaId: d.GetSchemaId(),
		Values:   d.GetValue().Strings(),
	}
}
