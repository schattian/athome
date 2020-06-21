package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/address/ent"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveAddress(ctx context.Context, in *pbaddress.RetrieveAddressRequest) (*pbaddress.Address, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.retrieveAddress(ctx, db, in)
}

func (s *Server) MeasureDistance(ctx context.Context, in *pbaddress.MeasureDistanceRequest) (*pbaddress.MeasureDistanceResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.measureDistance(ctx, db, in)
}

func (s *Server) measureDistance(ctx context.Context, db *sqlx.DB, in *pbaddress.MeasureDistanceRequest) (*pbaddress.MeasureDistanceResponse, error) {
	addrs, err := ent.FindAddresses(ctx, db, in.GetAAddressId(), in.GetBAddressId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindAddresses: %v", err)
	}
	if len(addrs) != 2 {
		return nil, status.Errorf(xerrors.InvalidArgument, "FindAddresses with ids: %v and %v", in.GetAAddressId(), in.GetBAddressId())
	}

	a, b := addrs[0], addrs[1]
	haversine := a.DistanceHaversine(b)
	manhattan := a.DistanceManhattan(b)

	return &pbaddress.MeasureDistanceResponse{
		HaversineInKilometers: haversine,
		ManhattanInKilometers: manhattan,
	}, nil
}

func (s *Server) retrieveAddress(ctx context.Context, db *sqlx.DB, in *pbaddress.RetrieveAddressRequest) (*pbaddress.Address, error) {
	addr, err := ent.FindAddress(ctx, db, in.GetAddressId())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "FindAddress with id: %v", in.GetAddressId())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindAddress: %v", err)
	}
	return addr.ToPb(), nil
}
