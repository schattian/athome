package ent

import "database/sql/driver"

func (sp *ProductCategory) GetId() uint64 {
	return sp.Id
}

func (sp *ProductCategory) SetId(id uint64) {
	sp.Id = id
}

func (sp *ProductCategory) SQLTable() string {
	return "product_categories"
}

func (sp *ProductCategory) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":        sp.Id,
		"name":      sp.Name,
		"parent_id": sp.ParentId,
	}
}
