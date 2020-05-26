package server

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/users/pbuser"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func TestServer_signUpStart(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbuser.SignUpStartRequest
	}
	tests := []struct {
		name       string
		args       args
		execStub   *sqlassist.QueryStubber
		want       *pbuser.SignUpStartResponse
		wantStatus xerrors.Code
	}{
		{
			name: "basic consumer",
			args: args{
				ctx: context.Background(),
				in:  onboardingToSignUpStartRequest(t, gOnboardings.Consumers.Foo),
			},
			execStub: &sqlassist.QueryStubber{
				Expect: "INSERT INTO onboardings", Rows: sqlmock.NewRows([]string{"id"}).AddRow(gOnboardings.Consumers.Foo.Id),
			},
			want:       &pbuser.SignUpStartResponse{OnboardingId: gOnboardings.Consumers.Foo.Id},
			wantStatus: xerrors.OK,
		},
		{
			name: "consumer qr err",
			args: args{
				ctx: context.Background(),
				in:  onboardingToSignUpStartRequest(t, gOnboardings.Consumers.Foo),
			},
			execStub: &sqlassist.QueryStubber{
				Err:    errors.New("foo"),
				Expect: "INSERT INTO onboardings", Rows: sqlmock.NewRows([]string{"id"}).AddRow(gOnboardings.Consumers.Foo.Id),
			},
			wantStatus: xerrors.Internal,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock := sqlhelp.MockDB(t)
			if tt.execStub != nil {
				tt.execStub.Stub(mock)
			}

			s := &Server{}
			got, err := s.signUpStart(tt.args.ctx, db, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.signUpStart() error = %v, wantStatus %v", err, tt.wantStatus)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Server.signUpStart() errored mismatch (-want +got): %s", diff)
			}
		})
	}
}
