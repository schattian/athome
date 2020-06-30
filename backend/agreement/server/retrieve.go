package server

import (
	"context"

	"github.com/athomecomar/athome/backend/agreement/agreementconf"
	"github.com/athomecomar/athome/pb/pbagreement"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func (s *Server) Retrieve(ctx context.Context, in *pbagreement.RetrieveRequest) (*pbagreement.RetrieveResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	return s.retrieve(ctx, in)
}

func (s *Server) retrieve(ctx context.Context, in *pbagreement.RetrieveRequest) (*pbagreement.RetrieveResponse, error) {
	token, err := retrieveToken(ctx, s.Redis, in.GetUserId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "retrieveToken: %v", err)
	}
	return &pbagreement.RetrieveResponse{
		AgreementToken:      token,
		AgreementTokenExpNs: uint64(agreementconf.GetAGREEMENT_TOKEN_EXP().Nanoseconds()),
	}, nil
}
