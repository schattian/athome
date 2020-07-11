package srvpurchases

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/checkout/internal/checkouttest"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/athome/pb/test/pbproductstest"
	"github.com/athomecomar/athome/pb/test/pbuserstest"
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

type prodViewerStub struct {
	req  *pbproducts.RetrieveProductsRequest
	resp *pbproducts.RetrieveProductsResponse
	err  error
}

type prodManagerStub struct {
	req  *pbproducts.ReserveStockRequest
	resp *emptypb.Empty
	err  error
}
type addrStub struct {
	req  *pbaddress.MeasureDistanceRequest
	resp *pbaddress.MeasureDistanceResponse
	err  error
}

type userStub struct {
	req  *pbusers.RetrieveUserRequest
	resp *pbusers.User
	err  error
}

func TestServer_createPurchase(t *testing.T) {
	type stubs struct {
		query []*sqlassist.QueryStubber
		prodV *prodViewerStub
		prodM *prodManagerStub
		user  []*userStub
	}
	type args struct {
		ctx    context.Context
		in     *pbcheckout.CreatePurchaseRequest
		userId uint64
	}
	tests := []struct {
		name string
		args args
		want *pbcheckout.CreatePurchaseResponse

		stubs      stubs
		wantStatus xerrors.Code
	}{
		{
			name: "unauthorized due role != consumer",
			args: args{
				ctx:    context.Background(),
				in:     &pbcheckout.CreatePurchaseRequest{Items: gPurchases.Foo.Items},
				userId: gPurchases.Foo.UserId,
			},
			stubs: stubs{
				user: []*userStub{
					{
						req:  &pbusers.RetrieveUserRequest{UserId: gPurchases.Foo.UserId},
						resp: gPbUsers.Merchants.Foo,
					},
				},
			},
			wantStatus: xerrors.PermissionDenied,
		},
		{
			name: "basic",
			args: args{
				ctx:    context.Background(),
				in:     &pbcheckout.CreatePurchaseRequest{Items: gPurchases.Foo.Items},
				userId: gPurchases.Foo.UserId,
			},
			stubs: stubs{
				query: []*sqlassist.QueryStubber{
					{
						Expect: storeql.ExecBoilerplate("INSERT", gPurchases.Foo),
						Rows: sqlmock.NewRows([]string{"id"}).
							AddRow(gPurchases.Foo.Id),
					},
					{
						Expect: storeql.ExecBoilerplate("INSERT", gStateChanges.Purchases.Created),
						Rows: sqlmock.NewRows([]string{"id"}).
							AddRow(gStateChanges.Purchases.Created.Id),
					},
				},
				user: []*userStub{
					{
						req:  &pbusers.RetrieveUserRequest{UserId: gPurchases.Foo.UserId},
						resp: gPbUsers.Consumers.Foo,
					},
					{
						req:  &pbusers.RetrieveUserRequest{UserId: gPurchases.Foo.MerchantId},
						resp: gPbUsers.Merchants.Foo,
					},
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
				prodM: &prodManagerStub{
					req:  &pbproducts.ReserveStockRequest{Order: pbutil.ToPbEntity(gPurchases.Foo)},
					resp: &emptypb.Empty{},
				},
			},
			want: &pbcheckout.CreatePurchaseResponse{
				PurchaseId: gPurchases.Foo.Id,
				Purchase: checkouttest.PurchaseToPb(
					t, checkouttest.PurchaseCreation(t, gPurchases.Foo),
					gPbProducts.Foo.A.Price*float64(gPurchases.Foo.Items[fooAProductId])+
						gPbProducts.Foo.B.Price*float64(gPurchases.Foo.Items[fooBProductId]),
				)},
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
			var users pbusers.ViewerClient
			if tt.stubs.user != nil {
				userMock := pbuserstest.NewMockViewerClient(ctrl)
				for _, stub := range tt.stubs.user {
					userMock.EXPECT().RetrieveUser(tt.args.ctx, stub.req).Return(stub.resp, stub.err)
				}
				users = userMock
			}

			var prodsM pbproducts.ManagerClient
			if stub := tt.stubs.prodM; stub != nil {
				prodMMock := pbproductstest.NewMockManagerClient(ctrl)
				prodMMock.EXPECT().CreateReserveStock(tt.args.ctx, stub.req).Return(stub.resp, stub.err)
				prodsM = prodMMock
			}

			s := &Server{}
			got, err := s.createPurchase(tt.args.ctx, db, tt.args.in, users, prodsV, prodsM, tt.args.userId)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("Server.createPurchase() error = %v, wantErr %v", err, tt.wantStatus)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreTypes(timestamppb.Timestamp{})); diff != "" {
				t.Errorf("Server.createPurchase()  mismatch (-want +got): %s", diff)
			}
		})
	}
}

// var addr pbaddress.AddressesClient
// if stub := tt.stubs.addr; stub != nil {
// 	addrMock := pbaddresstest.NewMockAddressesClient(ctrl)
// 	addrMock.EXPECT().MeasureDistance(tt.args.ctx, stub.req).Return(stub.resp, stub.err)
// 	addr = addrMock
// }
// addr: &addrStub{
// 	req: &pbaddress.MeasureDistanceRequest{
// 		AAddressId: gPbUsers.Merchants.Foo.AddressId,
// 		BAddressId: gPurchases.Foo.SrcAddressId,
// 	},
// 	resp: &pbaddress.MeasureDistanceResponse{
// 		ManhattanInKilometers: gPurchases.Foo.DistanceInKilometers,
// 	},
// },
