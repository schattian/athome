package sm

import (
	"github.com/athomecomar/athome/pb/pbcheckout"
)

type StateMachine struct {
	States []*State
}

func (sm *StateMachine) StateByName(s StateName) *State {
	if s == Cancelled {
		return CancelledState
	}

	for _, st := range sm.States {
		if st.Name == s {
			return st
		}
	}
	return nil
}

func (sm *StateMachine) ToPb() *pbcheckout.StateMachineResponse {
	var states []*pbcheckout.StateMachineResponse_StateDefinition
	for i, state := range sm.States {
		states = append(states, state.ToPb(int64(i)+1))
	}
	return &pbcheckout.StateMachineResponse{States: states}
}

func (sm *StateMachine) StateByStage(s int64) *State {
	if s == -1 {
		return CancelledState
	}

	if len(sm.States) < int(s) {
		return nil
	}
	return sm.States[s-1]
}

func (sm *StateMachine) First() *State {
	return sm.StateByStage(1)
}

func (sm *StateMachine) StageByName(s StateName) int64 {
	if s == Cancelled {
		return -1
	}
	for i, st := range sm.States {
		if st.Name == s {
			return int64(i) + 1
		}
	}
	return 0
}

var (
	PurchaseStateMachine = &StateMachine{
		States: []*State{
			{Name: PurchaseInit, Description: "draft was initialized", prevable: true, nextable: true, cancellable: true},
			{Name: PurchaseAddress, Description: "address fulfill is completed", prevable: true, nextable: true, cancellable: true},
			{Name: PurchaseShippingMethod, Description: "shipping method was selected", prevable: true, nextable: true, cancellable: true},
			{Name: PurchasePayment, Description: "purchase was paid", prevable: false, nextable: false, cancellable: false},
			{Name: PurchaseConfirmed, Description: "purchase was confirmed by merchant", prevable: false, nextable: false, cancellable: false},
			{Name: PurchaseInDelivery, Description: "purchase is on deliveror's hand", prevable: false, nextable: false, cancellable: false}, // only exists when shipments = 1
			{Name: PurchaseFinished, Description: "purchase is finished", prevable: false, nextable: false, cancellable: false},
		},
	}
)
