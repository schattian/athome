package server

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/pbuser"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) FetchOnboarding(ctx context.Context, in *pbuser.FetchOnboardingRequest) (*pbuser.FetchOnboardingResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := connDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "connDB: %v", err)
	}
	defer db.Close()
	onboarding, err := fetchOnboardingByToken(ctx, db, in)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "onboarding with id %v not found", in.GetToken())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "fetchOnboardingByToken: %v", err)
	}

	return onboardingToFetchOnboardingResponse(onboarding), nil
}

func onboardingToFetchOnboardingResponse(o *ent.Onboarding) *pbuser.FetchOnboardingResponse {
	return &pbuser.FetchOnboardingResponse{
		Email:   string(o.Email),
		Name:    string(o.Name),
		Role:    string(o.Role),
		Surname: string(o.Surname),
		Stage:   int64(o.Stage),
	}
}

type tokenizable interface {
	GetToken() uint64
}

func fetchOnboardingByToken(ctx context.Context, db *sqlx.DB, in tokenizable) (*ent.Onboarding, error) {
	o := &ent.Onboarding{}
	row := db.QueryRowxContext(ctx, `SELECT * FROM onboardings WHERE id=$1`, in.GetToken())
	err := row.StructScan(o)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return o, nil
}
