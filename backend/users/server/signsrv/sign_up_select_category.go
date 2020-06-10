package signsrv

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/xpbsemantic"
	"github.com/athomecomar/athome/backend/users/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SignUpSelectCategory(ctx context.Context, in *pbusers.SignUpSelectCategoryRequest) (*emptypb.Empty, error) {
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

	return s.signUpSelectCategory(ctx, db, sem, in, o)
}

func retrieveLatestOnboarding(ctx context.Context, db *sqlx.DB, oid uint64) (*ent.Onboarding, error) {
	previous, err := retrieveOnboardingByToken(ctx, db, oid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "onboarding with id %v not found", oid)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "retrieveOnboardingByToken: %v", err)
	}
	return previous, nil
}

func (s *Server) signUpSelectCategory(ctx context.Context, db *sqlx.DB, sem xpbsemantic.CategoriesClient, in *pbusers.SignUpSelectCategoryRequest, onboarding *ent.Onboarding) (e *emptypb.Empty, err error) {
	onboarding.Stage = onboarding.Stage.Next(onboarding.Role)

	err = onboarding.MustStage(field.SelectCategory)
	if err != nil {
		return nil, status.Errorf(xerrors.OutOfRange, "MustStage: %v", err)
	}

	cat, err := sem.RetrieveCategory(ctx, &pbsemantic.RetrieveCategoryRequest{CategoryId: in.GetCategoryId()})
	if err != nil {
		return nil, err
	}
	if cat.GetChilds() != nil {
		return nil, status.Errorf(xerrors.Internal, "cannot use non-low level category named %v", cat.GetName())
	}

	code, err := onboarding.ValidateByStage(ctx, db)
	if err != nil {
		return nil, status.Errorf(code, "ValidateByStage: %v", err)
	}

	onboarding.CategoryId = in.GetCategoryId()
	err = storeql.UpdateIntoDB(ctx, db, onboarding)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "UpdateIntoDB: %v", err)
	}

	return &emptypb.Empty{}, nil
}
