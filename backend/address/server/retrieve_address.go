package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/athomecomar/athome/backend/address/ent"
	"github.com/athomecomar/athome/backend/address/pb/pbaddress"
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
