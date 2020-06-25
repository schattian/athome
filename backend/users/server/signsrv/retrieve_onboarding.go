package signsrv

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveOnboarding(ctx context.Context, in *pbusers.RetrieveOnboardingRequest) (*pbusers.Onboarding, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	defer db.Close()
	o, err := retrieveLatestOnboarding(ctx, db, in.GetOnboardingId())
	if err != nil {
		return nil, err
	}

	return s.retrieveOnboarding(ctx, db, o)
}

func (s *Server) retrieveOnboarding(ctx context.Context, db *sqlx.DB, onboarding *ent.Onboarding) (*pbusers.Onboarding, error) {
	return onboarding.ToPb(), nil
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
