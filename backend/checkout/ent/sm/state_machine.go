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
	if s == -2 {
		return RejectedState
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
			{
				Name: PurchaseCreated, Description: "draft was initialized",
				consumer: all,
			},
			{
				Name: PurchaseAddressed, Description: "address fulfill is completed",
				consumer: all,
			},
			{
				Name: PurchaseShippingMethodSelected, Description: "shipping method was selected",
				consumer: all,
			},
			{
				Name: PurchasePaid, Description: "purchase was paid",
				merchant: onlyNext,
				consumer: onlyCancel,
			},
			{
				Name: PurchaseConfirmed, Description: "purchase was confirmed by merchant",
				//  state change validation is bind to shipping's one
				consumer: onlyNext,
			},
			{
				Name: PurchaseFinished, Description: "purchase is finished",
			},
		},
	}

	ShippingStateMachine = &StateMachine{
		States: []*State{
			{
				Name: ShippingCreated, Description: "waiting for being dispatched",
				serviceProvider: onlyNext, // agreed
			},
			{
				Name: ShippingTaken, Description: "the deliverer is coming",
				consumer: onlyNext, // agreed
			},
			{
				Name: ShippingFinished, Description: "shipping was finished",
			},
		},
	}

	PaymentStateMachine = &StateMachine{
		States: []*State{
			{
				Name: PaymentCreated, Description: "being processed",
			},
			{
				Name: PaymentFinished, Description: "payment was paid successfully",
			},
		},
	}
)
