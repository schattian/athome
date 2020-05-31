package server

import (
	"context"
	"strings"

	"github.com/athomecomar/athome/backend/identifier/pb/pbidentifier"
	"github.com/athomecomar/athome/backend/identifier/scraper"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/xerrors"
	"github.com/spf13/afero"
	"google.golang.org/grpc/status"
)

func (s *Server) InferLicenseByFullnamePsychologist(ctx context.Context, in *pbidentifier.InferByFullnameRequest) (*pbidentifier.InferLicenseResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	fs := afero.NewOsFs()
	return s.inferLicenseByFullname(ctx, fs, semprov.Psychologist, in)
}

func (s *Server) inferLicenseByFullname(ctx context.Context, fs afero.Fs, category semprov.Category, in *pbidentifier.InferByFullnameRequest) (*pbidentifier.InferLicenseResponse, error) {
	inferror, ok := scraper.InferrorByFullnameByCategory[category]
	if !ok {
		return nil, status.Errorf(xerrors.InvalidArgument, "invalid category %s", category)
	}
	name, surname := strings.TrimSpace(in.GetName()), strings.TrimSpace(in.GetSurname())
	license, err := inferror(fs, name, surname)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "%s inferror by fullname returned: %v", category, err)
	}
	if license == 0 {
		return nil, status.Errorf(xerrors.NotFound, "couldn't infer license for fullname: %s %s", name, surname)
	}
	return &pbidentifier.InferLicenseResponse{License: license}, nil
}
