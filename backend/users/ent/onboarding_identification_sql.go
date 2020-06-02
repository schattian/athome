package ent

import "database/sql/driver"

func (o *OnboardingIdentification) GetId() uint64 {
	return o.Id
}

func (o *OnboardingIdentification) SetId(id uint64) {
	o.Id = id
}

func (o *OnboardingIdentification) SQLTable() string {
	return "onboarding_identifications"
}

func (o *OnboardingIdentification) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":            o.Id,
		"onboarding_id": o.OnboardingId,
		"dni":           o.DNI,
		"name":          o.Name,
		"surname":       o.Surname,
		"license":       o.License,
		"tome":          o.Tome,
		"folio":         o.Folio,
		"cue":           o.CUE,
	}
}
