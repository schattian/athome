package data

import "database/sql/driver"

func (sp *ProductAttributeData) GetId() uint64 {
	return sp.Id
}

func (sp *ProductAttributeData) SetId(id uint64) {
	sp.Id = id
}

func (sp *ProductAttributeData) SQLTable() string {
	return "product_attribute_datas"
}

func (sp *ProductAttributeData) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":        sp.Id,
		"schema_id": sp.SchemaId,
		"user_id":   sp.UserId,

		"entity_id":    sp.EntityId,
		"entity_table": sp.EntityTable,

		"bool_value":     sp.BoolValue,
		"string_value":   sp.StringValue,
		"int_64_value":   sp.Int64Value,
		"float_64_value": sp.Float64Value,

		"sl_string_value":   sp.SlStringValue,
		"sl_int_64_value":   sp.SlInt64Value,
		"sl_float_64_value": sp.SlFloat64Value,
	}
}
