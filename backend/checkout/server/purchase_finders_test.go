package server

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestFindLatestPurchase(t *testing.T) {
	type args struct {
		ctx context.Context
		uId uint64
	}
	tests := []struct {
		name       string
		args       args
		want       *purchase.Purchase
		queryStubs []*sqlassist.QueryStubber
		wantStatus xerrors.Code
	}{
		{
			name:       "not found",
			wantStatus: xerrors.NotFound,
			args: args{
				ctx: context.Background(),
				uId: gPurchases.Foo.UserId,
			},
			queryStubs: []*sqlassist.QueryStubber{
				{

					Args: []driver.Value{gPurchases.Foo.UserId},
					Expect: `SELECT * FROM purchases WHERE id=(
            SELECT order_id FROM purchase_state_changes WHERE order_id IN (
                SELECT id FROM purchases WHERE user_id = $1
            )
             ORDER BY stage ASC, created_at DESC
        )`,
					Err: sql.ErrNoRows,
				},
			},
		},

		{
			name: "basic",
			args: args{
				ctx: context.Background(),
				uId: gPurchases.Foo.UserId,
			},
			queryStubs: []*sqlassist.QueryStubber{
				{

					Args: []driver.Value{gPurchases.Foo.UserId},
					Expect: `SELECT * FROM purchases WHERE id=(
            SELECT order_id FROM purchase_state_changes WHERE order_id IN (
                SELECT id FROM purchases WHERE user_id = $1
            )
             ORDER BY stage ASC, created_at DESC
        )`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gPurchases.Foo)).
						AddRow(storeql.SQLValues(gPurchases.Foo)...),
				},
			},
			want: gPurchases.Foo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}

			got, err := FindLatestPurchase(tt.args.ctx, db, tt.args.uId)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("FindLatestPurchase() error = %v, wantErr %v", err, tt.wantStatus)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreUnexported(timestamppb.Timestamp{})); diff != "" {
				t.Errorf("FindLatestPurchase() mismatch (-want +got): %s", diff)
			}

		})
	}
}
func TestFindPurchase(t *testing.T) {
	type args struct {
		ctx context.Context
		uId uint64
		oId uint64
	}
	tests := []struct {
		name       string
		args       args
		want       *purchase.Purchase
		queryStubs []*sqlassist.QueryStubber
		wantStatus xerrors.Code
	}{
		{
			name: "basic",
			args: args{
				ctx: context.Background(),
				uId: gPurchases.Foo.UserId,
				oId: gPurchases.Foo.Id,
			},
			queryStubs: []*sqlassist.QueryStubber{
				{

					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.UserId},
					Expect: `SELECT * FROM purchases WHERE id=$1 AND user_id=$2`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gPurchases.Foo)).
						AddRow(storeql.SQLValues(gPurchases.Foo)...),
				},
			},
			want: gPurchases.Foo,
		},
		{
			name: "no rows",
			args: args{
				ctx: context.Background(),
				uId: gPurchases.Foo.UserId,
				oId: gPurchases.Foo.Id,
			},
			queryStubs: []*sqlassist.QueryStubber{
				{

					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.UserId},
					Expect: `SELECT * FROM purchases WHERE id=$1 AND user_id=$2`,
					Err:    sql.ErrNoRows,
				},
			},
			wantStatus: xerrors.NotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}

			got, err := FindPurchase(tt.args.ctx, db, tt.args.oId, tt.args.uId)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("FindPurchase() error = %v, wantErr %v", err, tt.wantStatus)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreUnexported(timestamppb.Timestamp{})); diff != "" {
				t.Errorf("FindPurchase() mismatch (-want +got): %s", diff)
			}

		})
	}
}
