package ent

import "database/sql/driver"

func (u *Notification) GetId() uint64 {
	return u.Id
}

func (u *Notification) SetId(id uint64) {
	u.Id = id
}

func (u *Notification) SQLTable() string {
	return "notifications"
}

func (u *Notification) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":           u.Id,
		"priority":     u.Priority,
		"user_id":      u.UserId,
		"entity_table": u.EntityTable,
		"entity_id":    u.EntityId,
		"created_at":   u.CreatedAt.Time,
		"received_at":  u.ReceivedAt.Time,
		"seen_at":      u.SeenAt.Time,
	}
}
