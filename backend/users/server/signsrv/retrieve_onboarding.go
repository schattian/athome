package signsrv

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/internal/xpbsemantic"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveOnboarding(ctx context.Context, in *pbusers.RetrieveOnboardingRequest) (*pbusers.OnboardingDetail, error) {
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

	sem, semCloser, err := server.ConnCategories(ctx, o.Role)
	if err != nil {
		return nil, err
	}
	defer semCloser()

	return s.retrieveOnboarding(ctx, db, sem, o)
}

func (s *Server) retrieveOnboarding(ctx context.Context, db *sqlx.DB, sem xpbsemantic.CategoriesClient, onboarding *ent.Onboarding) (*pbusers.OnboardingDetail, error) {
	cat, err := onboarding.Category(ctx, sem)
	if err != nil {
		return nil, err
	}
	iden, err := onboarding.Identification(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Identification: %v", err)
	}
	return &pbusers.OnboardingDetail{
		Onboarding:     onboarding.ToPb(),
		Category:       server.PbSemanticCategoryToPbUserCategory(cat),
		Identification: iden.ToPb(),
	}, nil
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
