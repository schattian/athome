package server

import (
	"context"
	"testing"

	"github.com/athomecomar/athome/backend/auth/pbauth"
	"github.com/athomecomar/xerrors"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/status"
)

func TestServer_retrieveAuthentication(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbauth.RetrieveAuthenticationRequest
	}
	tests := []struct {
		name           string
		previousRecord *variadicTokens
		args           args
		want           *pbauth.RetrieveAuthenticationResponse
		wantStatus     xerrors.Code
	}{
		{
			name:           "normal auth",
			previousRecord: gTokens.Valid,
			args: args{
				ctx: context.Background(),
				in:  &pbauth.RetrieveAuthenticationRequest{AccessToken: gTokens.Valid.AccessToken},
			},
			want: &pbauth.RetrieveAuthenticationResponse{
				UserId: gTokens.Valid.UserId,
			},
			wantStatus: xerrors.OK,
		},
		{
			name: "db is empty",
			args: args{
				ctx: context.Background(),
				in:  &pbauth.RetrieveAuthenticationRequest{AccessToken: gTokens.Valid.AccessToken},
			},
			wantStatus: xerrors.Unauthenticated,
		},

		{
			name:           "another user is authenticated",
			previousRecord: gTokens.Expired,
			args: args{
				ctx: context.Background(),
				in:  &pbauth.RetrieveAuthenticationRequest{AccessToken: gTokens.Valid.AccessToken},
			},
			wantStatus: xerrors.Unauthenticated,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			miniredis := connMiniredis(t)
			redis := redisCli(miniredis)
			if tt.previousRecord != nil {
				tt.previousRecord.save(t, redis)
			}
			s := &Server{Redis: redis}

			got, err := s.retrieveAuthentication(tt.args.ctx, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("Server.retrieveAuthentication() error = %v, wantStatus %v", err, tt.wantStatus)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Server.retrieveAuthentication() errored mismatch (-want +got): %s", diff)
			}

		})
	}
}
