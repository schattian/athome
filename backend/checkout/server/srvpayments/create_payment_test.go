package srvpayments

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/internal/checkouttest"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbutil"
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

type prodViewerStub struct {
	req  *pbproducts.RetrieveProductsRequest
	resp *pbproducts.RetrieveProductsResponse
	err  error
}

func TestServer_createPurchasePayment(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbcheckout.CreatePaymentRequest
		p   *purchase.Purchase
	}
	tests := []struct {
		name       string
		prodV      *prodViewerStub
		args       args
		queries    []*sqlassist.QueryStubber
		want       *pbcheckout.CreatePaymentResponse
		wantStatus xerrors.Code
	}{
		{
			name: "basic",
			queries: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM shippings WHERE id=$1",
					Args:   []driver.Value{gPurchases.Foo.ShippingId},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gShippings.Foo)).
						AddRow(storeql.SQLValues(gShippings.Foo)...),
				},
				{
					Expect: "SELECT * FROM cards WHERE id=$1 AND user_id=$2",
					Args:   []driver.Value{gPayments.Purchases.Foo.CardId, gPayments.Purchases.Foo.UserId},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gCards.Foo)).
						AddRow(storeql.SQLValues(gCards.Foo)...),
				},
				{
					Expect: storeql.ExecBoilerplate("INSERT", gPayments.Purchases.Foo),
					Rows: sqlmock.NewRows([]string{"id"}).
						AddRow(gPayments.Purchases.Foo.Id),
				},
				{
					Expect: storeql.ExecBoilerplate("INSERT", gStateChanges.Payments.Created),
					Rows: sqlmock.NewRows([]string{"id"}).
						AddRow(gStateChanges.Payments.Created.Id),
				},
			},
			want: &pbcheckout.CreatePaymentResponse{Payment: checkouttest.PaymentToPb(t, gPayments.Purchases.Foo), PaymentId: gPayments.Purchases.Foo.Id},
			args: args{
				ctx: context.Background(),
				in: &pbcheckout.CreatePaymentRequest{
					Payment: &pbcheckout.PaymentInput{
						PaymentMethodId: gPayments.Purchases.Foo.PaymentMethodId,
						Installments:    gPayments.Purchases.Foo.Installments,
						CardId:          gPayments.Purchases.Foo.CardId,
						Order:           pbutil.ToPbEntity(gPurchases.Foo),
					},
				},
				p: gPurchases.Foo,
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queries {
				stub.Stub(mock)
			}
			ctrl := gomock.NewController(t)
			var prodsV pbproducts.ViewerClient
			if stub := tt.prodV; stub != nil {
				prodsMock := pbproductstest.NewMockViewerClient(ctrl)
				prodsMock.EXPECT().RetrieveProducts(tt.args.ctx, stub.req).Return(stub.resp, stub.err)
				prodsV = prodsMock
			}

			s := &Server{}
			got, err := s.createPurchasePayment(tt.args.ctx, db, tt.args.in, prodsV, tt.args.p)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("Server.createPurchasePayment() error = %v, wantErr %v", err, tt.wantStatus)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreTypes(timestamppb.Timestamp{})); diff != "" {
				t.Errorf("Server.createPurchasePayment()  mismatch (-want +got): %s", diff)
			}

		})
	}
}
