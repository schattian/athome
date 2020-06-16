package ent

import "database/sql/driver"

func (u *Message) GetId() uint64 {
	return u.Id
}

func (u *Message) SetId(id uint64) {
	u.Id = id
}

func (u *Message) SQLTable() string {
	return "messages"
}

func (u *Message) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":          u.Id,
		"receiver_id": u.ReceiverId,
		"sender_id":   u.SenderId,
		"body":        u.Body,
		"created_at":  u.CreatedAt,
		"received_at": u.ReceivedAt,
		"seen_at":     u.SeenAt,
	}
}
