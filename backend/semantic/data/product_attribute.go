package data

import (
	"github.com/athomecomar/athome/backend/semantic/data/value"
	"github.com/pkg/errors"
)

type ProductAttributeData struct {
	Id       uint64
	SchemaId uint64

	BoolValue    *value.Bool    `json:"bool_value,omitempty"`
	StringValue  *value.String  `json:"string_value,omitempty"`
	Int64Value   *value.Int64   `json:"int_64_value,omitempty"`
	Float64Value *value.Float64 `json:"float_64_value,omitempty"`

	SlStringValue  *value.SlString  `json:"sl_string_value,omitempty"`
	SlInt64Value   *value.SlInt64   `json:"sl_int_64_value,omitempty"`
	SlFloat64Value *value.SlFloat64 `json:"sl_float_64_value,omitempty"`
}

func NewProductAttributeData(t value.Type) (*ProductAttributeData, error) {
	pc := &ProductAttributeData{}
	switch t {
	case value.TypeInt64:
		pc.Int64Value = value.NilInt64
	case value.TypeBool:
		pc.BoolValue = value.NilBool
	case value.TypeFloat64:
		pc.Float64Value = value.NilFloat64
	case value.TypeString:
		pc.StringValue = value.NilString
	case value.TypeSlInt64:
		pc.SlInt64Value = value.NilSlInt64
	case value.TypeSlString:
		pc.SlStringValue = value.NilSlString
	case value.TypeSlFloat64:
		pc.SlFloat64Value = value.NilSlFloat64
	default:
		return nil, errors.New("invalid type given " + string(t))
	}
	return pc, nil
}

func (pc *ProductAttributeData) GetSchemaId() uint64 {
	return pc.SchemaId
}

func (pc *ProductAttributeData) SetSchemaId(i uint64) {
	pc.SchemaId = i
}

func (pc *ProductAttributeData) SetValueStrings(values ...string) error {
	if len(values) == 0 {

	}
	return nil
}

func (pc *ProductAttributeData) SetValue(v interface{}) error {
	var err error
	for _, value := range pc.values() {
		err = value.SetValue(v)
		if err == nil {
			break
		}
	}
	val := pc.GetValue()
	if val == nil || val.IsNil() {
		return errors.New("value cant be inserted")
	}
	return nil
}

func (pc *ProductAttributeData) GetValue() value.Value {
	for _, val := range pc.values() {
		if !val.IsNil() {
			return val
		}
	}
	return nil
}

func (pc *ProductAttributeData) values() []value.Value {
	return []value.Value{
		pc.BoolValue,
		pc.Float64Value,
		pc.StringValue,
		pc.Int64Value,

		pc.SlFloat64Value,
		pc.SlInt64Value,
		pc.SlStringValue,
	}
}
