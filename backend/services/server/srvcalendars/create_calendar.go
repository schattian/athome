package srvcalendars

import (
	"context"
	"strings"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/pb/pbauth"
	"github.com/athomecomar/athome/backend/services/pb/pbservices"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCalendar(ctx context.Context, in *pbservices.CreateCalendarRequest) (*pbservices.CreateCalendarResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	auth, authCloser, err := server.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	return s.createCalendar(ctx, db, auth, server.GetUserFromAccessToken, in)
}

func (s *Server) createCalendar(ctx context.Context, db *sqlx.DB, auth pbauth.AuthClient, authFn server.AuthFunc, in *pbservices.CreateCalendarRequest) (*pbservices.CreateCalendarResponse, error) {
	userId, err := authFn(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	body := in.GetBody()
	var availabilities []*ent.Availability
	for _, av := range body.GetAvailabilities() {
		availability, err := pbAvailabilityDataToAvailability(av)
		if err != nil {
			return nil, err
		}
		availabilities = append(availabilities, availability)
	}
	if ent.CheckOverlappingPairwise(availabilities) {
		return nil, status.Errorf(xerrors.InvalidArgument, "trying to perform a self-overlapping of availabilities")
	}

	avs, err := ent.AvailabilitiesByUserGroup(ctx, db, userId, body.GetGroupId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "AvailabilitiesByUser: %v", err)
	}
	for _, av := range avs {
		if av.CheckOverlappingPairwise(availabilities) {
			return nil, status.Errorf(xerrors.InvalidArgument, "tried to overlap availability")
		}
	}

	calendar := pbCalendarDataToCalendar(body, userId)
	err = storeql.InsertIntoDB(ctx, db, calendar)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "calendar InsertIntoDB: %v", err)
	}
	var storables []storeql.Storable
	for _, av := range availabilities {
		av.CalendarId = calendar.Id
		storables = append(storables, av)
	}
	err = storeql.InsertIntoDB(ctx, db, storables...)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "avs InsertIntoDB: %v", err)
	}

	resp := &pbservices.CreateCalendarResponse{
		CalendarId: calendar.Id,
		Data: &pbservices.CalendarData{
			GroupId: calendar.GroupId,
			Name:    calendar.Name,
		},
	}
	for _, av := range availabilities {
		resp.Data.Availabilities = append(resp.Data.Availabilities, availabilityToPbAvailabilityData(av))
	}
	return resp, nil
}

func availabilityToPbAvailabilityData(av *ent.Availability) *pbservices.AvailabilityData {
	return &pbservices.AvailabilityData{
		Dow:   strings.ToLower(av.DayOfWeek.String()),
		Start: &pbservices.TimeOfDay{Hour: av.StartHour, Minute: av.StartMinute},
		End:   &pbservices.TimeOfDay{Hour: av.EndHour, Minute: av.EndMinute},
	}
}

func pbAvailabilityDataToAvailability(in *pbservices.AvailabilityData) (*ent.Availability, error) {
	in.GetDow()

	dow, err := ent.DayOfWeekByName(in.GetDow())
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "DayOfWeekByName: %v", err)
	}
	return &ent.Availability{
		DayOfWeek: dow,

		StartHour:   in.GetStart().GetHour(),
		StartMinute: in.GetStart().GetMinute(),
		EndHour:     in.GetEnd().GetHour(),
		EndMinute:   in.GetEnd().GetMinute(),
	}, nil
}

func pbCalendarDataToCalendar(in *pbservices.CalendarData, uid uint64) *ent.Calendar {
	return &ent.Calendar{
		Name:    in.GetName(),
		GroupId: in.GetGroupId(),
		UserId:  uid,
	}
}

func calendarToPbCalendarData(c *ent.Calendar) *pbservices.CalendarData {
	return &pbservices.CalendarData{
		Name:    c.Name,
		GroupId: c.GroupId,
	}
}
