package signsrv

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/pbsemantictest"
	"github.com/athomecomar/athome/backend/users/internal/usertest"
	"github.com/athomecomar/athome/backend/users/internal/xpbsemantic"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func TestServer_signUpSelectCategory(t *testing.T) {
	type args struct {
		ctx      context.Context
		in       *pbusers.SignUpSelectCategoryRequest
		sem      xpbsemantic.CategoriesClient
		previous *ent.Onboarding
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
				ctx:      context.Background(),
				in:       &pbusers.SignUpSelectCategoryRequest{CategoryId: 3, OnboardingId: gOnboardings.ServiceProviders.Medic.Foo.Id},
				previous: usertest.SetStage(t, gOnboardings.ServiceProviders.Medic.Foo, field.SelectCategory),
				sem:      pbsemantictest.Client{},
			},
			wantStatus: xerrors.OutOfRange,
		},
		{
			name: "basic service-provider",
			args: args{
				ctx:      context.Background(),
				in:       &pbusers.SignUpSelectCategoryRequest{CategoryId: 3, OnboardingId: gOnboardings.ServiceProviders.Medic.Foo.Id},
				previous: usertest.SetStage(t, gOnboardings.ServiceProviders.Medic.Foo, field.Shared),
				sem:      pbsemantictest.Client{},
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
			_, err := s.signUpSelectCategory(tt.args.ctx, db, tt.args.sem, tt.args.in, tt.args.previous)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.signUpSelectCategory() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
		})
	}
}
