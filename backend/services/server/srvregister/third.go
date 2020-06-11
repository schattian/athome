package srvregister

import (
	"context"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/ent/stage"
	"github.com/athomecomar/athome/backend/services/pb/pbservices"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) Third(ctx context.Context, in *pbservices.ThirdRequest) (*pbservices.ThirdResponse, error) {
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

	reg, err := retrieveRegistryByUser(ctx, db, auth, in.GetAccessToken(), server.GetUserFromAccessToken)
	if err != nil {
		return nil, err
	}
	authCloser()

	return s.third(ctx, db, in, reg)
}

func (s *Server) third(ctx context.Context, db *sqlx.DB, in *pbservices.ThirdRequest, reg *ent.Registry) (*pbservices.ThirdResponse, error) {
	err := mustStage(reg.Stage, stage.Third)
	if err != nil {
		return nil, err
	}
	calendar, err := ent.FindCalendar(ctx, db, in.GetBody().GetCalendarId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindCalendar: %v", err)
	}
	if calendar.UserId != reg.UserId {
		return nil, status.Errorf(xerrors.PermissionDenied, "cannot use calendar with id: %v", calendar.Id)
	}

	reg = applyThirdRequestToRegistry(in.GetBody(), reg)

	svc := reg.ToService()
	err = storeql.InsertIntoDB(ctx, db, svc)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
	}
	err = storeql.DeleteFromDB(ctx, db, reg)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "DeleteFromDB: %v", err)
	}

	return &pbservices.ThirdResponse{
		Service:   svc.ToPb(),
		ServiceId: svc.Id,
	}, nil
}

func applyThirdRequestToRegistry(f *pbservices.ThirdRequest_Body, r *ent.Registry) *ent.Registry {
	r.CalendarId = f.GetCalendarId()
	return r
}
