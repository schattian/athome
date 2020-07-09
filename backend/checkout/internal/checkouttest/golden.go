package checkouttest

import (
	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/xtest/xload"
)

func Init(gPurchases *GoldenPurchases) {
	xload.DecodeJsonnet("purchases", gPurchases)
}

type GoldenPurchases struct {
	Foo *purchase.Purchase `json:"foo,omitempty"`
}
