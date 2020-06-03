package signsrv

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func TestServer_resetPassword(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbusers.ResetPasswordRequest
	}
	tests := []struct {
		name       string
		args       args
		queryStubs []*sqlassist.QueryStubber
		execStubs  []*sqlassist.ExecStubber
		wantStatus xerrors.Code
	}{
		{
			name: "basic change of pwd",
			args: args{
				ctx: context.Background(),
				in:  &pbusers.ResetPasswordRequest{Token: gTokens.Forgot.Valid, Password: fooPwd},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM users",
					Rows:   sqlmock.NewRows(storeql.SQLColumns(gUsers.Consumers.Foo)).AddRow(storeql.SQLValues(gUsers.Consumers.Foo)...),
				},
			},
			execStubs: []*sqlassist.ExecStubber{
				{
					Expect: "UPDATE users SET", Result: sqlmock.NewResult(1, 1),
				},
			},
			wantStatus: xerrors.OK,
		},
		{
			name: "expired token",
			args: args{
				ctx: context.Background(),
				in:  &pbusers.ResetPasswordRequest{Token: gTokens.Forgot.Expired, Password: fooPwd},
			},
			wantStatus: xerrors.InvalidArgument,
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
			_, err := s.resetPassword(tt.args.ctx, db, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.resetPassword() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
			if status.Code(err) != xerrors.OK {
				return
			}
		})
	}
}
