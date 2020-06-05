package data

import "github.com/athomecomar/athome/backend/semantic/data/value"

type ProductAttributeData struct {
	Id uint64

	BoolValue    value.Bool    `json:"bool_value,omitempty"`
	StringValue  value.String  `json:"string_value,omitempty"`
	Int64Value   value.Int64   `json:"int_64_value,omitempty"`
	Float64Value value.Float64 `json:"float_64_value,omitempty"`

	SlStringValue  value.SlString  `json:"sl_string_value,omitempty"`
	SlInt64Value   value.SlInt64   `json:"sl_int_64_value,omitempty"`
	SlFloat64Value value.SlFloat64 `json:"sl_float_64_value,omitempty"`
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
