package server

import (
	"context"
	"testing"
	"time"

	"github.com/athomecomar/athome/pb/pbagreement"
	"github.com/athomecomar/xerrors"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/status"
)

func TestServer_retrieve(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbagreement.RetrieveRequest
	}
	tests := []struct {
		name           string
		previousRecord *record
		args           args
		ttl            time.Duration
		randFn         randFunc
		want           *pbagreement.RetrieveResponse
		wantErr        xerrors.Code
	}{
		{
			name: "existing token",
			previousRecord: &record{
				UserId:         1,
				AgreementToken: "foo",
			},
			args: args{
				ctx: context.Background(),
				in:  &pbagreement.RetrieveRequest{UserId: 1},
			},
			ttl:  20 * time.Second,
			want: &pbagreement.RetrieveResponse{AgreementToken: "foo"},
		},
		{
			name: "non existant token",
			args: args{
				ctx: context.Background(),
				in:  &pbagreement.RetrieveRequest{UserId: 1},
			},
			randFn: func(int) (string, error) {
				return "foo", nil
			},
			want: &pbagreement.RetrieveResponse{AgreementToken: "foo"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			miniredis := connMiniredis(t)
			redis := redisCli(miniredis)
			if tt.previousRecord != nil {
				tt.previousRecord.save(t, redis, tt.ttl)
			}
			s := &Server{Redis: redis}

			got, err := s.retrieve(tt.args.ctx, tt.args.in, tt.randFn)
			if status.Code(err) != tt.wantErr {
				t.Errorf("Server.retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.AgreementTokenExpNs = uint64(tt.ttl.Nanoseconds())
			if tt.ttl == 0 {
				tt.want.AgreementTokenExpNs = got.AgreementTokenExpNs
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Server.retrieve() mismatch (-got +want): %s", diff)
			}
		})
	}
}
