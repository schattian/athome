package srvregister

import (
	"context"

	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveRegistry(ctx context.Context, in *pbservices.RetrieveRegistryRequest) (*pbservices.RegistryDetail, error) {
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

	return s.retrieveRegistry(ctx, db, server.GetUserFromAccessToken(auth, in.GetAccessToken()))
}

func (s *Server) retrieveRegistry(ctx context.Context, db *sqlx.DB, authFn server.AuthFunc) (*pbservices.RegistryDetail, error) {
	reg, err := retrieveRegistryByUser(ctx, db, authFn)
	if err != nil {
		return nil, err
	}
	var c *pbservices.Calendar
	var a map[uint64]*pbservices.Availability
	if reg.CalendarId != 0 {
		cal, err := reg.Calendar(ctx, db)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "Calendar: %v", err)
		}
		c = cal.ToPb()

		avs, err := cal.Availabilities(ctx, db)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "Availabilities: %v", err)
		}
		a = make(map[uint64]*pbservices.Availability)
		for _, av := range avs {
			a[av.Id] = av.ToPb()
		}
	}
	return &pbservices.RegistryDetail{
		RegistryId:     reg.Id,
		Registry:       reg.ToPb(),
		Calendar:       c,
		Availabilities: a,
	}, nil
}
