package signsrv

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/usertest"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func TestServer_signUpSelectCategory(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbuser.SignUpSelectCategoryRequest
	}
	tests := []struct {
		name       string
		args       args
		queryStubs []*sqlassist.QueryStubber
		execStubs  []*sqlassist.ExecStubber
		wantStatus xerrors.Code
	}{
		{
			name: "oor service-provider",
			args: args{
				ctx: context.Background(),
				in:  &pbuser.SignUpSelectCategoryRequest{CategoryName: semprov.Medic.Name, OnboardingId: gOnboardings.ServiceProviders.Medic.Foo.Id},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM onboardings",
					Rows:   sqlmock.NewRows(storeql.SQLColumns(gOnboardings.Consumers.Foo)).AddRow(storeql.SQLValues(usertest.SetStage(gOnboardings.ServiceProviders.Medic.Foo, field.SelectCategory))...),
				},
			},
			wantStatus: xerrors.OutOfRange,
		},
		{
			name: "basic service-provider",
			args: args{
				ctx: context.Background(),
				in:  &pbuser.SignUpSelectCategoryRequest{CategoryName: semprov.Medic.Name, OnboardingId: gOnboardings.ServiceProviders.Medic.Foo.Id},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM onboardings",
					Rows:   sqlmock.NewRows(storeql.SQLColumns(gOnboardings.Consumers.Foo)).AddRow(storeql.SQLValues(usertest.SetStage(gOnboardings.ServiceProviders.Medic.Foo, field.Shared))...),
				},
			},
			execStubs: []*sqlassist.ExecStubber{
				{Expect: "UPDATE onboardings SET", Result: sqlmock.NewResult(1, 1)},
			},
			wantStatus: xerrors.OK,
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
			for _, stub := range tt.execStubs {
				stub.Stub(mock)
			}
			s := &Server{}
			_, err := s.signUpSelectCategory(tt.args.ctx, db, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.signUpSelectCategory() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
		})
	}
}
