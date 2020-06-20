package sm

import (
	"errors"
	"fmt"
)

type StateChanger func(sm *StateMachine, s *State) (*State, error)

func Next(sm *StateMachine, s *State) (*State, error) {
	if !s.nextable {
		return nil, errors.New("state is not nextable")
	}
	stage := sm.StageByName(s.Name)
	if stage == 0 {
		return nil, errors.New("state cannot be found on sm by name " + string(s.Name))
	}
	state := sm.StateByStage(stage + 1)
	if state == nil {
		return nil, fmt.Errorf("state cannot be found on sm by stage %d", stage+1)
	}
	return state, nil
}

func Prev(sm *StateMachine, s *State) (*State, error) {
	if !s.prevable {
		return nil, errors.New("state is not prevable")
	}
	stage := sm.StageByName(s.Name)
	if stage == 0 {
		return nil, errors.New("state cannot be found on sm by name " + string(s.Name))
	}
	state := sm.StateByStage(stage - 1)
	if state == nil {
		return nil, fmt.Errorf("state cannot be found on sm by stage %d", stage-1)
	}
	return state, nil
}

func Cancel(sm *StateMachine, s *State) (*State, error) {
	if !s.cancellable {
		return nil, errors.New("state is not cancellable")
	}
	return CancelledState, nil
}
