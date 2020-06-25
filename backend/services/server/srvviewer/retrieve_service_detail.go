package srvviewer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveService(ctx context.Context, in *pbservices.RetrieveServiceRequest) (*pbservices.Service, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.retrieveService(ctx, db, in)
}

func (s *Server) retrieveService(
	ctx context.Context,
	db *sqlx.DB,
	in *pbservices.RetrieveServiceRequest,
) (*pbservices.Service, error) {
	svc, err := ent.FindService(ctx, db, in.GetServiceId())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "FindService with id: %v", in.GetServiceId())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindService: %v", err)
	}
	return svc.ToPb(), nil
}
