import {
  takeEvery,
} from 'redux-saga/effects';
import { actionsType } from './user.actions';

import fetchMe from './workers/fetchMe.workers';
import updateMe from './workers/updateMe.workers';

export default function* globalWatcher() {
  yield takeEvery(actionsType.FETCH_ME, fetchMe);
  yield takeEvery(actionsType.UPDATE_ME, updateMe);
}
