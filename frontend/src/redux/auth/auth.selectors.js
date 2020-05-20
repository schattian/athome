import { createSelector } from 'reselect';

export const getState = (state) => state.auth;

export const getToken = createSelector(getState, (authState) => (
  authState.token
));
export const isAuthenticated = createSelector(getState, (authState) => authState.valid);
export const getTokenValidateState = createSelector(getState, (authState) => (
  authState.valid
));
export const isLoading = createSelector(getState, (authState) => (
  authState.loading
));
export const getLastAction = createSelector(getState, (authState) => (
  authState.lastAction
));
