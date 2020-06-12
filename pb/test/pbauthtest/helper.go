package pbauthtest

import (
	"context"
	"testing"

	pbauth "github.com/athomecomar/athome/pb/pbauth"
	gomock "github.com/golang/mock/gomock"
)

func NewCtrlFromRetrieve(t *testing.T, ctx context.Context, access string, userId uint64) *gomock.Controller {
	t.Helper()
	ctrl := gomock.NewController(t)
	NewMockAuthClient(ctrl).
		EXPECT().
		RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access}).
		Return(&pbauth.RetrieveAuthenticationResponse{UserId: userId}, nil)
	return ctrl
}
