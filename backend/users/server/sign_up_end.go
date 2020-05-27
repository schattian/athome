package server

import (
	"context"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func (s *Server) SignUpEnd(ctx context.Context, in *pbuser.SignUpEndRequest) (*pbuser.SignUpEndResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := connDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "connDB: %v", err)
	}
	defer db.Close()
	return s.signUpEnd(ctx, db, in)
}

func (s *Server) signUpEnd(ctx context.Context, db *sqlx.DB, in *pbuser.SignUpEndRequest) (*pbuser.SignUpEndResponse, error) {
	previous, err := fetchOnboardingByToken(ctx, db, in.GetOnboardingId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "fetchOnboardingByToken: %v", err)
	}
	if previous == nil {
		return nil, status.Errorf(xerrors.NotFound, "onboarding with id %v not found", in.GetOnboardingId())
	}

	onboarding := previous.Next()
	code, err := onboarding.MustStage(field.End)
	if err != nil {
		return nil, status.Errorf(code, "MustStage: %v", err)
	}

	code, err = onboarding.ValidateByStage(ctx, db)
	if err != nil {
		return nil, status.Errorf(code, "ValidateByStage: %v", err)
	}

	user := onboarding.ToUser()
	err = user.AssignPassword(in.GetPassword())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "AssignPassword: %v", err)
	}

	err = storeql.InsertIntoDB(ctx, db, user)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
	}

	signedUser, err := userToSignInUser(user)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "DeleteFromDB: %v", err)
	}

	err = storeql.DeleteFromDB(ctx, db, onboarding)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "DeleteFromDB: %v", err)
	}
	return &pbuser.SignUpEndResponse{User: signedUser}, nil
}
