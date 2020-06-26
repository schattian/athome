package srvmanager

import (
	"context"

	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbshared"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func retrieveOrder(ctx context.Context, orderEnt *pbshared.Entity, access string) (ord order, err error) {
	cli, err := grpc.Dial(pbconf.Checkout.GetHost(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "grpc.Dial")
	}

	req := &pbcheckout.RetrieveOrderRequest{OrderId: orderEnt.EntityId, AccessToken: access}
	switch orderEnt.EntityTable {
	case "purchases":
		ord, err = pbcheckout.NewPurchasesClient(cli).Retrieve(ctx, req)
	case "reservations":
		ord, err = pbcheckout.NewReservationsClient(cli).Retrieve(ctx, req)
	}
	return
}

func getItemIds(order order) []uint64 {
	var ids []uint64
	for id := range order.GetItems() {
		ids = append(ids, id)
	}
	return ids
}

type order interface {
	GetUserId() uint64
	GetMerchantId() uint64
	GetItems() map[uint64]uint64
}
