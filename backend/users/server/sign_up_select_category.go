package server

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SignUpSelectCategory(ctx context.Context, in *pbuser.SignUpSelectCategoryRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := connDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "connDB: %v", err)
	}
	defer db.Close()
	return s.signUpSelectCategory(ctx, db, in)
}

func (s *Server) signUpSelectCategory(ctx context.Context, db *sqlx.DB, in *pbuser.SignUpSelectCategoryRequest) (e *emptypb.Empty, err error) {
	previous, err := fetchOnboardingByToken(ctx, db, in.GetOnboardingId())
	if errors.Is(err, sql.ErrNoRows) {
		err = status.Errorf(xerrors.NotFound, "onboarding with id %v not found", in.GetOnboardingId())
		return
	}
	if err != nil {
		err = status.Errorf(xerrors.Internal, "fetchOnboardingByToken: %v", err)
		return
	}
	onboarding := previous.Next()

	code, err := onboarding.MustStage(field.SelectCategory)
	if err != nil {
		return nil, status.Errorf(code, "MustStage: %v", err)
	}

	code, err = onboarding.ValidateByStage(ctx, db)
	if err != nil {
		return nil, status.Errorf(code, "ValidateByStage: %v", err)
	}

	err = onboarding.SetCategory(in.GetCategoryName())
	if err != nil {
		err = status.Errorf(xerrors.Internal, "onboarding.SetCategory: %v", err)
		return
	}
	err = storeql.UpdateIntoDB(ctx, db, onboarding)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "UpdateIntoDB: %v", err)
	}

	return &emptypb.Empty{}, nil
}
