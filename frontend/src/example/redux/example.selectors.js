import { createSelector } from 'reselect';

export const getState = (state) => state.example;

export const getFakeData = createSelector(getState, (exampleState) => (
  exampleState.fakeData
));
