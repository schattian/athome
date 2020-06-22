package srvcalendars

import (
	"context"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/ent/schedule"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateEvent(ctx context.Context, in *pbservices.CreateEventRequest) (*pbservices.CreateEventResponse, error) {
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
	users, usersCloser, err := pbutil.ConnUsersViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer usersCloser()

	return s.createEvent(ctx, db, users, server.GetUserFromAccessToken(auth, in.GetAccessToken()), in)
}

func (s *Server) createEvent(
	ctx context.Context,
	db *sqlx.DB,
	users pbusers.ViewerClient,
	authFn server.AuthFunc,
	in *pbservices.CreateEventRequest,
) (*pbservices.CreateEventResponse, error) {
	claimantId, err := authFn(ctx)
	if err != nil {
		return nil, err
	}
	e, err := ent.EventFromPb(in.GetEvent(), claimantId, in.GetCalendarId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "EventFromPb: %v", err)
	}
	claimant, err := users.RetrieveUser(ctx, &pbusers.RetrieveUserRequest{UserId: e.ClaimantId})
	if err != nil {
		return nil, err
	}

	return insertEvent(ctx, db, claimant.GetUser(), e)
}

func insertEvent(ctx context.Context, db *sqlx.DB, claimant *pbusers.User, event *ent.Event) (*pbservices.CreateEventResponse, error) {
	if role := claimant.GetRole(); role != "consumer" {
		return nil, status.Errorf(xerrors.PermissionDenied, "user with role %v cant create events", role)
	}
	c, err := ent.FindCalendar(ctx, db, event.CalendarId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindCalendar: %v", err)
	}
	avs, err := c.Availabilities(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Availabilities: %v", err)
	}
	if !schedule.CompareWithSlice(schedule.IsContained, event, ent.AvailabilitiesToTimeables(avs)...) {
		return nil, status.Errorf(xerrors.InvalidArgument, "there is no availability to store the event")
	}
	err = storeql.InsertIntoDB(ctx, db, event)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}
	resp := &pbservices.CreateEventResponse{EventId: event.Id, Event: event.ToPb()}
	return resp, nil
}
