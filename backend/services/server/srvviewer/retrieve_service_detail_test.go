package srvviewer

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/test/pbaddresstest"
	"github.com/athomecomar/athome/pb/test/pbuserstest"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/status"
)

func TestServer_retrieveServiceDetail(t *testing.T) {
	type usersStub struct {
		req  *pbusers.RetrieveUserRequest
		resp *pbusers.UserDetail
		err  error
	}
	type addrStub struct {
		req  *pbaddress.RetrieveAddressRequest
		resp *pbaddress.Address
		err  error
	}
	type args struct {
		ctx context.Context
		in  *pbservices.RetrieveServiceDetailRequest
	}
	tests := []struct {
		name       string
		addr       addrStub
		users      usersStub
		queryStubs []*sqlassist.QueryStubber
		args       args
		want       *pbservices.ServiceDetail
		wantCode   xerrors.Code
	}{

		{
			name: "not found",
			args: args{
				ctx: context.Background(),
				in:  &pbservices.RetrieveServiceDetailRequest{ServiceId: gServices.Foo.Id},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM services WHERE id=$1",
					Args:   []driver.Value{gServices.Foo.Id},
					Err:    sql.ErrNoRows,
				},
			},
			wantCode: xerrors.NotFound,
		},

		{
			name: "basic retrieve",
			args: args{
				ctx: context.Background(),
				in:  &pbservices.RetrieveServiceDetailRequest{ServiceId: gServices.Foo.Id},
			},
			addr: addrStub{
				req:  &pbaddress.RetrieveAddressRequest{AddressId: gServices.Foo.AddressId},
				resp: &pbaddress.Address{Alias: "foo", Department: 8},
			},
			users: usersStub{
				req:  &pbusers.RetrieveUserRequest{UserId: gServices.Foo.UserId},
				resp: &pbusers.UserDetail{User: &pbusers.User{Name: "foo"}},
			},

			want: &pbservices.ServiceDetail{
				User:    &pbservices.User{Name: "foo"},
				Address: &pbservices.Address{Department: 8},
				Service: gServices.Foo.ToPb(),
				Calendar: &pbservices.CalendarDetail{
					Calendar: gCalendars.Foo.Medic.A.ToPb(),
					Availabilities: map[uint64]*pbservices.Availability{
						gAvailabilities.Foo.Medic.First.A.Id: gAvailabilities.Foo.Medic.First.A.ToPb(),
						gAvailabilities.Foo.Medic.First.B.Id: gAvailabilities.Foo.Medic.First.B.ToPb(),
						gAvailabilities.Foo.Medic.First.C.Id: gAvailabilities.Foo.Medic.First.C.ToPb(),
					},
					Events: map[uint64]*pbservices.Event{
						gEvents.Foo.Medic.First.A.Id: gEvents.Foo.Medic.First.A.ToPb(),
						gEvents.Foo.Medic.First.B.Id: gEvents.Foo.Medic.First.B.ToPb(),
						gEvents.Foo.Medic.First.C.Id: gEvents.Foo.Medic.First.C.ToPb(),
					},
				},
			},

			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM services WHERE id=$1",
					Args:   []driver.Value{gServices.Foo.Id},
					Rows:   sqlmock.NewRows(storeql.SQLColumns(gServices.Foo)).AddRow(storeql.SQLValues(gServices.Foo)...),
				},
				{
					Expect: "SELECT * FROM calendars WHERE id=$1",
					Args:   []driver.Value{gServices.Foo.CalendarId},
					Rows:   sqlmock.NewRows(storeql.SQLColumns(gCalendars.Foo.Medic.A)).AddRow(storeql.SQLValues(gCalendars.Foo.Medic.A)...),
				},
				{
					Expect: "SELECT * FROM availabilities WHERE calendar_id=$1",
					Args:   []driver.Value{gServices.Foo.CalendarId},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gAvailabilities.Foo.Medic.First.A)).
						AddRow(storeql.SQLValues(gAvailabilities.Foo.Medic.First.A)...).
						AddRow(storeql.SQLValues(gAvailabilities.Foo.Medic.First.B)...).
						AddRow(storeql.SQLValues(gAvailabilities.Foo.Medic.First.C)...),
				},
				{
					Expect: "SELECT * FROM events WHERE calendar_id=$1",
					Args:   []driver.Value{gServices.Foo.CalendarId},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gEvents.Foo.Medic.First.A)).
						AddRow(storeql.SQLValues(gEvents.Foo.Medic.First.A)...).
						AddRow(storeql.SQLValues(gEvents.Foo.Medic.First.B)...).
						AddRow(storeql.SQLValues(gEvents.Foo.Medic.First.C)...),
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &Server{}
			ctrl := gomock.NewController(t)
			users := pbuserstest.NewMockViewerClient(ctrl)
			users.EXPECT().RetrieveUser(tt.args.ctx, tt.users.req).Return(tt.users.resp, tt.users.err)
			addr := pbaddresstest.NewMockAddressesClient(ctrl)
			addr.EXPECT().RetrieveAddress(tt.args.ctx, tt.addr.req).Return(tt.addr.resp, tt.addr.err)
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}
			got, err := s.retrieveServiceDetail(tt.args.ctx, db, addr, users, tt.args.in)
			if gotCode := status.Code(err); gotCode != tt.wantCode {
				t.Errorf("Server.retrieveServiceDetail() error = %v, wantErr %v", err, tt.wantCode)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Server.retrieveServiceDetail()  errored mismatch (-want +got): %s", diff)
			}
		})
	}
}
