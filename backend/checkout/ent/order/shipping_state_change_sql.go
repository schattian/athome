package order

import "database/sql/driver"

func (u *ShippingStateChange) GetId() uint64 {
	return u.Id
}

func (u *ShippingStateChange) SetId(id uint64) {
	u.Id = id
}

func (u *ShippingStateChange) SQLTable() string {
	return "shipping_state_changes"
}

func (u *ShippingStateChange) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":          u.Id,
		"name":        u.Name,
		"created_at":  u.CreatedAt,
		"stage":       u.Stage,
		"shipping_id": u.ShippingId,
	}
}
