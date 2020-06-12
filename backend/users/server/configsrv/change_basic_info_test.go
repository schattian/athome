package configsrv

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/test/pbauthtest"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func TestServer_changeBasicInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbusers.ChangeBasicInfoRequest
	}
	type authStub struct {
		err error
		uid uint64
	}

	tests := []struct {
		name       string
		args       args
		queryStubs []*sqlassist.QueryStubber
		authStub   authStub
		execStub   *sqlassist.ExecStubber
		wantStatus xerrors.Code
	}{
		{
			name: "ok",
			args: args{
				ctx: context.Background(),
				in: &pbusers.ChangeBasicInfoRequest{
					Name: "buzz", Surname: "lightyear",
				},
			},
			authStub: authStub{
				uid: gUsers.Consumers.Foo.Id,
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM users",
					Rows: sqlmock.NewRows(
						storeql.SQLColumns(
							gUsers.Consumers.Foo,
						)).AddRow(
						storeql.SQLValues(gUsers.Consumers.Foo)...,
					),
				},
			},
			execStub: &sqlassist.ExecStubber{
				Expect: "UPDATE users SET", Result: sqlmock.NewResult(1, 1),
			},
			wantStatus: xerrors.OK,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			c := pbauthtest.NewCtrlFromRetrieve(ctrl, tt.args.ctx, tt.args.in.AccessToken, tt.authStub.uid)

			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}
			tt.execStub.Stub(mock)
			s := &Server{}
			_, err := s.changeBasicInfo(tt.args.ctx, db, c, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.changePassword() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
		})
	}
}
