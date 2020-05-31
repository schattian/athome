package server

import (
	"context"

	"github.com/athomecomar/athome/backend/identifier/pb/pbidentifier"
	"github.com/athomecomar/athome/backend/identifier/validate"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func (s *Server) ValidateLicensePsychologist(ctx context.Context, in *pbidentifier.ValidateLicenseRequest) (*pbidentifier.ValidateLicenseResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	return s.validateLicense(ctx, semprov.Psychologist, in)
}

func (s *Server) validateLicense(ctx context.Context, c semprov.Category, in *pbidentifier.ValidateLicenseRequest) (*pbidentifier.ValidateLicenseResponse, error) {
	verifier, ok := validate.ByCategory[c]
	if !ok {
		return nil, status.Errorf(xerrors.InvalidArgument, "invalid category %s", c)
	}
	valid, err := verifier(in.GetDni(), in.GetLicense())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "%s verifier returned: %v", c, err)
	}
	return &pbidentifier.ValidateLicenseResponse{Valid: valid}, nil
}
