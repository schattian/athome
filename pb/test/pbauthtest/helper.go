package pbauthtest

import (
	"context"

	pbauth "github.com/athomecomar/athome/pb/pbauth"
	gomock "github.com/golang/mock/gomock"
)

func NewCtrlFromRetrieve(ctrl *gomock.Controller, ctx context.Context, access string, userId uint64) pbauth.AuthClient {
	c := NewMockAuthClient(ctrl)
	c.EXPECT().
		RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access}).
		Return(&pbauth.RetrieveAuthenticationResponse{UserId: userId}, nil)
	return c
}
