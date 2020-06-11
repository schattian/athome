package ent

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/athomecomar/athome/backend/services/ent/schedule"
	"github.com/athomecomar/athome/pb/pbservices"
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

func AvailabilityFromPb(in *pbservices.Availability) (*Availability, error) {
	in.GetDow()

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
		Start:      &pbservices.TimeOfDay{Hour: av.StartHour, Minute: av.StartMinute},
		End:        &pbservices.TimeOfDay{Hour: av.EndHour, Minute: av.EndMinute},
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
	row := storeql.Where(ctx, db, &Availability{}, "id=$1", id)
	av := &Availability{}
	err := row.StructScan(av)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return av, nil
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
