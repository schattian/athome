package srvviewer

import (
	"context"
	"database/sql/driver"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/test/pbuserstest"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/status"
)

func TestServer_searchServices(t *testing.T) {
	type usersStub struct {
		req  *pbusers.RetrieveUserRequest
		resp *pbusers.UserDetail
		err  error
	}

	type args struct {
		ctx context.Context
		sem pbsemantic.ServiceProvidersClient
		img pbimages.ImagesClient
		in  *pbservices.SearchServicesRequest
	}
	tests := []struct {
		name       string
		args       args
		queryStubs []*sqlassist.QueryStubber
		users      []usersStub
		want       *pbservices.SearchServicesResponse
		wantCode   xerrors.Code
	}{
		{
			name: "basic paginated search",
			users: []usersStub{
				{
					req:  &pbusers.RetrieveUserRequest{UserId: gServices.Foo.UserId},
					resp: &pbusers.UserDetail{User: &pbusers.User{Name: "fooUserName"}},
				},
				{
					req:  &pbusers.RetrieveUserRequest{UserId: gServices.Bar.UserId},
					resp: &pbusers.UserDetail{User: &pbusers.User{Name: "barUserName"}},
				},
			},
			args: args{
				ctx: context.Background(),
				in: &pbservices.SearchServicesRequest{
					Query: "FÓóBáR",
					Page: &pbservices.PageRequest{
						Cursor: b64EncodeId(gServices.Foo.Id + 1),
						Size:   2,
					},
				},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: `SELECT (*) FROM services
    WHERE lower(unaccent(title)) ILIKE ESCAPE $1
    AND id < ` + strconv.Itoa(int(gServices.Foo.Id)+1) +
						` ORDER BY id DESC LIMIT $2`,
					Args: []driver.Value{"foobar", 2},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gServices.Foo)).
						AddRow(storeql.SQLValues(gServices.Foo)...).
						AddRow(storeql.SQLValues(gServices.Bar)...),
				},
				{
					Expect: "SELECT COUNT(*) FROM services WHERE lower(unaccent(title)) ILIKE ESCAPE $1",
					Args:   []driver.Value{"foobar"},
					Rows:   sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(4),
				},
			},
			want: &pbservices.SearchServicesResponse{
				Services: map[uint64]*pbservices.ServiceSearchResult{
					gServices.Foo.Id: {
						Service: &pbservices.ServiceSearchResult_Service{Title: gServices.Foo.Title, Price: gServices.Foo.PbPrice()},
						User:    &pbservices.User{Name: "fooUserName"},
					},
					gServices.Bar.Id: {
						Service: &pbservices.ServiceSearchResult_Service{Title: gServices.Bar.Title, Price: gServices.Bar.PbPrice()},
						User:    &pbservices.User{Name: "barUserName"},
					},
				},
				Page: &pbservices.PageResponse{
					NextCursor: b64EncodeId(gServices.Bar.Id),
					TotalSize:  4,
				},
			},
		},

		{
			name: "nil cursor given (first iteration)",
			users: []usersStub{
				{
					req:  &pbusers.RetrieveUserRequest{UserId: gServices.Foo.UserId},
					resp: &pbusers.UserDetail{User: &pbusers.User{Name: "fooUserName"}},
				},
				{
					req:  &pbusers.RetrieveUserRequest{UserId: gServices.Bar.UserId},
					resp: &pbusers.UserDetail{User: &pbusers.User{Name: "barUserName"}},
				},
			},
			args: args{
				ctx: context.Background(),
				in: &pbservices.SearchServicesRequest{
					Query: "FÓóBáR",
					Page: &pbservices.PageRequest{
						Size: 2,
					},
				},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: `SELECT (*) FROM services
    WHERE lower(unaccent(title)) ILIKE ESCAPE $1
    ORDER BY id DESC LIMIT $2`,
					Args: []driver.Value{"foobar", 2},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gServices.Foo)).
						AddRow(storeql.SQLValues(gServices.Foo)...).
						AddRow(storeql.SQLValues(gServices.Bar)...),
				},
				{
					Expect: "SELECT COUNT(*) FROM services WHERE lower(unaccent(title)) ILIKE ESCAPE $1",
					Args:   []driver.Value{"foobar"},
					Rows:   sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(4),
				},
			},
			want: &pbservices.SearchServicesResponse{
				Services: map[uint64]*pbservices.ServiceSearchResult{
					gServices.Foo.Id: {
						Service: &pbservices.ServiceSearchResult_Service{Title: gServices.Foo.Title, Price: gServices.Foo.PbPrice()},
						User:    &pbservices.User{Name: "fooUserName"},
					},
					gServices.Bar.Id: {
						Service: &pbservices.ServiceSearchResult_Service{Title: gServices.Bar.Title, Price: gServices.Bar.PbPrice()},
						User:    &pbservices.User{Name: "barUserName"},
					},
				},
				Page: &pbservices.PageResponse{
					NextCursor: b64EncodeId(gServices.Bar.Id),
					TotalSize:  4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			users := pbuserstest.NewMockViewerClient(ctrl)
			for _, stub := range tt.users {
				users.EXPECT().RetrieveUser(tt.args.ctx, stub.req).Return(stub.resp, stub.err).AnyTimes()
			}
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}
			s := &Server{}
			got, err := s.searchServices(tt.args.ctx, db, tt.args.sem, users, tt.args.img, tt.args.in)
			if gotCode := status.Code(err); gotCode != tt.wantCode {
				t.Errorf("Server.searchServices() error = %v, wantErr %v", err, tt.wantCode)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Server.searchServices()  errored mismatch (-want +got): %s", diff)
			}
		})
	}
}
