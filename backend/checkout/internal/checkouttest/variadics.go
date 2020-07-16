package checkouttest

import (
	"github.com/athomecomar/athome/backend/checkout/ent/payment"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbusers"
)

type variadicPbUsers struct {
	Foo *pbusers.User `json:"foo,omitempty"`
	Bar *pbusers.User `json:"bar,omitempty"`
}
type variadicPaymentStateChanges struct {
	Cancelled *sm.StateChange `json:"cancelled,omitempty"`
	Created   *sm.StateChange `json:"created,omitempty"`
	Finished  *sm.StateChange `json:"finished,omitempty"`
	Rejected  *sm.StateChange
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
type variadicPayments struct {
	Foo *payment.Payment
}
type variadicShippingStateChanges struct {
	Cancelled *sm.StateChange `json:"cancelled,omitempty"`
	Created   *sm.StateChange `json:"created,omitempty"`
	Taken     *sm.StateChange
	Finished  *sm.StateChange `json:"finished,omitempty"`
}

type variadicPbProducts struct {
	A *pbproducts.Product `json:"a,omitempty"`
	B *pbproducts.Product `json:"b,omitempty"`
}

type variadicPbAddresses struct {
	Foo *pbaddress.Address
}
type variadicPbEvents struct {
	Medic    *variationEvent
	Delivery *variationEvent
}

type variationEvent struct {
	First *subVariationEvents
}
type subVariationEvents struct {
	A *pbservices.Event
	B *pbservices.Event
	C *pbservices.Event
}
