package configsrv

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/users/internal/pbauthtest"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

const fooOldPwd = "foopassword3"
const fooNewPwd = "barpassword4"

func TestServer_changePassword(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbusers.ChangePasswordRequest
	}
	tests := []struct {
		name       string
		args       args
		c          pbauth.AuthClient
		oldPwd     string
		queryStubs []*sqlassist.QueryStubber
		execStub   *sqlassist.ExecStubber
		wantStatus xerrors.Code
	}{
		{
			name:   "ok",
			oldPwd: fooOldPwd,
			args: args{
				ctx: context.Background(),
				in:  &pbusers.ChangePasswordRequest{OldPassword: fooOldPwd, NewPassword: fooNewPwd},
			},
			c: pbauthtest.Client{Retrieve: &pbauth.RetrieveAuthenticationResponse{UserId: gUsers.Consumers.Foo.Id}},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM users",
					Rows: sqlmock.NewRows(
						storeql.SQLColumns(
							gUsers.Consumers.Foo,
						)).AddRow(
						storeql.SQLValues(assignPassword(t, gUsers.Consumers.Foo, fooOldPwd))...,
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
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}
			tt.execStub.Stub(mock)
			s := &Server{}
			_, err := s.changePassword(tt.args.ctx, db, tt.c, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.changePassword() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
		})
	}
}
