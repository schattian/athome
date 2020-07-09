package checkouttest

import (
	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/ent/shipping"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/xtest/xload"
)

func Init(gPurchases *GoldenPurchases, gShippings *GoldenShippings, gPbProducts *GoldenPbProducts, gStateChanges *GoldenStateChanges) {
	xload.DecodeJsonnet("purchases", gPurchases)
	xload.DecodeJsonnet("products", gPbProducts)
	xload.DecodeJsonnet("shippings", gShippings)
	xload.DecodeJsonnet("state_changes", gStateChanges)
}

type GoldenStateChanges struct {
	Purchases *variadicPurchaseStateChanges `json:"purchases,omitempty"`
}

type variadicPurchaseStateChanges struct {
	Cancelled              *sm.StateChange `json:"cancelled,omitempty"`
	Created                *sm.StateChange `json:"created,omitempty"`
	Addressed              *sm.StateChange `json:"addressed,omitempty"`
	ShippingMethodSelected *sm.StateChange `json:"shipping_method_selected,omitempty"`
	Paid                   *sm.StateChange `json:"paid,omitempty"`
	Confirmed              *sm.StateChange `json:"confirmed,omitempty"`
	Finished               *sm.StateChange `json:"finished,omitempty"`
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
	A *pbproducts.Product `json:"a,omitempty"`
	B *pbproducts.Product `json:"b,omitempty"`
}
