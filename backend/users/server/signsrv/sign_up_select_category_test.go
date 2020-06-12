package signsrv

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/usertest"
	"github.com/athomecomar/athome/backend/users/internal/xpbsemantictest"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func TestServer_signUpSelectCategory(t *testing.T) {
	type args struct {
		ctx      context.Context
		in       *pbusers.SignUpSelectCategoryRequest
		previous *ent.Onboarding
	}
	type semanticStub struct {
		err  error
		resp *pbsemantic.Category
	}
	tests := []struct {
		name         string
		args         args
		queryStubs   []*sqlassist.QueryStubber
		execStubs    []*sqlassist.ExecStubber
		semanticStub semanticStub
		wantStatus   xerrors.Code
	}{
		{
			name: "oor service-provider",
			args: args{
				ctx:      context.Background(),
				in:       &pbusers.SignUpSelectCategoryRequest{CategoryId: 3, OnboardingId: gOnboardings.ServiceProviders.Medic.Foo.Id},
				previous: usertest.SetStage(t, gOnboardings.ServiceProviders.Medic.Foo, field.SelectCategory),
			},
			semanticStub: semanticStub{err: nil, resp: nil},
			wantStatus:   xerrors.OutOfRange,
		},
		{
			name: "basic service-provider",
			args: args{
				ctx:      context.Background(),
				in:       &pbusers.SignUpSelectCategoryRequest{CategoryId: 3, OnboardingId: gOnboardings.ServiceProviders.Medic.Foo.Id},
				previous: usertest.SetStage(t, gOnboardings.ServiceProviders.Medic.Foo, field.Shared),
			},
			semanticStub: semanticStub{err: nil, resp: nil},
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
			ctrl := gomock.NewController(t)
			c := xpbsemantictest.NewMockCategoriesClient(ctrl)
			c.EXPECT().
				RetrieveCategory(tt.args.ctx, &pbsemantic.RetrieveCategoryRequest{CategoryId: tt.args.in.CategoryId}).
				Return(tt.semanticStub.resp, tt.semanticStub.err)
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}
			for _, stub := range tt.execStubs {
				stub.Stub(mock)
			}
			s := &Server{}
			_, err := s.signUpSelectCategory(tt.args.ctx, db, c, tt.args.in, tt.args.previous)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.signUpSelectCategory() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
		})
	}
}
