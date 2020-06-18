package sm

import (
	"github.com/athomecomar/athome/pb/pbcheckout"
)

type StateMachine struct {
	States []*State
}

func (sm *StateMachine) StateByName(s string) *State {
	for _, st := range sm.States {
		if st.Name == s {
			return st
		}
	}
	return nil
}

func (s *State) ToPb(stage uint64) *pbcheckout.StateMachineResponse_StateDefinition {
	return &pbcheckout.StateMachineResponse_StateDefinition{
		Name:        s.Name,
		Stage:       stage,
		Description: s.Description,
	}
}
func (sm *StateMachine) ToPb() *pbcheckout.StateMachineResponse {
	var states []*pbcheckout.StateMachineResponse_StateDefinition
	for i, state := range sm.States {
		states = append(states, state.ToPb(uint64(i)+1))
	}
	return &pbcheckout.StateMachineResponse{States: states}
}

func (sm *StateMachine) StateByStage(s uint64) *State {
	if len(sm.States) < int(s) {
		return nil
	}
	return sm.States[s-1]
}

func (sm *StateMachine) First() *State {
	return sm.StateByStage(1)
}

func (sm *StateMachine) StageByName(s string) uint64 {
	for i, st := range sm.States {
		if st.Name == s {
			return uint64(i) + 1
		}
	}
	return 0
}

type State struct {
	Name        string
	Description string

	IsCancellable bool

	IsPrevable bool
	IsNextable bool
}

var (
	PurchaseStateMachine = &StateMachine{
		States: []*State{
			{Name: "address", Description: "address fulfill is needed", IsPrevable: true, IsNextable: true, IsCancellable: true},
		},
	}
)
