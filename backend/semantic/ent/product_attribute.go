package ent

import (
	"github.com/athomecomar/athome/backend/semantic/ent/value"
)

type ProductAttribute struct {
	Id uint64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ValueType value.Type `json:"value_type,omitempty"`

	BoolValue    value.Bool    `json:"bool_value,omitempty"`
	StringValue  value.String  `json:"string_value,omitempty"`
	Int64Value   value.Int64   `json:"int_64_value,omitempty"`
	Float64Value value.Float64 `json:"float_64_value,omitempty"`

	SlStringValue  value.SlString  `json:"sl_string_value,omitempty"`
	SlInt64Value   value.SlInt64   `json:"sl_int_64_value,omitempty"`
	SlFloat64Value value.SlFloat64 `json:"sl_float_64_value,omitempty"`

	CategoryId uint64 `json:"category_id,omitempty"`
}

func (pc *ProductAttribute) GetValue() value.Value {
	for _, val := range pc.values() {
		if !val.IsNil() {
			return val
		}
	}
	return nil
}

func (pc *ProductAttribute) SetValueType(v value.Type) {
	pc.ValueType = v
}

func (pc *ProductAttribute) GetValueType() value.Type {
	return pc.ValueType
}

func (pc *ProductAttribute) GetName() string {
	return pc.Name
}

func (pc *ProductAttribute) SetName(s string) {
	pc.Name = s
}

func (pc *ProductAttribute) GetCategoryId() uint64 {
	return pc.CategoryId
}

func (pc *ProductAttribute) SetCategoryId(p uint64) {
	pc.CategoryId = p
}

func (pc *ProductAttribute) values() []value.Value {
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
