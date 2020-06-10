package signsrv

import (
	"context"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func (s *Server) SignUpEnd(ctx context.Context, in *pbusers.SignUpEndRequest) (*pbusers.SignUpEndResponse, error) {
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

	return s.signUpEnd(ctx, db, in, o)
}

func (s *Server) signUpEnd(ctx context.Context, db *sqlx.DB, in *pbusers.SignUpEndRequest, onboarding *ent.Onboarding) (*pbusers.SignUpEndResponse, error) {
	onboarding.Stage = onboarding.Stage.Next(onboarding.Role)

	err := onboarding.MustStage(field.End)
	if err != nil {
		return nil, status.Errorf(xerrors.OutOfRange, "MustStage: %v", err)
	}

	code, err := onboarding.ValidateByStage(ctx, db)
	if err != nil {
		return nil, status.Errorf(code, "ValidateByStage: %v", err)
	}

	user, _, err := onboarding.Close(ctx, db, in.GetPassword())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "onboarding.Close: %v", err)
	}
	signedUser, err := userToSignInUser(user)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "userToSignInUser: %v", err)
	}
	return &pbusers.SignUpEndResponse{User: signedUser}, nil
}
