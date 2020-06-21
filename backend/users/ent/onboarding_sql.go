package ent

import "database/sql/driver"

func (o *Onboarding) GetId() uint64 {
	return o.Id
}

func (o *Onboarding) SetId(id uint64) {
	o.Id = id
}

func (o *Onboarding) SQLTable() string {
	return "onboardings"
}

func (o *Onboarding) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":          o.Id,
		"email":       o.Email,
		"category_id": o.CategoryId,
		"address_id":  o.AddressId,
		"stage":       o.Stage,
		"role":        o.Role,
		"name":        o.Name,
		"surname":     o.Surname,
	}
}
