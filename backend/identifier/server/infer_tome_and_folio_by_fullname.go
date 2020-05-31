package server

import (
	"context"
	"strings"

	"github.com/athomecomar/athome/backend/identifier/infer"
	"github.com/athomecomar/athome/backend/identifier/pb/pbidentifier"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/xerrors"
	"github.com/spf13/afero"
	"google.golang.org/grpc/status"
)

func (s *Server) InferTomeAndFolioByFullname(ctx context.Context, category *semprov.Category, in *pbidentifier.InferByFullnameRequest) (*pbidentifier.InferTomeAndFolioResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	fs := afero.NewOsFs()
	return s.inferTomeAndFolioByFullname(ctx, fs, category, in)
}

func (s *Server) inferTomeAndFolioByFullname(ctx context.Context, fs afero.Fs, category *semprov.Category, in *pbidentifier.InferByFullnameRequest) (*pbidentifier.InferTomeAndFolioResponse, error) {
	inferror, ok := infer.TomeAndFolioByFullnameByCategory[category]
	if !ok {
		return nil, status.Errorf(xerrors.InvalidArgument, "invalid category %s", category)
	}
	name, surname := strings.TrimSpace(in.GetName()), strings.TrimSpace(in.GetSurname())
	tome, folio, err := inferror(fs, name, surname)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "%s inferror by fullname returned: %v", category, err)
	}
	if tome == 0 || folio == 0 {
		return nil, status.Errorf(xerrors.NotFound, "couldn't infer tome or folio for fullname: %s %s", name, surname)
	}
	return &pbidentifier.InferTomeAndFolioResponse{Tome: tome, Folio: folio}, nil
}

func (s *Server) InferTomeAndFolioByFullnameLawyer(ctx context.Context, in *pbidentifier.InferByFullnameRequest) (*pbidentifier.InferTomeAndFolioResponse, error) {
	return s.InferTomeAndFolioByFullname(ctx, semprov.Lawyer, in)
}

func (s *Server) InferTomeAndFolioByFullnameAttorney(ctx context.Context, in *pbidentifier.InferByFullnameRequest) (*pbidentifier.InferTomeAndFolioResponse, error) {
	return s.InferTomeAndFolioByFullname(ctx, semprov.Attorney, in)
}
