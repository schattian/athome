package pbauthtest

import (
	"context"

	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client struct {
	Create   *pbauth.CreateAuthenticationResponse
	Retrieve *pbauth.RetrieveAuthenticationResponse

	CreateErr   error
	RetrieveErr error
	DeleteErr   error
}

func (c Client) DeleteAuthentication(ctx context.Context, in *pbauth.DeleteAuthenticationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	return &emptypb.Empty{}, c.DeleteErr
}

func (c Client) CreateAuthentication(ctx context.Context, in *pbauth.CreateAuthenticationRequest, opts ...grpc.CallOption) (*pbauth.CreateAuthenticationResponse, error) {
	return c.Create, c.CreateErr
}

func (c Client) RetrieveAuthentication(ctx context.Context, in *pbauth.RetrieveAuthenticationRequest, opts ...grpc.CallOption) (*pbauth.RetrieveAuthenticationResponse, error) {
	return c.Retrieve, c.RetrieveErr
}
