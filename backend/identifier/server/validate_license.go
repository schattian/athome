package server

import (
	"context"

	"github.com/athomecomar/athome/backend/identifier/validate"
	"github.com/athomecomar/athome/pb/pbidentifier"
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

func (s *Server) validateLicense(_ context.Context, c *semprov.Category, in *pbidentifier.ValidateLicenseRequest) (*pbidentifier.ValidateLicenseResponse, error) {
	verifier, ok := validate.ByCategory[c]
	if !ok {
		return nil, status.Errorf(xerrors.InvalidArgument, "invalid category %v", c)
	}
	valid, err := verifier(in.GetDni(), in.GetLicense())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "%v verifier returned: %v", c, err)
	}
	return &pbidentifier.ValidateLicenseResponse{Valid: valid}, nil
}
