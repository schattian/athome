package server

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

const fooPwd = "foopassword3"

func TestServer_signUpEnd(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbuser.SignUpEndRequest
	}
	tests := []struct {
		name       string
		args       args
		queryStubs []*sqlassist.QueryStubber
		execStubs  []*sqlassist.ExecStubber
		want       *pbuser.SignUpEndResponse
		wantStatus xerrors.Code
	}{
		{
			name: "basic consumer",
			args: args{
				ctx: context.Background(),
				in:  onboardingToSignUpEndRequest(gOnboardings.Consumers.Foo, fooPwd),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM onboardings",
					Rows:   sqlmock.NewRows(storeql.SQLColumns(gOnboardings.Consumers.Foo)).AddRow(storeql.SQLValues(setStage(gOnboardings.Consumers.Foo, field.Shared))...),
				},
				{
					Expect: "INSERT INTO users", Rows: sqlmock.NewRows([]string{"id"}).AddRow(gUsers.Consumers.Foo.Id),
				},
			},
			execStubs: []*sqlassist.ExecStubber{
				{
					Expect: "DELETE FROM onboardings", Result: sqlmock.NewResult(0, 1),
				},
			},
			want:       &pbuser.SignUpEndResponse{User: userToSignInUserUnsafe(t, gUsers.Consumers.Foo)},
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
			got, err := s.signUpEnd(tt.args.ctx, db, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.signUpEnd() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
			tt.want.User.SignToken = got.User.GetSignToken()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Server.signUpEnd() errored mismatch (-want +got): %s", diff)
			}

			if status.Code(err) != xerrors.OK {
				return
			}

			claims, err := claimJwt(got.User.GetSignToken(), userconf.GetSIGN_JWT_SECRET)
			if err != nil {
				t.Fatalf("Server.signUpEnd() errored parsing generated sign token: %v", err)
			}
			userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
			if err != nil {
				t.Fatalf("Server.signUpEnd() errored parsing generated sign token: %v", err)
			}
			if userId != tt.want.User.Id {
				t.Errorf("Server.signUpEnd() errored mismatch user id assigned to sign token. Got %v, want %v", userId, tt.want.User.Id)
			}
			if err := claims.Valid(); err != nil {
				t.Errorf("Server.signUpEnd() given invalid sign token: %v", err)
			}

		})
	}
}
