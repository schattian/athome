package ent

import "database/sql/driver"

func (u *Event) GetId() uint64 {
	return u.Id
}

func (u *Event) SetId(id uint64) {
	u.Id = id
}

func (u *Event) SQLTable() string {
	return "events"
}

func (u *Event) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":           u.Id,
		"claimant_id":  u.ClaimantId,
		"order_id":     u.OrderId,
		"calendar_id":  u.CalendarId,
		"is_confirmed": u.IsConfirmed,
		"day_of_week":  u.DayOfWeek,
		"start_hour":   u.StartHour,
		"start_minute": u.StartMinute,
		"end_hour":     u.EndHour,
		"end_minute":   u.EndMinute,
	}
}
