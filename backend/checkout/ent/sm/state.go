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
	Cancelled       StateName = "cancelled"
	PurchaseAddress StateName = "address"
)

var (
	CancelledState = &State{
		Name: Cancelled, Description: "cancelled order",
		prevable: false, nextable: false, cancellable: false,
	}
)
