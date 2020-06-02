package signsrv

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbidentifier"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/semantic/semerr"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SignUpIdentification(ctx context.Context, in *pbuser.SignUpIdentificationRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	defer db.Close()

	conn, err := grpc.Dial(userconf.GetAUTH_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, userconf.GetAUTH_ADDR())
	}

	c := pbidentifier.NewIdentifierClient(conn)
	defer conn.Close()
	return s.signUpIdentification(ctx, db, c, in)
}

func (s *Server) signUpIdentification(ctx context.Context, db *sqlx.DB, c pbidentifier.IdentifierClient, in *pbuser.SignUpIdentificationRequest) (e *emptypb.Empty, err error) {
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

	var oi *ent.OnboardingIdentification
	switch onboarding.Role {
	case field.Merchant:
	case field.ServiceProvider:
		oi, err = signUpIdentificationServiceProvider(ctx, c, in, onboarding)
	}
	if err != nil {
		return nil, err
	}

	err = storeql.InsertIntoDB(ctx, db, oi)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func signUpIdentificationServiceProvider(
	ctx context.Context,
	c pbidentifier.IdentifierClient,
	in *pbuser.SignUpIdentificationRequest,
	o *ent.Onboarding,
) (oi *ent.OnboardingIdentification, err error) {
	switch o.Category {
	case semprov.Medic.Name:
		oi, err = signUpIdentificationMedic(ctx, c, in.GetMedic())
	case semprov.Lawyer.Name:
		oi, err = signUpIdentificationLawyer(ctx, c, in.GetLawyer())
	default:
		err = semerr.ErrProviderCategoryNotFound
	}
	return
}

func signUpIdentificationMedic(
	ctx context.Context,
	c pbidentifier.IdentifierClient,
	in *pbuser.SignUpIdentificationRequest_Medic,
) (*ent.OnboardingIdentification, error) {
	resp, err := c.InferLicenseByFullnameMedic(ctx, nameableToInferByFullnameRequest(in))
	if err != nil {
		return nil, err
	}
	return &ent.OnboardingIdentification{License: resp.GetLicense(), Name: field.Name(in.GetName()), Surname: field.Surname(in.GetSurname())}, nil
}

func signUpIdentificationLawyer(
	ctx context.Context,
	c pbidentifier.IdentifierClient,
	in *pbuser.SignUpIdentificationRequest_Lawyer,
) (*ent.OnboardingIdentification, error) {
	resp, err := c.InferTomeAndFolioByFullnameLawyer(ctx, nameableToInferByFullnameRequest(in))
	if err != nil {
		return nil, err
	}
	return &ent.OnboardingIdentification{Tome: resp.GetTome(), Folio: resp.GetFolio()}, nil
}

type nameable interface {
	GetName() string
	GetSurname() string
}

func nameableToInferByFullnameRequest(n nameable) *pbidentifier.InferByFullnameRequest {
	return &pbidentifier.InferByFullnameRequest{Name: n.GetName(), Surname: n.GetSurname()}
}
