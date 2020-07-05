package shipping

import "database/sql/driver"

func (u *Shipping) GetId() uint64 {
	return u.Id
}

func (u *Shipping) SetId(id uint64) {
	u.Id = id
}

func (u *Shipping) SQLTable() string {
	return "shippings"
}

func (u *Shipping) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":                 u.Id,
		"user_id":            u.UserId,
		"event_id":           u.EventId,
		"shipping_method_id": u.ShippingMethodId,

		"order_price":               u.OrderPrice,
		"order_duration_in_minutes": u.OrderDurationInMinutes,

		"real_price":               u.RealPrice,
		"real_duration_in_minutes": u.RealDurationInMinutes,
	}
}
