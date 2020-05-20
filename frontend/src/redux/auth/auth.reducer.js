/* eslint-disable operator-assignment */
/* eslint-disable no-param-reassign */
import produce from 'immer';
import { actionsType } from './auth.actions';

const initialState = {
  token: null,
  loading: null, // Keep it null for initial, and not false to detect if check token end
  valid: null,
  lastAction: null,
};

export default function authReducer(state = initialState, action) {
  switch (action.type) {
    case actionsType.SET_TOKEN: {
      const nextState = produce(state, (draftState) => {
        draftState.token = action.payload;
        draftState.lastAction = action.type;
      });
      return nextState;
    }
    case actionsType.ON_CHECK_TOKEN_START: {
      const nextState = produce(state, (draftState) => {
        draftState.loading = true;
        draftState.valid = false;
        draftState.lastAction = action.type;
      });
      return nextState;
    }
    case actionsType.ON_CHECK_TOKEN_END: {
      const nextState = produce(state, (draftState) => {
        draftState.loading = false;
        draftState.valid = true;
        draftState.lastAction = action.type;
      });
      return nextState;
    }
    case actionsType.ON_CHECK_TOKEN_FAILED:
    case actionsType.ON_CHECK_TOKEN_FAILED_NETWORK_ERROR: {
      const nextState = produce(state, (draftState) => {
        draftState.loading = false;
        draftState.valid = false;
        draftState.token = null;
        draftState.lastAction = action.type;
      });
      return nextState;
    }
    case actionsType.ON_LOGIN_END:
    case actionsType.ON_REGISTER_END: {
      const nextState = produce(state, (draftState) => {
        draftState.loading = false;
        draftState.valid = true;
        draftState.token = action.token;
        draftState.lastAction = action.type;
      });
      return nextState;
    }
    case actionsType.ON_LOGOUT_END: {
      const nextState = produce(state, (draftState) => {
        draftState.loading = initialState.loading;
        draftState.valid = initialState.valid;
        draftState.token = initialState.token;
        draftState.lastAction = action.type;
      });
      return nextState;
    }
    case actionsType.SET_LOADING: {
      const nextState = produce(state, (draftState) => {
        draftState.lastAction = action.type;
        draftState.loading = action.payload;
      });
      return nextState;
    }
    default:
      return state;
  }
}
