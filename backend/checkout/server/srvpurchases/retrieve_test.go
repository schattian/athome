package srvpurchases

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/checkout/internal/checkouttest"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/test/pbproductstest"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	fooAProductId uint64 = 2134512
	fooBProductId uint64 = 432432432
)

func TestServer_retrieve(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbcheckout.RetrieveOrderRequest
		uid uint64
	}
	type prodsStub struct {
		req  *pbproducts.RetrieveProductsRequest
		resp *pbproducts.RetrieveProductsResponse
		err  error
	}

	tests := []struct {
		name string
		args args
		want *pbcheckout.Purchase

		prodsStub  *prodsStub
		queryStubs []*sqlassist.QueryStubber

		wantStatus xerrors.Code
	}{
		{
			name: "unauthorized",
			args: args{
				ctx: context.Background(),
				in:  &pbcheckout.RetrieveOrderRequest{OrderId: gPurchases.Foo.Id},
				uid: 2,
			},
			queryStubs: []*sqlassist.QueryStubber{
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
			},
			wantStatus: xerrors.PermissionDenied,
		},
		{
			name: "basic, consumer POV",
			args: args{
				ctx: context.Background(),
				in:  &pbcheckout.RetrieveOrderRequest{OrderId: gPurchases.Foo.Id},
				uid: gPurchases.Foo.UserId,
			},
			prodsStub: &prodsStub{
				req: &pbproducts.RetrieveProductsRequest{Ids: gPurchases.Foo.ProductIds()},
				resp: &pbproducts.RetrieveProductsResponse{
					Products: map[uint64]*pbproducts.Product{
						fooAProductId: gPbProducts.Foo.A,
						fooBProductId: gPbProducts.Foo.B,
					},
				},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM purchases WHERE id=$1",
					Args:   []driver.Value{gPurchases.Foo.Id},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gPurchases.Foo)).
						AddRow(storeql.SQLValues(gPurchases.Foo)...),
				},
			},
			want: checkouttest.PurchaseToPb(t,
				gPurchases.Foo,
				gPbProducts.Foo.A.Price*float64(gPurchases.Foo.Items[fooAProductId])+
					gPbProducts.Foo.B.Price*float64(gPurchases.Foo.Items[fooBProductId]),
			),
		},
		{
			name: "basic, merchant POV",
			args: args{
				ctx: context.Background(),
				in:  &pbcheckout.RetrieveOrderRequest{OrderId: gPurchases.Foo.Id},
				uid: gPurchases.Foo.MerchantId,
			},
			prodsStub: &prodsStub{
				req: &pbproducts.RetrieveProductsRequest{Ids: gPurchases.Foo.ProductIds()},
				resp: &pbproducts.RetrieveProductsResponse{
					Products: map[uint64]*pbproducts.Product{
						fooAProductId: gPbProducts.Foo.A,
						fooBProductId: gPbProducts.Foo.B,
					},
				},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM purchases WHERE id=$1",
					Args:   []driver.Value{gPurchases.Foo.Id},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gPurchases.Foo)).
						AddRow(storeql.SQLValues(gPurchases.Foo)...),
				},
			},
			want: checkouttest.PurchaseToPb(t,
				gPurchases.Foo,
				gPbProducts.Foo.A.Price*float64(gPurchases.Foo.Items[fooAProductId])+
					gPbProducts.Foo.B.Price*float64(gPurchases.Foo.Items[fooBProductId]),
			),
		},
		{
			name: "basic, deliverer POV",
			args: args{
				ctx: context.Background(),
				in:  &pbcheckout.RetrieveOrderRequest{OrderId: gPurchases.Foo.Id},
				uid: gShippings.Foo.UserId,
			},
			prodsStub: &prodsStub{
				req: &pbproducts.RetrieveProductsRequest{Ids: gPurchases.Foo.ProductIds()},
				resp: &pbproducts.RetrieveProductsResponse{
					Products: map[uint64]*pbproducts.Product{
						fooAProductId: gPbProducts.Foo.A,
						fooBProductId: gPbProducts.Foo.B,
					},
				},
			},
			queryStubs: []*sqlassist.QueryStubber{
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
			},
			want: checkouttest.PurchaseToPb(t,
				gPurchases.Foo,
				gPbProducts.Foo.A.Price*float64(gPurchases.Foo.Items[fooAProductId])+
					gPbProducts.Foo.B.Price*float64(gPurchases.Foo.Items[fooBProductId]),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}

			var prods pbproducts.ViewerClient
			if stub := tt.prodsStub; stub != nil {
				ctrl := gomock.NewController(t)
				prodsMock := pbproductstest.NewMockViewerClient(ctrl)
				prodsMock.EXPECT().RetrieveProducts(tt.args.ctx, stub.req).Return(stub.resp, stub.err)
				prods = prodsMock
			}

			s := &Server{}
			got, err := s.retrieve(tt.args.ctx, db, tt.args.in, prods, tt.args.uid)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("Server.retrieve() error = %v, wantErr %v", err, tt.wantStatus)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreUnexported(timestamppb.Timestamp{})); diff != "" {
				t.Errorf("Server.retrieve() mismatch (-want +got): %s", diff)
			}
		})
	}
}
