package signsrv

import (
	"context"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/xpbsemantic"
	"github.com/athomecomar/athome/backend/users/pb/pbidentifier"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/semantic/semerr"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SignUpIdentification(ctx context.Context, in *pbusers.SignUpIdentificationRequest) (*emptypb.Empty, error) {
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

	iden, idenCloser, err := server.ConnIdentifier(ctx)
	if err != nil {
		return nil, err
	}
	defer idenCloser()
	sem, semCloser, err := server.ConnCategories(ctx, o.Role)
	if err != nil {
		return nil, err
	}
	defer semCloser()

	return s.signUpIdentification(ctx, db, sem, iden, in, o)
}

func (s *Server) signUpIdentification(
	ctx context.Context,
	db *sqlx.DB,
	sem xpbsemantic.CategoriesClient,
	iden pbidentifier.IdentifierClient,
	in *pbusers.SignUpIdentificationRequest,
	onboarding *ent.Onboarding,
) (e *emptypb.Empty, err error) {
	onboarding.Stage = onboarding.Stage.Next(onboarding.Role)
	cat, err := onboarding.Category(ctx, sem)
	if err != nil {
		return nil, err
	}

	var oi *ent.OnboardingIdentification
	switch onboarding.Role {
	case field.Merchant:
	case field.ServiceProvider:
		oi, err = signUpIdentificationServiceProvider(ctx, iden, in, cat.GetIdentificationTemplate())
	}
	if err != nil {
		return
	}

	err = storeql.InsertIntoDB(ctx, db, oi)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
	}
	err = storeql.UpdateIntoDB(ctx, db, onboarding)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "UpdateIntoDB: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func signUpIdentificationServiceProvider(
	ctx context.Context,
	c pbidentifier.IdentifierClient,
	in *pbusers.SignUpIdentificationRequest,
	identificationTemplate string,
) (oi *ent.OnboardingIdentification, err error) {
	switch identificationTemplate {
	case semprov.Medic.Name:
		oi, err = signUpIdentificationMedic(ctx, c, in.GetMedic())
	case semprov.Attorney.Name:
		oi, err = signUpIdentificationAttorney(ctx, c, in.GetAttorney())
	case semprov.Lawyer.Name:
		oi, err = signUpIdentificationLawyer(ctx, c, in.GetLawyer())
	case semprov.Psychologist.Name:
		oi, err = signUpIdentificationPsychologist(ctx, c, in.GetDni(), in.GetPsychologist())
	default:
		err = status.Error(xerrors.InvalidArgument, semerr.ErrProviderCategoryNotFound.Error())
	}
	if err != nil {
		return
	}
	return
}

func signUpIdentificationPsychologist(
	ctx context.Context,
	c pbidentifier.IdentifierClient,
	dni uint64,
	in *pbusers.SignUpIdentificationRequest_Psychologist,
) (*ent.OnboardingIdentification, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	resp, err := c.ValidateLicensePsychologist(ctx, &pbidentifier.ValidateLicenseRequest{License: in.GetLicense(), Dni: dni})
	if err != nil {
		return nil, err
	}
	if !resp.Valid {
		return nil, status.Error(xerrors.InvalidArgument, "license doesnt match with dni")
	}
	return &ent.OnboardingIdentification{License: in.GetLicense()}, nil
}

func signUpIdentificationMedic(
	ctx context.Context,
	c pbidentifier.IdentifierClient,
	in *pbusers.SignUpIdentificationRequest_Medic,
) (*ent.OnboardingIdentification, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	resp, err := c.InferLicenseByFullnameMedic(ctx, nameableToInferByFullnameRequest(in))
	if err != nil {
		return nil, err
	}
	return &ent.OnboardingIdentification{License: resp.GetLicense(), Name: field.Name(in.GetName()), Surname: field.Surname(in.GetSurname())}, nil
}

func signUpIdentificationAttorney(
	ctx context.Context,
	c pbidentifier.IdentifierClient,
	in *pbusers.SignUpIdentificationRequest_Attorney,
) (*ent.OnboardingIdentification, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	resp, err := c.InferTomeAndFolioByFullnameAttorney(ctx, nameableToInferByFullnameRequest(in))
	if err != nil {
		return nil, err
	}
	return &ent.OnboardingIdentification{Tome: resp.GetTome(), Folio: resp.GetFolio()}, nil
}

func signUpIdentificationLawyer(
	ctx context.Context,
	c pbidentifier.IdentifierClient,
	in *pbusers.SignUpIdentificationRequest_Lawyer,
) (*ent.OnboardingIdentification, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
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
