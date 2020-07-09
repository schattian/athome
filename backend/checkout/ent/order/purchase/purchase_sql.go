package purchase

import "database/sql/driver"

func (u *Purchase) GetId() uint64 {
	return u.Id
}

func (u *Purchase) SetId(id uint64) {
	u.Id = id
}

func (u *Purchase) SQLTable() string {
	return string(u.OrderClass())
}

func (u *Purchase) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":                     u.Id,
		"user_id":                u.UserId,
		"distance_in_kilometers": u.DistanceInKilometers,
		"dest_address_id":        u.DestAddressId,
		"src_address_id":         u.SrcAddressId,
		"shipping_id":            u.ShippingId,
		"merchant_id":            u.MerchantId,
		"created_at":             u.CreatedAt,
		"updated_at":             u.UpdatedAt,
		"items":                  u.Items,
	}
}
