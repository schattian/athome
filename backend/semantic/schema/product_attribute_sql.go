package schema

import "database/sql/driver"

func (sp *ProductAttributeSchema) GetId() uint64 {
	return sp.Id
}

func (sp *ProductAttributeSchema) SetId(id uint64) {
	sp.Id = id
}

func (sp *ProductAttributeSchema) SQLTable() string {
	return "product_attribute_schemas"
}

func (sp *ProductAttributeSchema) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":          sp.Id,
		"name":        sp.Name,
		"category_id": sp.CategoryId,
		"value_type":  sp.ValueType,
	}
}
