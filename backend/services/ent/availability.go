package ent

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Availability struct {
	Id         uint64 `json:"id,omitempty"`
	CalendarId uint64 `json:"calendar_id,omitempty"`

	DayOfWeek   time.Weekday `json:"day_of_week,omitempty"`
	StartHour   int64
	EndHour     int64
	StartMinute int64
	EndMinute   int64
}

var days = [...]time.Weekday{
	time.Sunday,
	time.Monday,
	time.Tuesday,
	time.Wednesday,
	time.Thursday,
	time.Friday,
	time.Saturday,
}

func TimeFromHourMinute(h, m int) time.Time {
	return time.Date(0, 0, 0, h, m, 0, 0, nil)
}

func DayOfWeekByName(s string) (time.Weekday, error) {
	for _, d := range days {
		if strings.EqualFold(d.String(), s) {
			return d, nil
		}
	}
	return 0, fmt.Errorf("couldn't find day of week by name %s", s)
}

func FindAvailability(ctx context.Context, db *sqlx.DB, id uint64) (*Availability, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM availabilities WHERE id=$1`, id)
	av := &Availability{}
	err := row.StructScan(av)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return av, nil
}

func (av *Availability) StartAbs() int64 {
	return absTs(av.StartHour, av.StartMinute)
}

func absTs(h, m int64) int64 {
	return h*60 + m
}

func (av *Availability) EndAbs() int64 {
	return absTs(av.EndHour, av.EndMinute)
}

func (av *Availability) CheckOverlappingPairwise(avs []*Availability) bool {
	for _, avi := range avs {
		if av.Overlaps(avi) {
			return true
		}
	}
	return false
}

func CheckOverlappingPairwise(avs []*Availability) bool {
	for i, avi := range avs {
		for j, avj := range avs {
			if j == i {
				continue
			}
			if avi.Overlaps(avj) {
				return true
			}
		}
	}
	return false
}

func (av *Availability) Overlaps(a *Availability) bool {
	if av.DayOfWeek != a.DayOfWeek {
		return false
	}

	startAbs, endAbs := av.StartAbs(), av.EndAbs()
	if startAbs >= endAbs {
		return false
	}
	if endAbs <= startAbs {
		return false
	}

	return true
}
