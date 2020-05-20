/* eslint-disable operator-assignment */
/* eslint-disable no-param-reassign */
import produce from 'immer';
import { actionsType } from './global.actions';

const initialState = {
  count: 0,
  notifications: [],
  windowWidth: null,
  windowHeight: null,
  networkOnline: null,
};

export default function globalReducer(state = initialState, action) {
  switch (action.type) {
    case actionsType.INCREMENT: {
      const nextState = produce(state, (draftState) => {
        draftState.count = draftState.count + 1;
      });
      return nextState;
    }
    case actionsType.DECREMENT: {
      const nextState = produce(state, (draftState) => {
        draftState.count = draftState.count - 1;
      });
      return nextState;
    }
    case actionsType.SEND_NOTIFICATION: {
      const nextState = produce(state, (draftState) => {
        draftState.notifications.push({
          key: action.key,
          ...action.notification,
        });
      });
      return nextState;
    }
    case actionsType.CLOSE_NOTIFICATION: {
      const nextState = produce(state, (draftState) => {
        draftState.notifications = draftState.notifications.map((notification) => (
          (action.dismissAll || notification.key === action.key)
            ? { ...notification, dismissed: true }
            : { ...notification }
        ));
      });
      return nextState;
    }
    case actionsType.REMOVE_NOTIFICATION: {
      const nextState = produce(state, (draftState) => {
        draftState.notifications = draftState.notifications.filter(
          (notification) => notification.key !== action.key,
        );
      });
      return nextState;
    }
    case actionsType.SET_WINDOW_SIZE: {
      const nextState = produce(state, (draftState) => {
        draftState.windowWidth = action.width;
        draftState.windowHeight = action.height;
      });
      return nextState;
    }
    case actionsType.SET_NETWORK_ONLINE: {
      const nextState = produce(state, (draftState) => {
        draftState.networkOnline = action.payload;
      });
      return nextState;
    }
    default:
      return state;
  }
}
