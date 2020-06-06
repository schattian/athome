package signsrv

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveOnboarding(ctx context.Context, in *pbusers.RetrieveOnboardingRequest) (*pbusers.RetrieveOnboardingResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	defer db.Close()
	return s.retrieveOnboarding(ctx, db, in)
}

func (s *Server) retrieveOnboarding(ctx context.Context, db *sqlx.DB, in *pbusers.RetrieveOnboardingRequest) (*pbusers.RetrieveOnboardingResponse, error) {
	onboarding, err := retrieveOnboardingByToken(ctx, db, in.GetOnboardingId())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "onboarding with id %v not found", in.GetOnboardingId())
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "retrieveOnboardingByToken: %v", err)
	}
	return onboardingToRetrieveOnboardingResponse(onboarding), nil
}

func onboardingToRetrieveOnboardingResponse(o *ent.Onboarding) *pbusers.RetrieveOnboardingResponse {
	return &pbusers.RetrieveOnboardingResponse{
		Email:   string(o.Email),
		Name:    string(o.Name),
		Role:    string(o.Role),
		Surname: string(o.Surname),
		Stage:   int64(o.Stage),
	}
}

func retrieveOnboardingByToken(ctx context.Context, db *sqlx.DB, token uint64) (*ent.Onboarding, error) {
	o := &ent.Onboarding{}
	row := db.QueryRowxContext(ctx, `SELECT * FROM onboardings WHERE id=$1`, token)
	err := row.StructScan(o)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return o, nil
}
