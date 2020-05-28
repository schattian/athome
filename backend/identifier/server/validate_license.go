package server

import (
	"context"
	"errors"

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
	valid, err := scraper.VerifyLicense(semprov.Category(in.GetCategory()), in.GetDni(), in.GetLicense())
	if err != nil {
		if errors.Is(err, semerr.ErrProviderCategoryNotFound) {
			return nil, status.Errorf(xerrors.InvalidArgument, "invalid category %s: %v", in.GetCategory(), err)
		}
		// return nil, status.Error
	}
	return &pbidentifier.ValidateLicenseResponse{Valid: valid}, nil
}
