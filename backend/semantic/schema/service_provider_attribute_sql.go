package schema

import "database/sql/driver"

func (sp *ServiceProviderAttributeSchema) GetId() uint64 {
	return sp.Id
}

func (sp *ServiceProviderAttributeSchema) SetId(id uint64) {
	sp.Id = id
}

func (sp *ServiceProviderAttributeSchema) SQLTable() string {
	return "service_proivder_attribute_schemas"
}

func (sp *ServiceProviderAttributeSchema) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":          sp.Id,
		"name":        sp.Name,
		"category_id": sp.CategoryId,
		"value_type":  sp.ValueType,
	}
}
