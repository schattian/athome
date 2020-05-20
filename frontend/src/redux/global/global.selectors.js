import { createSelector } from 'reselect';

export const getState = (state) => state.global;

export const getCount = createSelector(getState, (globalState) => (
  globalState.count
));
export const isOnline = createSelector(getState, (globalState) => (
  globalState.networkOnline
));
export const getWindowSize = createSelector(getState, (globalState) => ({
  windowWidth: globalState.windowWidth,
  windowHeight: globalState.windowHeight,
}));
