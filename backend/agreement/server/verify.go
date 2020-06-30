package server

import (
	"context"

	"github.com/athomecomar/athome/pb/pbagreement"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Verify(ctx context.Context, in *pbagreement.VerifyRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	return s.verify(ctx, in)
}

func (s *Server) verify(ctx context.Context, in *pbagreement.VerifyRequest) (*emptypb.Empty, error) {
	token, err := retrieveToken(ctx, s.Redis, in.GetAgreedUserId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "retrieveToken: %v", err)
	}
	if token != in.GetAgreementToken() {
		return nil, status.Error(xerrors.PermissionDenied, "invalid token given")
	}
	_, err = createToken(ctx, s.Redis, in.GetAgreedUserId(), randString)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "createToken: %v", err)
	}
	return &emptypb.Empty{}, nil
}
