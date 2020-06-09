package ent

import "database/sql/driver"

func (u *Availability) GetId() uint64 {
	return u.Id
}

func (u *Availability) SetId(id uint64) {
	u.Id = id
}

func (u *Availability) SQLTable() string {
	return "availabilities"
}

func (u *Availability) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":           u.Id,
		"calendar_id":  u.CalendarId,
		"day_of_week":  u.DayOfWeek,
		"start_hour":   u.StartHour,
		"start_minute": u.StartMinute,
		"end_hour":     u.EndHour,
		"end_minute":   u.EndMinute,
	}
}
