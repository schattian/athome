package signsrv

import (
	"context"

	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/xerrors"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func (s *Server) SignUpStart(ctx context.Context, in *pbusers.SignUpStartRequest) (*pbusers.SignUpStartResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	defer db.Close()
	return s.signUpStart(ctx, db, in)
}

func (s *Server) signUpStart(ctx context.Context, db *sqlx.DB, in *pbusers.SignUpStartRequest) (*pbusers.SignUpStartResponse, error) {
	onboarding := signUpStartRequestToOnboarding(in).Next()
	code, err := onboarding.ValidateByStage(ctx, db)
	if err != nil {
		return nil, status.Errorf(code, "ValidateByStage: %v", err)
	}
	err = storeql.InsertIntoDB(ctx, db, onboarding)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
	}
	return onboardingToSignUpStartResponse(onboarding), nil
}

func signUpStartRequestToOnboarding(in *pbusers.SignUpStartRequest) *ent.Onboarding {
	return &ent.Onboarding{Role: field.Role(in.GetRole())}
}

func onboardingToSignUpStartResponse(o *ent.Onboarding) *pbusers.SignUpStartResponse {
	return &pbusers.SignUpStartResponse{
		OnboardingId: o.Id,
	}
}
