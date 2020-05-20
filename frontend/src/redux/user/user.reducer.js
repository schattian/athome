/* eslint-disable operator-assignment */
/* eslint-disable no-param-reassign */
import produce from 'immer';
import { actionsType } from './user.actions';

const initialState = {
  user: null,
};

export default function userReducer(state = initialState, action) {
  switch (action.type) {
    case actionsType.SET_USER: {
      const nextState = produce(state, (draftState) => {
        draftState.user = action.payload;
      });
      return nextState;
    }
    case actionsType.REMOVE_USER: {
      const nextState = produce(state, (draftState) => {
        draftState.user = initialState.user;
      });
      return nextState;
    }

    default:
      return state;
  }
}
