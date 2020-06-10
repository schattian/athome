package signsrv

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/usertest"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func TestServer_signUpShared(t *testing.T) {
	type args struct {
		ctx      context.Context
		in       *pbusers.SignUpSharedRequest
		previous *ent.Onboarding
	}
	tests := []struct {
		name       string
		args       args
		queryStubs []*sqlassist.QueryStubber
		execStub   *sqlassist.ExecStubber
		wantStatus xerrors.Code
	}{
		{
			name: "basic consumer",
			args: args{
				ctx:      context.Background(),
				in:       usertest.OnboardingToSignUpSharedRequest(gOnboardings.Consumers.Foo),
				previous: usertest.SetStage(t, gOnboardings.Consumers.Foo, field.Start),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT COUNT(id) FROM users",
					Rows:   sqlmock.NewRows([]string{"COUNT(id)"}).AddRow(0),
				},
			},
			execStub: &sqlassist.ExecStubber{
				Expect: "UPDATE onboardings SET", Result: sqlmock.NewResult(1, 1),
			},
			wantStatus: xerrors.OK,
		},
		{
			name: "already exists basic consumer",
			args: args{
				ctx:      context.Background(),
				in:       usertest.OnboardingToSignUpSharedRequest(gOnboardings.Consumers.Foo),
				previous: usertest.SetStage(t, gOnboardings.Consumers.Foo, field.Start),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT COUNT(id) FROM users",
					Rows:   sqlmock.NewRows([]string{"COUNT(id)"}).AddRow(1),
				},
			},
			execStub: &sqlassist.ExecStubber{
				Expect: "UPDATE onboardings SET", Result: sqlmock.NewResult(1, 1),
			},
			wantStatus: xerrors.AlreadyExists,
		},
		{
			name: "invalid stage basic consumer",
			args: args{
				ctx:      context.Background(),
				in:       usertest.OnboardingToSignUpSharedRequest(gOnboardings.Consumers.Foo),
				previous: usertest.SetStage(t, gOnboardings.Consumers.Foo, field.Shared),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT COUNT(id) FROM users",
					Rows:   sqlmock.NewRows([]string{"COUNT(id)"}).AddRow(0),
				},
			},
			execStub: &sqlassist.ExecStubber{
				Expect: "UPDATE onboardings SET", Result: sqlmock.NewResult(1, 1),
			},
			wantStatus: xerrors.OutOfRange,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}
			if tt.execStub != nil {
				tt.execStub.Stub(mock)
			}
			s := &Server{}
			_, err := s.signUpShared(tt.args.ctx, db, tt.args.in, tt.args.previous)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.signUpShared() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
		})
	}
}
