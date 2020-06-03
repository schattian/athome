package ent

import "database/sql/driver"

func (sp *MerchantCategory) GetId() uint64 {
	return sp.Id
}

func (sp *MerchantCategory) SetId(id uint64) {
	sp.Id = id
}

func (sp *MerchantCategory) SQLTable() string {
	return "merchant_categories"
}

func (sp *MerchantCategory) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":        sp.Id,
		"name":      sp.Name,
		"parent_id": sp.ParentId,
	}
}
