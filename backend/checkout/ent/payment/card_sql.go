package payment

import "database/sql/driver"

func (u *Card) GetId() uint64 {
	return u.Id
}

func (u *Card) SetId(id uint64) {
	u.Id = id
}

func (u *Card) SQLTable() string {
	return "cards"
}

func (u *Card) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":               u.Id,
		"user_id":          u.UserId,
		"number_hash":      u.NumberHash,
		"last_four_digits": u.LastFourDigits,
		"cvv_hash":         u.CVVHash,
		"expiry_month":     u.ExpiryMonth,
		"expiry_year":      u.ExpiryYear,
		"holder_dni":       u.HolderDNI,
		"holder_name":      u.HolderName,
	}
}
