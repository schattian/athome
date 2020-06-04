package ent

import "database/sql/driver"

func (sp *ProductAttribute) GetId() uint64 {
	return sp.Id
}

func (sp *ProductAttribute) SetId(id uint64) {
	sp.Id = id
}

func (sp *ProductAttribute) SQLTable() string {
	return "product_attributes"
}

func (sp *ProductAttribute) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":          sp.Id,
		"name":        sp.Name,
		"category_id": sp.CategoryId,
		"value_type":  sp.ValueType,

		"bool_value":     sp.BoolValue,
		"string_value":   sp.StringValue,
		"int_64_value":   sp.Int64Value,
		"float_64_value": sp.Float64Value,

		"sl_string_value":   sp.SlStringValue,
		"sl_int_64_value":   sp.SlInt64Value,
		"sl_float_64_value": sp.SlFloat64Value,
	}
}
