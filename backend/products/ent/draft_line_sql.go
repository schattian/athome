package ent

import "database/sql/driver"

func (u *DraftLine) GetId() uint64 {
	return u.Id
}

func (u *DraftLine) SetId(id uint64) {
	u.Id = id
}

func (u *DraftLine) SQLTable() string {
	return "draft_lines"
}

func (u *DraftLine) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":       u.Id,
		"draft_id": u.DraftId,

		"category_id": u.CategoryId,
		"title":       u.Title,
	}
}
