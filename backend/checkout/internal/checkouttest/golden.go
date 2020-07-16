package checkouttest

import (
	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/ent/payment"
	"github.com/athomecomar/athome/backend/checkout/ent/shipping"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/xtest/xload"
)

func Init(gPurchases *GoldenPurchases, gShippings *GoldenShippings, gPbProducts *GoldenPbProducts, gStateChanges *GoldenStateChanges, gPbUsers *GoldenPbUsers,
	gPbAddresses *GoldenPbAddresses,
	gCards *GoldenCards, gPayments *GoldenPayments,
	gPbEvents *GoldenPbEvents,
	gPbServices *GoldenPbServices,
) {
	xload.DecodeJsonnet("purchases", gPurchases)
	xload.DecodeJsonnet("addresses", gPbAddresses)
	xload.DecodeJsonnet("users", gPbUsers)
	xload.DecodeJsonnet("products", gPbProducts)
	xload.DecodeJsonnet("shippings", gShippings)
	xload.DecodeJsonnet("state_changes", gStateChanges)
	xload.DecodeJsonnet("payments", gPayments)
	xload.DecodeJsonnet("cards", gCards)
	xload.DecodeJsonnet("events", gPbEvents)
	xload.DecodeJsonnet("services", gPbServices)
}

type GoldenPbServices struct {
	Delivery *pbservices.Service `json:"delivery,omitempty"`
}

type GoldenPbEvents struct {
	Foo *variadicPbEvents
}
type GoldenPayments struct {
	Purchases *variadicPayments `json:"purchases,omitempty"`
}

type GoldenCards struct {
	Foo *payment.Card
}

type GoldenStateChanges struct {
	Purchases *variadicPurchaseStateChanges `json:"purchases,omitempty"`
	Shippings *variadicShippingStateChanges `json:"shippings,omitempty"`
	Payments  *variadicPaymentStateChanges
}
type GoldenPbAddresses struct {
	Consumers        *variadicPbAddresses `json:"consumers,omitempty"`
	Merchants        *variadicPbAddresses `json:"merchants,omitempty"`
	ServiceProviders *variadicPbAddresses `json:"service_providers,omitempty"`
}

type GoldenPbUsers struct {
	Consumers        *variadicPbUsers `json:"consumers,omitempty"`
	Merchants        *variadicPbUsers `json:"merchants,omitempty"`
	ServiceProviders *variadicPbUsers `json:"service_providers,omitempty"`
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
