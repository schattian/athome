package payment

import "database/sql/driver"

func (u *PaymentStateChange) GetId() uint64 {
	return u.Id
}

func (u *PaymentStateChange) SetId(id uint64) {
	u.Id = id
}

func (u *PaymentStateChange) SQLTable() string {
	return "payment_state_changes"
}

func (u *PaymentStateChange) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":         u.Id,
		"name":       u.Name,
		"created_at": u.CreatedAt,
		"stage":      u.Stage,
		"entity_id":  u.EntityId,
	}
}
