package server

import (
	"context"
	"log"

	"github.com/athomecomar/storeql"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pbuser"
	"github.com/athomecomar/xerrors"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func (s *Server) SignUpStart(ctx context.Context, in *pbuser.SignUpStartRequest) (*pbuser.SignUpStartResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := connDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "connDB: %v", err)
	}
	defer db.Close()

	onboarding := signUpStartRequestToOnboarding(in).Next()

	code, err := onboarding.ValidateByStage(ctx, db)
	if err != nil {
		return nil, status.Errorf(code, "ValidateByStage: %v", err)
	}

	err = storeql.InsertIntoDB(ctx, db, onboarding)
	log.Println(onboarding)
	log.Println(err)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
	}
	return onboardingToSignUpStartResponse(onboarding), nil
}

func signUpStartRequestToOnboarding(in *pbuser.SignUpStartRequest) *ent.Onboarding {
	return &ent.Onboarding{Role: field.Role(in.GetRole())}
}

func onboardingToSignUpStartResponse(o *ent.Onboarding) *pbuser.SignUpStartResponse {
	return &pbuser.SignUpStartResponse{
		Token: o.Id,
	}
}
