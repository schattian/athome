import {
  takeEvery, put, delay, select,
} from 'redux-saga/effects';
import { actionsType } from './global.actions';
import { getCount } from './global.selectors';
import { startLoading, stopLoading, LoaderTypes } from '../loaders';

function* onCountChange() {
  try {
    const current = yield select(getCount);
    yield put(startLoading({ loaderType: LoaderTypes.ON_COUNT_CHANGE }));
    yield delay(500);
    if (current === 5) {
      throw new Error('throw error on 5');
    }
    yield put(stopLoading({ loaderType: LoaderTypes.ON_COUNT_CHANGE }));
  } catch (error) {
    yield put(stopLoading({ loaderType: LoaderTypes.ON_COUNT_CHANGE, error }));
  }
}

export default function* globalWatcher() {
  yield takeEvery(actionsType.INCREMENT, onCountChange);
  yield takeEvery(actionsType.DECREMENT, onCountChange);
}
