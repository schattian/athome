package srvpurchases

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/internal/checkouttest"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/test/pbaddresstest"
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

type addrMeasureStub struct {
	req  *pbaddress.MeasureDistanceRequest
	resp *pbaddress.MeasureDistanceResponse
	err  error
}
type addrRetrieveStub struct {
	req  *pbaddress.RetrieveAddressRequest
	resp *pbaddress.Address
	err  error
}
type addrStub struct {
	measure  *addrMeasureStub
	retrieve *addrRetrieveStub
}

func TestServer_assignAddress(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbcheckout.AssignDestAddressRequest
		o   *purchase.Purchase
	}
	tests := []struct {
		name         string
		args         args
		addrStub     *addrStub
		queries      []*sqlassist.QueryStubber
		execs        []*sqlassist.ExecStubber
		want         *emptypb.Empty
		wantPurchase *purchase.Purchase
		wantStatus   xerrors.Code
		wantErr      bool
	}{
		{
			name: "unauthorized:another user's address id",
			args: args{
				ctx: context.Background(),
				in:  &pbcheckout.AssignDestAddressRequest{DestAddressId: gPurchases.Foo.DestAddressId},
				o:   checkouttest.PurchaseZeroDestAddressId(t, gPurchases.Foo),
			},
			wantStatus:   xerrors.PermissionDenied,
			wantPurchase: checkouttest.PurchaseZeroDestAddressId(t, gPurchases.Foo),
			addrStub: &addrStub{
				retrieve: &addrRetrieveStub{
					req:  &pbaddress.RetrieveAddressRequest{AddressId: gPurchases.Foo.DestAddressId},
					resp: gPbAddresses.Merchants.Foo,
				},
				measure: &addrMeasureStub{},
			},
		},
		{
			name: "basic",
			args: args{
				ctx: context.Background(),
				in:  &pbcheckout.AssignDestAddressRequest{DestAddressId: gPurchases.Foo.DestAddressId},
				o:   checkouttest.PurchaseZeroDestAddressId(t, gPurchases.Foo),
			},
			wantPurchase: gPurchases.Foo,
			execs: []*sqlassist.ExecStubber{
				{
					Expect: "UPDATE purchases SET", Result: sqlmock.NewResult(1, 1),
				},
			},
			addrStub: &addrStub{
				retrieve: &addrRetrieveStub{
					req:  &pbaddress.RetrieveAddressRequest{AddressId: gPurchases.Foo.DestAddressId},
					resp: gPbAddresses.Consumers.Foo,
				},
				measure: &addrMeasureStub{
					req: &pbaddress.MeasureDistanceRequest{
						AAddressId: gPurchases.Foo.SrcAddressId,
						BAddressId: gPurchases.Foo.DestAddressId,
					},
					resp: &pbaddress.MeasureDistanceResponse{ManhattanInKilometers: gPurchases.Foo.DistanceInKilometers},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queries {
				stub.Stub(mock)
			}
			for _, stub := range tt.execs {
				stub.Stub(mock)
			}

			ctrl := gomock.NewController(t)
			var addr pbaddress.AddressesClient
			if stub := tt.addrStub; stub != nil {
				addrMock := pbaddresstest.NewMockAddressesClient(ctrl)
				addrMock.EXPECT().RetrieveAddress(tt.args.ctx, stub.retrieve.req).Return(stub.retrieve.resp, stub.retrieve.err)
				addrMock.EXPECT().MeasureDistance(tt.args.ctx, stub.measure.req).Return(stub.measure.resp, stub.measure.err)
				addr = addrMock
			}
			s := &Server{}
			tt.args.o = checkouttest.CopyPurchase(t, tt.args.o)
			_, err := s.assignAddress(tt.args.ctx, db, tt.args.in, addr, tt.args.o)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("Server.assignAddress() error = %v, wantErr %v", err, tt.wantStatus)
				return
			}
			if diff := cmp.Diff(tt.wantPurchase, tt.args.o, cmpopts.IgnoreTypes(timestamppb.Timestamp{})); diff != "" {
				t.Errorf("Server.assignAddress()  mismatch (-want +got): %s", diff)
			}
		})
	}
}
