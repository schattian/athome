/* eslint-disable no-param-reassign */
import produce from 'immer';
import { actionsType } from './loaders.actions';
import { LoaderStatus, LoaderTypes } from './loaderTypesEnum';

const initialState = {};
Object.keys(LoaderTypes).forEach((key) => {
  initialState[key] = { status: null, error: null };
});

export default function loadersReducer(state = initialState, action) {
  const { payload } = action;
  const { loaderType, error } = payload || {};
  switch (action.type) {
    case actionsType.START_LOADING: {
      const nextState = produce(state, (draftState) => {
        draftState[loaderType] = { status: LoaderStatus.LOADING, error: null };
      });
      return nextState;
    }
    case actionsType.STOP_LOADING: {
      const nextState = produce(state, (draftState) => {
        // eslint-disable-next-line no-underscore-dangle
        const _error = (error ? (error.message || error) : null);
        draftState[loaderType] = {
          status: _error ? LoaderStatus.ERROR : LoaderStatus.END,
          error: _error,
        };
      });
      return nextState;
    }
    default:
      return state;
  }
}
