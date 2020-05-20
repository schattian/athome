
import {
  put,
} from 'redux-saga/effects';
import ApiService, { httpRequest } from 'src/services/api';
import { startLoading, stopLoading, LoaderTypes } from 'src/redux/loaders';
import { setUser } from '../user.actions';

function* fetchFakeData() {
  try {
    yield put(startLoading({ loaderType: LoaderTypes.FETCH_ME }));
    const response = yield httpRequest(ApiService.fetchMe);
    yield put(setUser(response.data));
    yield put(stopLoading({ loaderType: LoaderTypes.FETCH_ME }));
  } catch (error) {
    yield put(stopLoading({ loaderType: LoaderTypes.FETCH_ME, error }));
  }
}

export default fetchFakeData;
