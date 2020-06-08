package ent

import "database/sql/driver"

func (u *Product) GetId() uint64 {
	return u.Id
}

func (u *Product) SetId(id uint64) {
	u.Id = id
}

func (u *Product) SQLTable() string {
	return "products"
}

func (u *Product) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":      u.Id,
		"user_id": u.UserId,

		"category_id": u.CategoryId,
		"title":       u.Title,

		"price": u.Price,
		"stock": u.Stock,

		"image_ids": u.ImageIds,
	}
}
