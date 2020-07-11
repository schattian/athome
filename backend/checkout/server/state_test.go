package server

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/internal/checkouttest"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func TestMustPrevState(t *testing.T) {
	type args struct {
		ctx      context.Context
		stateful sm.Stateful
		desired  sm.StateName
		uid      uint64
	}
	tests := []struct {
		name       string
		args       args
		queryStubs []*sqlassist.QueryStubber
		wantStatus xerrors.Code
	}{
		{
			name: "basic",
			args: args{
				ctx:      context.Background(),
				stateful: gPurchases.Foo,
				desired:  sm.PurchasePaid,
				uid:      gPurchases.Foo.UserId,
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Purchases.ShippingMethodSelected)).
						AddRow(storeql.SQLValues(gStateChanges.Purchases.ShippingMethodSelected)...),
				},
			},
		},
		{
			name: "not allowed",
			args: args{
				ctx:      context.Background(),
				stateful: gPurchases.Foo,
				desired:  sm.PurchaseConfirmed,
				uid:      gPurchases.Foo.UserId,
			},
			wantStatus: xerrors.OutOfRange,
			queryStubs: []*sqlassist.QueryStubber{
				{
					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Purchases.Paid)).
						AddRow(storeql.SQLValues(gStateChanges.Purchases.Paid)...),
				},
			},
		},
		{
			name: "invalid prev state",
			args: args{
				ctx:      context.Background(),
				stateful: gPurchases.Foo,
				desired:  sm.PurchasePaid,
				uid:      gPurchases.Foo.UserId,
			},
			wantStatus: xerrors.InvalidArgument,
			queryStubs: []*sqlassist.QueryStubber{
				{
					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Purchases.Addressed)).
						AddRow(storeql.SQLValues(gStateChanges.Purchases.Addressed)...),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}

			err := MustPrevState(tt.args.ctx, db, tt.args.stateful, tt.args.desired, tt.args.uid)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("MustPrevState() error = %v, wantErr %v", err, tt.wantStatus)
			}
		})
	}
}

func TestChangeState(t *testing.T) {
	type args struct {
		ctx          context.Context
		stateChanger sm.StateChanger
		s            sm.Stateful
		uid          uint64
	}
	tests := []struct {
		name       string
		args       args
		queryStubs []*sqlassist.QueryStubber
		wantStatus xerrors.Code
	}{

		{
			name: "order without shipping: next: confirmed",
			args: args{
				ctx:          context.Background(),
				stateChanger: sm.Next,
				s:            checkouttest.PurchaseZeroShippingId(t, gPurchases.Foo),
				uid:          gPurchases.Foo.UserId,
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Purchases.Confirmed)).
						AddRow(storeql.SQLValues(gStateChanges.Purchases.Confirmed)...),
				},
				{
					Expect: storeql.ExecBoilerplate("INSERT", gStateChanges.Purchases.ShippingMethodSelected),
					Rows: sqlmock.NewRows([]string{"id"}).
						AddRow(gStateChanges.Purchases.Finished.Id),
				},
			},
		},
		{
			name: "order with shipping: next: confirmed",
			args: args{
				ctx:          context.Background(),
				stateChanger: sm.Next,
				s:            gPurchases.Foo,
				uid:          gPurchases.Foo.UserId,
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Purchases.Confirmed)).
						AddRow(storeql.SQLValues(gStateChanges.Purchases.Confirmed)...),
				},
				{
					Args:   []driver.Value{gShippings.Foo.Id},
					Expect: `SELECT * FROM shippings WHERE id=$1`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gShippings.Foo)).
						AddRow(storeql.SQLValues(gShippings.Foo)...),
				},
				{
					Args:   []driver.Value{gShippings.Foo.Id, gShippings.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Shippings.Created)).
						AddRow(storeql.SQLValues(gStateChanges.Shippings.Finished)...),
				},
				{
					Expect: storeql.ExecBoilerplate("INSERT", gStateChanges.Purchases.ShippingMethodSelected),
					Rows: sqlmock.NewRows([]string{"id"}).
						AddRow(gStateChanges.Purchases.Finished.Id),
				},
			},
		},
		{
			name: "order with unfinished shipping: next: confirmed - invalid",
			args: args{
				ctx:          context.Background(),
				stateChanger: sm.Next,
				s:            gPurchases.Foo,
				uid:          gPurchases.Foo.UserId,
			},
			wantStatus: xerrors.InvalidArgument,
			queryStubs: []*sqlassist.QueryStubber{
				{
					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Purchases.Confirmed)).
						AddRow(storeql.SQLValues(gStateChanges.Purchases.Confirmed)...),
				},
				{
					Args:   []driver.Value{gShippings.Foo.Id},
					Expect: `SELECT * FROM shippings WHERE id=$1`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gShippings.Foo)).
						AddRow(storeql.SQLValues(gShippings.Foo)...),
				},
				{
					Args:   []driver.Value{gShippings.Foo.Id, gShippings.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Shippings.Created)).
						AddRow(storeql.SQLValues(gStateChanges.Shippings.Created)...),
				},
				{
					Expect: storeql.ExecBoilerplate("INSERT", gStateChanges.Purchases.ShippingMethodSelected),
					Rows: sqlmock.NewRows([]string{"id"}).
						AddRow(gStateChanges.Purchases.Finished.Id),
				},
			},
		},

		{
			name: "next: paid",
			args: args{
				ctx:          context.Background(),
				stateChanger: sm.Next,
				s:            gPurchases.Foo,
				uid:          gPurchases.Foo.MerchantId,
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Purchases.Paid)).
						AddRow(storeql.SQLValues(gStateChanges.Purchases.Paid)...),
				},
				{
					Expect: storeql.ExecBoilerplate("INSERT", gStateChanges.Purchases.ShippingMethodSelected),
					Rows: sqlmock.NewRows([]string{"id"}).
						AddRow(gStateChanges.Purchases.Confirmed.Id),
				},
			},
		},

		{
			name: "next: not allowed",
			args: args{
				ctx:          context.Background(),
				stateChanger: sm.Next,
				s:            gPurchases.Foo,
				uid:          gPurchases.Foo.UserId,
			},
			wantStatus: xerrors.PermissionDenied,
			queryStubs: []*sqlassist.QueryStubber{
				{
					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Purchases.Paid)).
						AddRow(storeql.SQLValues(gStateChanges.Purchases.Paid)...),
				},
			},
		},
		{
			name: "next: addressed",
			args: args{
				ctx:          context.Background(),
				stateChanger: sm.Next,
				s:            gPurchases.Foo,
				uid:          gPurchases.Foo.UserId,
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Purchases.Addressed)).
						AddRow(storeql.SQLValues(gStateChanges.Purchases.Addressed)...),
				},
				{
					Expect: storeql.ExecBoilerplate("INSERT", gStateChanges.Purchases.ShippingMethodSelected),
					Rows: sqlmock.NewRows([]string{"id"}).
						AddRow(gStateChanges.Purchases.Addressed.Id),
				},
			},
		},
		{
			name: "next: invalid",
			args: args{
				ctx:          context.Background(),
				stateChanger: sm.Next,
				s:            gPurchases.Foo,
				uid:          gPurchases.Foo.UserId,
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.SQLTable()},
					Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
					Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Purchases.Addressed)).
						AddRow(storeql.SQLValues(gStateChanges.Purchases.Addressed)...),
				},
				{
					Expect: storeql.ExecBoilerplate("INSERT", gStateChanges.Purchases.Addressed),
					Rows: sqlmock.NewRows([]string{"id"}).
						AddRow(gStateChanges.Purchases.Addressed.Id),
				},
			},
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}

			err := ChangeState(tt.args.ctx, db, tt.args.stateChanger, tt.args.s, tt.args.uid)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("ChangeState() error = %v, wantStatus %v", err, tt.wantStatus)
			}
		})
	}
}
