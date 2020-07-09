package srvpurchases

import (
	"context"
	"database/sql/driver"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/test/pbproductstest"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/status"
)

func TestServer_retrieve(t *testing.T) {
	type args struct {
		ctx   context.Context
		in    *pbcheckout.RetrieveOrderRequest
		prods pbproducts.ViewerClient
		uid   uint64
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
			name: "basic",
			args: args{
				ctx: context.Background(),
				in:  &pbcheckout.RetrieveOrderRequest{},
				uid: gPurchases.Foo.UserId,
			},
			prodsStub: &prodsStub{
				req: &pbproducts.RetrieveProductsRequest{Ids: gPurchases.Foo.ProductIds()},
				resp: &pbproducts.RetrieveProductsResponse{
					Products: map[uint64]*pbproducts.Product{
						gPurchases.Foo.ProductIds()[0]: gPbProducts.Foo.A,
						gPurchases.Foo.ProductIds()[1]: gPbProducts.Foo.B,
					},
				},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM purchases WHERE id = $1",
					Args:   []driver.Value{gPurchases.Foo.Id},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gPurchases.Foo)).
						AddRow(storeql.SQLValues(gPurchases.Foo)...),
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}
