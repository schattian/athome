package srvcalendars

import (
	"context"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/ent/schedule"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbusers"
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
	auth, authCloser, err := server.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()
	users, usersCloser, err := server.ConnUsers(ctx)
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
	u, err := users.RetrieveUser(ctx, &pbusers.RetrieveUserRequest{UserId: claimantId})
	if err != nil {
		return nil, err
	}
	if role := u.GetUser().GetRole(); role != "consumer" {
		return nil, status.Errorf(xerrors.PermissionDenied, "user with role %v cant create events", role)
	}
	c, err := ent.FindCalendar(ctx, db, in.GetCalendarId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindCalendar: %v", err)
	}
	avs, err := c.Availabilities(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Availabilities: %v", err)
	}
	e, err := ent.EventFromPb(in.GetEvent())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "EventFromPb: %v", err)
	}
	e.ClaimantId = claimantId
	if !schedule.CompareWithSlice(schedule.IsContained, e, ent.AvailabilitiesToTimeables(avs)...) {
		return nil, status.Errorf(xerrors.InvalidArgument, "there is no availability to store the event")
	}
	err = storeql.InsertIntoDB(ctx, db, e)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}
	resp := &pbservices.CreateEventResponse{EventId: e.Id, Event: e.ToPb()}
	return resp, nil
}
