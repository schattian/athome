package pbutil

import (
	"errors"
	"math"

	"github.com/athomecomar/athome/pb/pbshared"
)

func DiffTimeOfDay(t1, t2 *pbshared.TimeOfDay) int64 {
	a := t1.GetHour()*60 + t1.GetMinute()
	b := t2.GetHour()*60 + t2.GetMinute()
	return b - a
}

func RestTimeOfDay(t1 *pbshared.TimeOfDay, mins int64) (*pbshared.TimeOfDay, error) {
	hours := int64(math.Floor(float64(mins) / 60))
	if hours > t1.GetHour() {
		return nil, errors.New("dow time overflow")
	}
	minute := t1.GetMinute() - (mins - hours*60)
	if minute < 0 {
		hours += 1
		minute += 60
	}
	hour := t1.GetHour() - hours
	return &pbshared.TimeOfDay{Hour: hour, Minute: minute}, nil
}
