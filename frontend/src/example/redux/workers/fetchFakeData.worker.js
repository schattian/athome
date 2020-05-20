
import {
  put,
} from 'redux-saga/effects';
import ApiService, { httpRequest } from 'src/services/api';
import { startLoading, stopLoading, LoaderTypes } from 'src/redux/loaders';
import { setFakeData } from '../example.actions';

function* fetchFakeData() {
  try {
    yield put(startLoading({ loaderType: LoaderTypes.FETCH_FAKE_DATA }));
    const response = yield httpRequest(ApiService.fetchFakeData);
    yield put(setFakeData(response.data));
    yield put(stopLoading({ loaderType: LoaderTypes.FETCH_FAKE_DATA }));
  } catch (error) {
    yield put(stopLoading({ loaderType: LoaderTypes.FETCH_FAKE_DATA, error }));
  }
}

export default fetchFakeData;
