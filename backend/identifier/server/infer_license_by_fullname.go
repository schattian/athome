package server

import (
	"context"
	"strings"

	"github.com/athomecomar/athome/backend/identifier/infer"
	"github.com/athomecomar/athome/pb/pbidentifier"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/xerrors"
	"github.com/spf13/afero"
	"google.golang.org/grpc/status"
)

func (s *Server) InferLicenseByFullname(ctx context.Context, c *semprov.Category, in *pbidentifier.InferByFullnameRequest) (*pbidentifier.InferLicenseResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	fs := afero.NewOsFs()
	return s.inferLicenseByFullname(ctx, fs, c, in)
}

func (s *Server) inferLicenseByFullname(_ context.Context, fs afero.Fs, category *semprov.Category, in *pbidentifier.InferByFullnameRequest) (*pbidentifier.InferLicenseResponse, error) {
	inferror, ok := infer.LicenseByFullnameByCategory[category]
	if !ok {
		return nil, status.Errorf(xerrors.InvalidArgument, "invalid category %v", category)
	}
	name, surname := strings.TrimSpace(in.GetName()), strings.TrimSpace(in.GetSurname())
	license, err := inferror(fs, name, surname)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "%v inferror by fullname returned: %v", category, err)
	}
	if license == 0 {
		return nil, status.Errorf(xerrors.NotFound, "couldn't infer license for fullname: %s %s", name, surname)
	}
	return &pbidentifier.InferLicenseResponse{License: license}, nil
}

func (s *Server) InferLicenseByFullnameMedic(ctx context.Context, in *pbidentifier.InferByFullnameRequest) (*pbidentifier.InferLicenseResponse, error) {
	return s.InferLicenseByFullname(ctx, semprov.Medic, in)
}
