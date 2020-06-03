package ent

import "database/sql/driver"

func (sp *ServiceProviderCategory) GetId() uint64 {
	return sp.Id
}

func (sp *ServiceProviderCategory) SetId(id uint64) {
	sp.Id = id
}

func (sp *ServiceProviderCategory) SQLTable() string {
	return "service_provider_categories"
}

func (sp *ServiceProviderCategory) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":        sp.Id,
		"name":      sp.Name,
		"parent_id": sp.ParentId,
	}
}
