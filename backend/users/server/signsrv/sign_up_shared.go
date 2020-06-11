package signsrv

import (
	"context"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SignUpShared(ctx context.Context, in *pbusers.SignUpSharedRequest) (*emptypb.Empty, error) {
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

	return s.signUpShared(ctx, db, in, o)
}

func (s *Server) signUpShared(ctx context.Context, db *sqlx.DB, in *pbusers.SignUpSharedRequest, previous *ent.Onboarding) (*emptypb.Empty, error) {
	onboarding := signUpSharedRequestToOnboarding(previous, in)
	onboarding.Stage = onboarding.Stage.Next(onboarding.Role)

	err := onboarding.MustStage(field.Shared)
	if err != nil {
		return nil, status.Errorf(xerrors.OutOfRange, "MustStage: %v", err)
	}

	code, err := onboarding.ValidateByStage(ctx, db)
	if err != nil {
		return nil, status.Errorf(code, "ValidateByStage: %v", err)
	}

	err = storeql.UpdateIntoDB(ctx, db, onboarding)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "UpdateIntoDB: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func signUpSharedRequestToOnboarding(prev *ent.Onboarding, in *pbusers.SignUpSharedRequest) *ent.Onboarding {
	prev.Surname = field.Surname(in.GetSurname())
	prev.Name = field.Name(in.GetName())
	prev.Email = field.Email(in.GetEmail())
	return prev
}
