package server

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func TestServer_signUpShared(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbuser.SignUpSharedRequest
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
				ctx: context.Background(),
				in:  onboardingToSignUpSharedRequest(gOnboardings.Consumers.Foo),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM onboardings",
					Rows:   sqlmock.NewRows(storeql.SQLColumns(gOnboardings.Consumers.Foo)).AddRow(storeql.SQLValues(setStage(gOnboardings.Consumers.Foo, field.Start))...),
				},
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
				ctx: context.Background(),
				in:  onboardingToSignUpSharedRequest(gOnboardings.Consumers.Foo),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM onboardings",
					Rows:   sqlmock.NewRows(storeql.SQLColumns(gOnboardings.Consumers.Foo)).AddRow(storeql.SQLValues(setStage(gOnboardings.Consumers.Foo, field.Start))...),
				},
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
				ctx: context.Background(),
				in:  onboardingToSignUpSharedRequest(gOnboardings.Consumers.Foo),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM onboardings",
					Rows:   sqlmock.NewRows(storeql.SQLColumns(gOnboardings.Consumers.Foo)).AddRow(storeql.SQLValues(setStage(gOnboardings.Consumers.Foo, field.Shared))...),
				},
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
		{
			name: "unexistant onboarding",
			args: args{
				ctx: context.Background(),
				in:  onboardingToSignUpSharedRequest(gOnboardings.Consumers.Foo),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM onboardings",
					Err:    sql.ErrNoRows,
				},
			},
			wantStatus: xerrors.NotFound,
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
			_, err := s.signUpShared(tt.args.ctx, db, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.signUpShared() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
		})
	}
}
