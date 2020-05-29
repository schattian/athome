package server

import (
	"context"

	"github.com/athomecomar/athome/backend/identifier/pb/pbidentifier"
	"github.com/athomecomar/athome/backend/identifier/scraper"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/xerrors"
	"github.com/spf13/afero"
	"google.golang.org/grpc/status"
)

func (s *Server) InferLicenseByFullname(ctx context.Context, in *pbidentifier.InferLicenseByFullnameRequest) (*pbidentifier.InferLicenseByFullnameResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	return s.inferLicenseByFullname(ctx, in)
}

func (s *Server) inferLicenseByFullname(ctx context.Context, in *pbidentifier.InferLicenseByFullnameRequest) (*pbidentifier.InferLicenseByFullnameResponse, error) {
	inferror, ok := scraper.InferrorByFullnameByCategory[semprov.Category(in.GetCategory())]
	if !ok {
		return nil, status.Errorf(xerrors.InvalidArgument, "invalid category %s", in.GetCategory())
	}
	license, err := inferror(afero.NewOsFs(), in.GetName(), in.GetSurname())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "%s inferror by fullname returned: %v", in.GetCategory(), err)
	}
	if license == 0 {
		return nil, status.Errorf(xerrors.NotFound, "couldn't infer license for fullname: %s %s", in.GetName(), in.GetSurname())
	}
	return &pbidentifier.InferLicenseByFullnameResponse{License: license}, nil
}
