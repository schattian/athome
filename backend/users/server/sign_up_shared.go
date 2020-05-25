package server

import (
	"context"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pbuser"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func (s *Server) SignUpShared(ctx context.Context, in *pbuser.SignUpSharedRequest) (*pbuser.SignUpSharedResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := connDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "connDB: %v", err)
	}
	defer db.Close()
	return s.signUpShared(ctx, db, in)
}

func (s *Server) signUpShared(ctx context.Context, db *sqlx.DB, in *pbuser.SignUpSharedRequest) (*pbuser.SignUpSharedResponse, error) {
	previous, err := fetchOnboardingByToken(ctx, db, in.GetOnboardingId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "fetchOnboardingByToken: %v", err)
	}

	onboarding := signUpSharedRequestToOnboarding(previous, in).Next()

	code, err := onboarding.MustStage(field.Shared)
	if err != nil {
		return nil, status.Errorf(code, "MustStage: %v", err)
	}

	code, err = onboarding.ValidateByStage(ctx, db)
	if err != nil {
		return nil, status.Errorf(code, "ValidateByStage: %v", err)
	}

	err = storeql.UpdateIntoDB(ctx, db, onboarding)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "UpdateIntoDB: %v", err)
	}

	return onboardingToSignUpSharedResponse(onboarding), nil
}

func signUpSharedRequestToOnboarding(prev *ent.Onboarding, in *pbuser.SignUpSharedRequest) *ent.Onboarding {
	prev.Surname = field.Surname(in.GetSurname())
	prev.Name = field.Name(in.GetName())
	prev.Email = field.Email(in.GetEmail())
	return prev
}

func onboardingToSignUpSharedResponse(o *ent.Onboarding) *pbuser.SignUpSharedResponse {
	return &pbuser.SignUpSharedResponse{
		OnboardingId: o.Id,
	}
}
