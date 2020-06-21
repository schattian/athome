package ent

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/athomecomar/athome/backend/services/ent/schedule"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbshared"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

type Availability struct {
	Id         uint64 `json:"id,omitempty"`
	CalendarId uint64 `json:"calendar_id,omitempty"`

	DayOfWeek   time.Weekday `json:"day_of_week,omitempty"`
	StartHour   int64        `json:"start_hour,omitempty"`
	EndHour     int64        `json:"end_hour,omitempty"`
	StartMinute int64        `json:"start_minute,omitempty"`
	EndMinute   int64        `json:"end_minute,omitempty"`
}

func availabilitiesContainingRange(ctx context.Context, db *sqlx.DB, dow time.Weekday, from, to *pbshared.TimeOfDay) ([]*Availability, error) {
	qr := `
        SELECT * FROM availabilities WHERE 
        dow = $1 
        AND
        (start_hour < $2 OR (start_hour = $2 AND start_minute <= $3))
        AND
        (end_hour > $4 OR (end_hour = $4 AND end_minute >= $5))
    `
	rows, err := db.QueryxContext(ctx, qr, dow, from.GetHour(), from.GetMinute(), to.GetHour(), to.GetMinute())
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()
	var cs []*Availability
	for rows.Next() {
		c := &Availability{}
		err := rows.StructScan(c)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func AvailabilityFromPb(in *pbservices.Availability) (*Availability, error) {
	dow, err := DayOfWeekByName(in.GetDow())
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "DayOfWeekByName: %v", err)
	}
	return &Availability{
		DayOfWeek: dow,

		CalendarId:  in.GetCalendarId(),
		StartHour:   in.GetStart().GetHour(),
		StartMinute: in.GetStart().GetMinute(),
		EndHour:     in.GetEnd().GetHour(),
		EndMinute:   in.GetEnd().GetMinute(),
	}, nil
}

func (av *Availability) ToPb() *pbservices.Availability {
	return &pbservices.Availability{
		Dow:        strings.ToLower(av.DayOfWeek.String()),
		CalendarId: av.CalendarId,
		Start:      &pbshared.TimeOfDay{Hour: av.StartHour, Minute: av.StartMinute},
		End:        &pbshared.TimeOfDay{Hour: av.EndHour, Minute: av.EndMinute},
	}
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
	av := &Availability{}
	row := storeql.Where(ctx, db, av, "id=$1", id)
	err := row.StructScan(av)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return av, nil
}

func FindEvent(ctx context.Context, db *sqlx.DB, id uint64) (*Event, error) {
	e := &Event{}
	row := storeql.Where(ctx, db, &Event{}, "id=$1", id)
	err := row.StructScan(e)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return e, nil
}

func AvailabilitiesToTimeables(avs []*Availability) (as []schedule.Scheduleable) {
	for _, av := range avs {
		as = append(as, av)
	}
	return
}

func (av *Availability) GetDayOfWeek() time.Weekday { return av.DayOfWeek }

func (av *Availability) GetStartHour() int64   { return av.StartHour }
func (av *Availability) GetStartMinute() int64 { return av.StartMinute }

func (av *Availability) GetEndHour() int64   { return av.EndHour }
func (av *Availability) GetEndMinute() int64 { return av.EndMinute }
