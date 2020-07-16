package srvshippings

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/internal/checkouttest"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/test/pbservicestest"
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

type calendarStub struct {
	req  *pbservices.CreateShippingEventRequest
	resp *pbservices.CreateEventResponse
	err  error
}

type servicesStub struct {
	req  *pbservices.RetrieveServiceRequest
	resp *pbservices.Service
	err  error
}

func TestServer_createShipping(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbcheckout.CreateShippingRequest
		p   *purchase.Purchase
	}
	type stubs struct {
		exec  []*sqlassist.ExecStubber
		query []*sqlassist.QueryStubber
		cals  *calendarStub
		svcs  *servicesStub
	}
	tests := []struct {
		name  string
		args  args
		stubs stubs

		want       *pbcheckout.CreateShippingResponse
		wantStatus xerrors.Code
	}{
		{
			name: "basic",
			args: args{
				ctx: context.Background(),
				in: &pbcheckout.CreateShippingRequest{
					Shipping: &pbcheckout.CreateShippingRequest_Shipping{
						ShippingMethodId: gShippings.Foo.ShippingMethodId,
						Time:             gPbEvents.Foo.Delivery.First.A.End,
						Dow:              gPbEvents.Foo.Delivery.First.A.Dow,
					},
				},
				p: checkouttest.PurchaseZeroShippingId(t, gPurchases.Foo),
			},

			stubs: stubs{
				query: []*sqlassist.QueryStubber{
					{
						Expect: "SELECT * FROM shippings WHERE user_id=$1",
						Err:    sql.ErrNoRows,
					},
					{
						Expect: storeql.ExecBoilerplate("INSERT", gShippings.Foo),
						Rows: sqlmock.NewRows([]string{"id"}).
							AddRow(gShippings.Foo.Id),
					},
					{
						Expect: storeql.ExecBoilerplate("INSERT", gStateChanges.Shippings.Created),
						Rows: sqlmock.NewRows([]string{"id"}).
							AddRow(gStateChanges.Shippings.Created.Id),
					},
				},
				exec: []*sqlassist.ExecStubber{
					{
						Expect: "UPDATE purchases SET",
						Result: sqlmock.NewResult(1, 1),
					},
				},
				cals: &calendarStub{
					req: &pbservices.CreateShippingEventRequest{
						End:       gPbEvents.Foo.Delivery.First.A.End,
						Dow:       gPbEvents.Foo.Delivery.First.A.Dow,
						ServiceId: gShippings.Foo.ShippingMethodId,
					},
					resp: &pbservices.CreateEventResponse{
						EventId: gShippings.Foo.EventId,
						Event:   gPbEvents.Foo.Delivery.First.A,
					},
				},
				svcs: &servicesStub{
					req: &pbservices.RetrieveServiceRequest{
						ServiceId: gShippings.Foo.ShippingMethodId,
					},
					resp: gPbServices.Delivery,
				},
			},
			want: &pbcheckout.CreateShippingResponse{
				ShippingId: gShippings.Foo.Id,
				Shipping:   gShippings.Foo.ToPb(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.stubs.query {
				stub.Stub(mock)
			}
			for _, stub := range tt.stubs.exec {
				stub.Stub(mock)
			}

			ctrl := gomock.NewController(t)
			var cals pbservices.CalendarsClient
			var svcs pbservices.ViewerClient
			if stub := tt.stubs.cals; stub != nil {
				calsMock := pbservicestest.NewMockCalendarsClient(ctrl)
				calsMock.EXPECT().CreateShippingEvent(tt.args.ctx, stub.req).Return(stub.resp, stub.err)
				cals = calsMock
			}
			if stub := tt.stubs.svcs; stub != nil {
				svcsMock := pbservicestest.NewMockViewerClient(ctrl)
				svcsMock.EXPECT().RetrieveService(tt.args.ctx, stub.req).Return(stub.resp, stub.err)
				svcs = svcsMock
			}

			s := &Server{}
			got, err := s.createShipping(tt.args.ctx, db, tt.args.in, svcs, cals, tt.args.p)
			if status.Code(err) != tt.wantStatus {
				t.Errorf("Server.createShipping() error = %v, wantErr %v", err, tt.wantStatus)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreTypes(timestamppb.Timestamp{})); diff != "" {
				t.Errorf("Server.createShipping()  mismatch (-want +got): %s", diff)
			}
		})
	}
}
