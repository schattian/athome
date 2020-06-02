package ent

import "database/sql/driver"

func (o *Identification) GetId() uint64 {
	return o.Id
}

func (o *Identification) SetId(id uint64) {
	o.Id = id
}

func (o *Identification) SQLTable() string {
	return "identifications"
}

func (o *Identification) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":       o.Id,
		"user_id":  o.UserId,
		"verified": o.Verified,
		"dni":      o.DNI,
		"name":     o.Name,
		"surname":  o.Surname,
		"license":  o.License,
		"tome":     o.Tome,
		"folio":    o.Folio,
		"cue":      o.CUE,
	}
}
