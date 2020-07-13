package srvpurchases

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/checkout/internal/checkouttest"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/athome/pb/test/pbproductstest"
	"github.com/athomecomar/athome/pb/test/pbservicestest"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type calendarStub struct {
	req  *pbservices.ConfirmEventRequest
	resp *emptypb.Empty
	err  error
}

func TestServer_confirmPurchase(t *testing.T) {
	type stubs struct {
		query []*sqlassist.QueryStubber
		cals  *calendarStub
		prodV *prodViewerStub
		prodM *prodManagerReserveStub
	}

	type args struct {
		ctx    context.Context
		in     *pbcheckout.ConfirmPurchaseRequest
		userId uint64
	}
	tests := []struct {
		name  string
		stubs stubs
		args  args

		want       *pbcheckout.Purchase
		wantStatus xerrors.Code
	}{
		{
			name: "unauthorized: being provider",
			args: args{
				ctx:    context.Background(),
				in:     &pbcheckout.ConfirmPurchaseRequest{PurchaseId: gPurchases.Foo.Id},
				userId: gPurchases.Foo.UserId,
			},
			wantStatus: xerrors.PermissionDenied,
			stubs: stubs{
				query: []*sqlassist.QueryStubber{
					{
						Expect: "SELECT * FROM purchases WHERE id=$1",
						Args:   []driver.Value{gPurchases.Foo.Id},
						Rows: sqlmock.NewRows(storeql.SQLColumns(gPurchases.Foo)).
							AddRow(storeql.SQLValues(gPurchases.Foo)...),
					},
				},
			},
		},
		{
			name: "unauthorized: being consumer",
			args: args{
				ctx:    context.Background(),
				in:     &pbcheckout.ConfirmPurchaseRequest{PurchaseId: gPurchases.Foo.Id},
				userId: gPurchases.Foo.UserId,
			},
			wantStatus: xerrors.PermissionDenied,
			stubs: stubs{
				query: []*sqlassist.QueryStubber{
					{
						Expect: "SELECT * FROM purchases WHERE id=$1",
						Args:   []driver.Value{gPurchases.Foo.Id},
						Rows: sqlmock.NewRows(storeql.SQLColumns(gPurchases.Foo)).
							AddRow(storeql.SQLValues(gPurchases.Foo)...),
					},
				},
			},
		},
		{
			name: "basic",
			args: args{
				ctx:    context.Background(),
				in:     &pbcheckout.ConfirmPurchaseRequest{PurchaseId: gPurchases.Foo.Id},
				userId: gPurchases.Foo.MerchantId,
			},
			stubs: stubs{
				query: []*sqlassist.QueryStubber{
					{
						Expect: "SELECT * FROM purchases WHERE id=$1",
						Args:   []driver.Value{gPurchases.Foo.Id},
						Rows: sqlmock.NewRows(storeql.SQLColumns(gPurchases.Foo)).
							AddRow(storeql.SQLValues(gPurchases.Foo)...),
					},
					{
						Expect: "SELECT * FROM shippings WHERE id=$1",
						Args:   []driver.Value{gPurchases.Foo.ShippingId},
						Rows: sqlmock.NewRows(storeql.SQLColumns(gShippings.Foo)).
							AddRow(storeql.SQLValues(gShippings.Foo)...),
					},
					{
						Args:   []driver.Value{gPurchases.Foo.Id, gPurchases.Foo.SQLTable()},
						Expect: `SELECT * FROM state_changes WHERE entity_id=$1 AND entity_table=$2 ORDER BY created_at`,
						Rows: sqlmock.NewRows(storeql.SQLColumns(gStateChanges.Purchases.Paid)).
							AddRow(storeql.SQLValues(gStateChanges.Purchases.Paid)...),
					},
					{
						Expect: storeql.ExecBoilerplate("INSERT", gStateChanges.Purchases.Paid),
						Rows: sqlmock.NewRows([]string{"id"}).
							AddRow(gStateChanges.Purchases.Paid.Id),
					},
					{
						Expect: "SELECT * FROM shippings WHERE id=$1",
						Args:   []driver.Value{gPurchases.Foo.ShippingId},
						Rows: sqlmock.NewRows(storeql.SQLColumns(gShippings.Foo)).
							AddRow(storeql.SQLValues(gShippings.Foo)...),
					},
				},
				cals: &calendarStub{
					req:  &pbservices.ConfirmEventRequest{EventId: gShippings.Foo.EventId},
					resp: &emptypb.Empty{},
				},
				prodV: &prodViewerStub{
					req: &pbproducts.RetrieveProductsRequest{Ids: gPurchases.Foo.ProductIds()},
					resp: &pbproducts.RetrieveProductsResponse{
						Products: map[uint64]*pbproducts.Product{
							fooAProductId: gPbProducts.Foo.A,
							fooBProductId: gPbProducts.Foo.B,
						},
					},
				},
				prodM: &prodManagerReserveStub{
					req:  &pbproducts.ReserveStockRequest{Order: pbutil.ToPbEntity(gPurchases.Foo)},
					resp: &emptypb.Empty{},
				},
			},
			want: checkouttest.PurchaseToPb(t,
				gPurchases.Foo,
				gShippings.Foo.OrderPrice.Float64()+
					gPbProducts.Foo.A.Price*float64(gPurchases.Foo.Items[fooAProductId])+
					gPbProducts.Foo.B.Price*float64(gPurchases.Foo.Items[fooBProductId]),
			),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.stubs.query {
				stub.Stub(mock)
			}

			ctrl := gomock.NewController(t)
			var prodsV pbproducts.ViewerClient
			if stub := tt.stubs.prodV; stub != nil {
				prodsMock := pbproductstest.NewMockViewerClient(ctrl)
				prodsMock.EXPECT().RetrieveProducts(tt.args.ctx, stub.req).Return(stub.resp, stub.err)
				prodsV = prodsMock
			}
			var prodsM pbproducts.ManagerClient
			if stub := tt.stubs.prodM; stub != nil {
				prodMMock := pbproductstest.NewMockManagerClient(ctrl)
				prodMMock.EXPECT().ConfirmReserveStock(tt.args.ctx, stub.req).Return(stub.resp, stub.err)
				prodsM = prodMMock
			}
			var cals pbservices.CalendarsClient
			if stub := tt.stubs.cals; stub != nil {
				calsMock := pbservicestest.NewMockCalendarsClient(ctrl)
				calsMock.EXPECT().ConfirmEvent(tt.args.ctx, stub.req).Return(stub.resp, stub.err)
				cals = calsMock
			}

			s := &Server{}
			got, err := s.confirmPurchase(tt.args.ctx, db, tt.args.in, cals, prodsV, prodsM, tt.args.userId)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("Server.confirmPurchase() error = %v, wantErr %v", err, tt.wantStatus)
				return
			}
			if diff := cmp.Diff(tt.want, got.GetPurchase(), cmpopts.IgnoreTypes(timestamppb.Timestamp{})); diff != "" {
				t.Errorf("Server.confirmPurchase()  mismatch (-want +got): %s", diff)
			}

		})
	}
}
