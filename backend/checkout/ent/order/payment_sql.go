package order

import "database/sql/driver"

func (u *Payment) GetId() uint64 {
	return u.Id
}

func (u *Payment) SetId(id uint64) {
	u.Id = id
}

func (u *Payment) SQLTable() string {
	return "payments"
}

func (u *Payment) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":                u.Id,
		"user_id":           u.UserId,
		"payment_method_id": u.PaymentMethodId,
		"card_id":           u.CardId,
		"entity_id":         u.EntityId,
		"entity_table":      u.EntityTable,
		"amount":            u.Amount,
		"created_at":        u.CreatedAt,
		"updated_at":        u.UpdatedAt,
		"installments":      u.Installments,
	}
}
