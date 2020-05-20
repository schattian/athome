import { createSelector } from 'reselect';

export const getState = (state) => state.user;

export const getUser = createSelector(getState, (userState) => (
  userState.user
));
