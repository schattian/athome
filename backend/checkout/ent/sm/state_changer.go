package sm

import (
	"errors"
	"fmt"
)

type StateChanger func(sm *StateMachine, s *State, entity Stateful, uid uint64) (*State, error)

func Next(sm *StateMachine, s *State, entity Stateful, uid uint64) (*State, error) {
	permissions := s.userPermissions(entity, uid)
	if !permissions.nextable {
		return nil, errors.New("you are not authorized to next this")
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

func Prev(sm *StateMachine, s *State, entity Stateful, uid uint64) (*State, error) {
	permissions := s.userPermissions(entity, uid)
	if !permissions.prevable {
		return nil, errors.New("you are not authorized to prev this")
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

func Cancel(sm *StateMachine, s *State, entity Stateful, uid uint64) (*State, error) {
	permissions := s.userPermissions(entity, uid)
	if !permissions.cancellable {
		return nil, errors.New("you are not authorized to cancel this")
	}
	return CancelledState, nil
}
