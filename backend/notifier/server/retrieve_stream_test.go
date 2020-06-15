package server

import (
	"context"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/notifier/internal/notifiertest"
	"github.com/athomecomar/athome/pb/pbnotifier"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestServer_retrieveStream(t *testing.T) {
	var now = time.Now()
	type args struct {
		ctx    context.Context
		userId uint64
		offset time.Time
	}
	tests := []struct {
		name       string
		queryStubs []*sqlassist.QueryStubber
		args       args
		want       *pbnotifier.RetrieveManyResponse
		wantCode   xerrors.Code
	}{
		{
			name: "base",
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: `SELECT * FROM notifications WHERE user_id=$1 AND created_at>=$2 ORDER BY created_at`,
					Args:   []driver.Value{4, now},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gNotifications.Consumers.Foo)).
						AddRow(storeql.SQLValues(gNotifications.Consumers.Foo)...).
						AddRow(storeql.SQLValues(gNotifications.Consumers.Bar)...),
				},
			},
			wantCode: xerrors.OK,
			args: args{
				ctx:    context.Background(),
				userId: 4,
				offset: now,
			},
			want: &pbnotifier.RetrieveManyResponse{
				Notifications: map[uint64]*pbnotifier.Notification{
					gNotifications.Consumers.Foo.Id: notifiertest.NotificationToPb(t, gNotifications.Consumers.Foo),
					gNotifications.Consumers.Bar.Id: notifiertest.NotificationToPb(t, gNotifications.Consumers.Bar),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{}
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}

			got, err := s.retrieveStream(tt.args.ctx, db, tt.args.userId, tt.args.offset)
			if status.Code(err) != tt.wantCode {
				t.Errorf("Server.retrieveStream() error = %v, wantCode %v", err, tt.wantCode)
				return
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(timestamppb.Timestamp{})); diff != "" {
				t.Errorf("Server.retrieveStream() mismatch (-got +want): %s", diff)
				return
			}
		})
	}
}
