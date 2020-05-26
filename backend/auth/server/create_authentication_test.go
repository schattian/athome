package server

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/athomecomar/athome/backend/auth/pbauth"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func TestServer_createAuthentication(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbauth.CreateAuthenticationRequest
	}
	tests := []struct {
		name           string
		previousRecord *variadicTokens
		args           args
		userId         uint64
		wantStatus     xerrors.Code
	}{
		{
			name: "basic sign",
			args: args{
				ctx: context.Background(),
				in:  &pbauth.CreateAuthenticationRequest{SignToken: gTokens.Valid.SignToken},
			},
			userId:     gTokens.Valid.UserId,
			wantStatus: xerrors.OK,
		},
		{
			name:           "already signed, but expired",
			previousRecord: gTokens.Expired,
			args: args{
				ctx: context.Background(),
				in:  &pbauth.CreateAuthenticationRequest{SignToken: gTokens.Valid.SignToken},
			},
			userId:     gTokens.Valid.UserId,
			wantStatus: xerrors.OK,
		},
		{
			name:           "already signed, valid",
			previousRecord: gTokens.Valid,
			args: args{
				ctx: context.Background(),
				in:  &pbauth.CreateAuthenticationRequest{SignToken: gTokens.Valid.SignToken},
			},
			userId:     gTokens.Valid.UserId,
			wantStatus: xerrors.OK,
		},
		{
			name: "expired sign token",
			args: args{
				ctx: context.Background(),
				in:  &pbauth.CreateAuthenticationRequest{SignToken: gTokens.ExpiredSign.SignToken},
			},
			userId:     gTokens.Valid.UserId,
			wantStatus: xerrors.InvalidArgument,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			miniredis := connMiniredis(t)
			redis := redisCli(miniredis)
			if tt.previousRecord != nil {
				tt.previousRecord.save(t, redis)
			}
			s := &Server{Redis: redis}

			got, err := s.createAuthentication(tt.args.ctx, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("Server.createAuthentication() error = %v, wantStatus %v", err, tt.wantStatus)
				return
			}
			isCreated := err == nil
			retrieve, err := s.retrieveAuthentication(tt.args.ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: got.GetAccessToken()})
			if err != nil {
				if isCreated {
					t.Errorf("Server.createAuthentication() error = %v, when desired successful creation", err)
				}
				return
			}
			if retrieve.UserId != tt.userId {
				t.Errorf("Server.createAuthentication() mismatch retrieved userId after sign")
			}
		})
	}
}

func connMiniredis(t *testing.T) *miniredis.Miniredis {
	t.Helper()
	miniredis, err := miniredis.Run()
	if err != nil {
		t.Fatalf("miniredis failed to conn: %v", err)
	}

	return miniredis
}
