package sm

import "database/sql/driver"

func (u *StateChange) GetId() uint64 {
	return u.Id
}

func (u *StateChange) SetId(id uint64) {
	u.Id = id
}

func (u *StateChange) SQLTable() string {
	return "state_changes"
}

func (u *StateChange) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":           u.Id,
		"name":         u.Name,
		"created_at":   u.CreatedAt,
		"stage":        u.Stage,
		"entity_id":    u.EntityId,
		"entity_table": u.EntityTable,
	}
}
