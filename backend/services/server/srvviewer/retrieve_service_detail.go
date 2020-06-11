package srvviewer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/pb/pbaddress"
	"github.com/athomecomar/athome/backend/services/pb/pbservices"
	"github.com/athomecomar/athome/backend/services/pb/pbusers"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveServiceDetail(ctx context.Context, in *pbservices.RetrieveServiceDetailRequest) (*pbservices.RetrieveServiceDetailResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	addr, addrCloser, err := server.ConnAddress(ctx)
	if err != nil {
		return nil, err
	}
	defer addrCloser()
	users, usersCloser, err := server.ConnUsers(ctx)
	if err != nil {
		return nil, err
	}
	defer usersCloser()
	return s.retrieveServiceDetail(ctx, db, addr, users, in)
}

func (s *Server) retrieveServiceDetail(
	ctx context.Context,
	db *sqlx.DB,
	addr pbaddress.AddressesClient,
	users pbusers.ViewerClient,
	in *pbservices.RetrieveServiceDetailRequest,
) (*pbservices.RetrieveServiceDetailResponse, error) {
	svc, err := ent.FindService(ctx, db, in.GetServiceId())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "FindService with id: %v", in.GetServiceId())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindService: %v", err)
	}
	c, err := server.RetrieveCalendarDetail(ctx, db, svc.CalendarId)
	if err != nil {
		return nil, err
	}
	user, err := svc.User(ctx, users)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "svc.User: %v", err)
	}
	address, err := svc.Address(ctx, addr)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "svc.Address: %v", err)
	}

	resp := &pbservices.RetrieveServiceDetailResponse{
		Service:  svc.ToPb(),
		Address:  address,
		User:     user,
		Calendar: c,
	}
	return resp, nil
}
