package order

import "database/sql/driver"

func (u *PurchaseStateChange) GetId() uint64 {
	return u.Id
}

func (u *PurchaseStateChange) SetId(id uint64) {
	u.Id = id
}

func (u *PurchaseStateChange) SQLTable() string {
	return "purchase_state_changes"
}

func (u *PurchaseStateChange) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":         u.Id,
		"name":       u.Name,
		"created_at": u.CreatedAt,
		"stage":      u.Stage,
		"entity_id":  u.EntityId,
	}
}
