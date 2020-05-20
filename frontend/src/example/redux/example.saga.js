
import { takeEvery } from 'redux-saga/effects';
import { actionsType } from './example.actions';
import fetchFakeData from './workers/fetchFakeData.worker';

export default function* exampleWatcher() {
  yield takeEvery(actionsType.FETCH_FAKE_DATA, fetchFakeData);
}
