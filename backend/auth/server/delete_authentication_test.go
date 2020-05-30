package server

import (
	"context"
	"testing"

	"github.com/athomecomar/athome/backend/auth/pb/pbauth"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func TestServer_deleteAuthentication(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbauth.DeleteAuthenticationRequest
	}
	tests := []struct {
		name           string
		previousRecord *variadicTokens
		args           args
		wantStatus     xerrors.Code
	}{
		{
			name:           "normal auth",
			previousRecord: gTokens.Valid,
			args: args{
				ctx: context.Background(),
				in:  &pbauth.DeleteAuthenticationRequest{AccessToken: gTokens.Valid.AccessToken},
			},
			wantStatus: xerrors.OK,
		},
		{
			name:           "expired token given",
			previousRecord: gTokens.Valid,
			args: args{
				ctx: context.Background(),
				in:  &pbauth.DeleteAuthenticationRequest{AccessToken: gTokens.Expired.AccessToken},
			},
			wantStatus: xerrors.InvalidArgument,
		},
		{
			name: "db is empty",
			args: args{
				ctx: context.Background(),
				in:  &pbauth.DeleteAuthenticationRequest{AccessToken: gTokens.Valid.AccessToken},
			},
			wantStatus: xerrors.Unauthenticated,
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

			_, err := s.deleteAuthentication(tt.args.ctx, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.deleteAuthentication() error = %v, wantStatus %v", err, tt.wantStatus)
			}
		})
	}
}
