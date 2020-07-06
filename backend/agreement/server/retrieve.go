package server

import (
	"context"

	"github.com/athomecomar/athome/pb/pbagreement"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func (s *Server) Retrieve(ctx context.Context, in *pbagreement.RetrieveRequest) (*pbagreement.RetrieveResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	return s.retrieve(ctx, in, randInt)
}

func (s *Server) retrieve(ctx context.Context, in *pbagreement.RetrieveRequest, randFn randFunc) (*pbagreement.RetrieveResponse, error) {
	token, err := retrieveOrCreateToken(ctx, s.Redis, in.GetUserId(), randFn)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "retrieveOrCreateToken: %v", err)
	}
	ttl, err := ttl(ctx, s.Redis, in.GetUserId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "ttl: %v", err)
	}
	return &pbagreement.RetrieveResponse{
		AgreementToken:      token,
		AgreementTokenExpNs: uint64(ttl.Nanoseconds()),
	}, nil
}
