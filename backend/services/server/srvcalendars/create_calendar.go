package srvcalendars

import (
	"context"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/ent/schedule"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbutil"
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
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	return s.createCalendar(ctx, db, server.GetUserFromAccessToken(auth, in.GetAccessToken()), in)
}

func (s *Server) createCalendar(ctx context.Context, db *sqlx.DB, authFn server.AuthFunc, in *pbservices.CreateCalendarRequest) (*pbservices.CreateCalendarResponse, error) {
	userId, err := authFn(ctx)
	if err != nil {
		return nil, err
	}
	var availabilities []*ent.Availability
	for _, av := range in.GetAvailabilities() {
		availability, err := ent.AvailabilityFromPb(av)
		if err != nil {
			return nil, err
		}
		availabilities = append(availabilities, availability)
	}
	availabilitiesTimer := ent.AvailabilitiesToTimeables(availabilities)

	if schedule.ComparePairwise(schedule.NotNullIntersection, availabilitiesTimer...) {
		return nil, status.Errorf(xerrors.InvalidArgument, "trying to perform a self-overlapping of availabilities")
	}

	pbCalendar := in.GetCalendar()
	avs, err := ent.AvailabilitiesByUserGroup(ctx, db, userId, pbCalendar.GetGroupId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "AvailabilitiesByUser: %v", err)
	}
	for _, av := range avs {
		if schedule.CompareWithSlice(schedule.NotNullIntersection, av, availabilitiesTimer...) {
			return nil, status.Errorf(xerrors.InvalidArgument, "tried to overlap availability")
		}
	}
	calendar := ent.CalendarFromPb(pbCalendar, userId)

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
		Calendar:   calendar.ToPb(),
	}
	resp.Availabilities = make(map[uint64]*pbservices.Availability)
	for _, av := range availabilities {
		resp.Availabilities[av.Id] = av.ToPb()
	}
	return resp, nil
}
