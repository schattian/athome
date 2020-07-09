package checkouttest

import (
	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/ent/shipping"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/xtest/xload"
)

func Init(gPurchases *GoldenPurchases, gShippings *GoldenShippings, gProducts *GoldenPbProducts) {
	xload.DecodeJsonnet("purchases", gPurchases)
	xload.DecodeJsonnet("products", gProducts)
	xload.DecodeJsonnet("shippings", gShippings)
}

type GoldenPurchases struct {
	Foo *purchase.Purchase `json:"foo,omitempty"`
}

type GoldenShippings struct {
	Foo *shipping.Shipping `json:"foo,omitempty"`
}
type GoldenPbProducts struct {
	Foo *variadicPbProducts `json:"foo,omitempty"`
	Bar *variadicPbProducts `json:"bar,omitempty"`
}
type variadicPbProducts struct {
	A *pbproducts.Product
	B *pbproducts.Product
}
