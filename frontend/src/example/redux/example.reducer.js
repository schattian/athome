/* eslint-disable operator-assignment */
/* eslint-disable no-param-reassign */
import produce from 'immer';
import { actionsType } from './example.actions';

const initialState = {
  fakeData: null,
};

export default function exampleReducer(state = initialState, action) {
  switch (action.type) {
    case actionsType.SET_FAKE_DATA: {
      const nextState = produce(state, (draftState) => {
        draftState.fakeData = action.payload;
      });
      return nextState;
    }
    default:
      return state;
  }
}
