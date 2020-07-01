package sm

import "github.com/athomecomar/athome/pb/pbcheckout"

type StateName string

type State struct {
	Name        StateName
	Description string

	cancellable bool
	prevable    bool
	nextable    bool
}

func (s *State) ToPb(stage int64) *pbcheckout.StateMachineResponse_StateDefinition {
	return &pbcheckout.StateMachineResponse_StateDefinition{
		Name:        string(s.Name),
		Stage:       stage,
		Description: s.Description,
	}
}

const (
	Cancelled StateName = "cancelled"

	// purchase is created, and is redirected to address
	PurchaseCreated StateName = "purchase:created"
	// purchase has an address assigned
	PurchaseAddressed StateName = "purchase:addressed"
	// purchase has a shipping method selected
	PurchaseShippingMethodSelected StateName = "purchase:shipping_method_selected"
	// purchase has been paid
	PurchasePaid StateName = "purchase:paid"
	// purchase has been confirmed by the merchant
	PurchaseConfirmed StateName = "purchase:confirmed"
	// purchase has been shipped successfully
	PurchaseShipped StateName = "purchase:shipped"
	// purchase has finished successfully
	PurchaseFinished StateName = "purchase:finished"

	// shipping has been created, and is waiting to be dispatched
	ShippingCreated StateName = "shipping:created"
	// shipping has been dispatched by the merchant, and is ready to be taken
	ShippingDispatched StateName = "shipping:dispatched"
	// shipping has been taken and the deliverer is going to the consumer
	ShippingTaken StateName = "shipping:taken"
	// shipping has finished successfully
	ShippingFinished StateName = "shipping:finished"

	// payment has been created, and is being processed
	PaymentCreated StateName = "payment:created"
	// payments was paid successfully
	PaymentFinished StateName = "payment:finished"
	// payment was rejected due some reason
	PaymentRejected StateName = "payment:rejected"
)

var (
	CancelledState = &State{
		Name: Cancelled, Description: "cancelled order",
		prevable: false, nextable: false, cancellable: false,
	}
)
