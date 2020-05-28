package server

import (
	"context"

	"github.com/athomecomar/athome/backend/identifier/pb/pbidentifier"
	"github.com/athomecomar/athome/backend/identifier/scraper"
	"github.com/athomecomar/semantic/semerr"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func (s *Server) ValidateLicense(ctx context.Context, in *pbidentifier.ValidateLicenseRequest) (*pbidentifier.ValidateLicenseResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	return s.validateLicense(ctx, in)
}

func (s *Server) validateLicense(ctx context.Context, in *pbidentifier.ValidateLicenseRequest) (*pbidentifier.ValidateLicenseResponse, error) {
	verifier, ok := scraper.VerifierByCategory[semprov.Category(in.GetCategory())]
	if !ok {
		return nil, status.Errorf(xerrors.InvalidArgument, "invalid category %s: %v", in.GetCategory(), semerr.ErrProviderCategoryNotFound)
	}
	valid, err := verifier(in.GetDni(), in.GetLicense())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "%s verifier returned: %v", in.GetCategory(), err)
	}
	return &pbidentifier.ValidateLicenseResponse{Valid: valid}, nil
}
