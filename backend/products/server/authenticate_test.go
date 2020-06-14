package server

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/athome/backend/products/internal/productstest"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/test/pbauthtest"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/golang/mock/gomock"
)

func Test_retrieveLatestDraft(t *testing.T) {
	type authStub struct {
		req  *pbauth.RetrieveAuthenticationRequest
		resp *pbauth.RetrieveAuthenticationResponse
		err  error
	}
	type args struct {
		ctx         context.Context
		accessToken string
	}
	tests := []struct {
		name       string
		auth       authStub
		queryStubs []*sqlassist.QueryStubber
		args       args
		want       *ent.Draft
		wantErr    bool
	}{

		{
			name: "existing draft",
			args: args{
				ctx:         context.Background(),
				accessToken: "fooAccessToken",
			},
			auth: authStub{
				req:  &pbauth.RetrieveAuthenticationRequest{AccessToken: "fooAccessToken"},
				resp: &pbauth.RetrieveAuthenticationResponse{UserId: gDrafts.Foo.UserId},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM drafts WHERE user_id=$1",
					Args:   []driver.Value{gDrafts.Foo.UserId},
					Rows:   sqlmock.NewRows(storeql.SQLColumns(gDrafts.Foo)).AddRow(storeql.SQLValues(gDrafts.Foo)...),
				},
			},
			want: gDrafts.Foo,
		},

		{
			name: "no draft (creates new one)",
			args: args{
				ctx:         context.Background(),
				accessToken: "fooAccessToken",
			},
			auth: authStub{
				req:  &pbauth.RetrieveAuthenticationRequest{AccessToken: "fooAccessToken"},
				resp: &pbauth.RetrieveAuthenticationResponse{UserId: gDrafts.Foo.UserId},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM drafts WHERE user_id=$1",
					Args:   []driver.Value{gDrafts.Foo.UserId},
					Err:    sql.ErrNoRows,
				},
				{
					Expect: "INSERT INTO drafts", Rows: sqlmock.NewRows([]string{"id"}).AddRow(gDrafts.Foo.Id),
					Args: []driver.Value{stage.First, gDrafts.Foo.UserId},
				},
			},
			want: productstest.SetStage(t, gDrafts.Foo, stage.First),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			auth := pbauthtest.NewMockAuthClient(ctrl)
			auth.EXPECT().RetrieveAuthentication(tt.args.ctx, tt.auth.req).Return(tt.auth.resp, tt.auth.err)

			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}

			got, err := retrieveLatestDraft(tt.args.ctx, db, auth, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("retrieveLatestDraft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("retrieveLatestDraft() = %v, want %v", got, tt.want)
			}
		})
	}
}
