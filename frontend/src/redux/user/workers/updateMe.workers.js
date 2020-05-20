import {
  put,
} from 'redux-saga/effects';
import ApiService, { httpRequest } from 'src/services/api';
import { startLoading, stopLoading, LoaderTypes } from 'src/redux/loaders';
import { setUser } from '../user.actions';

function* updateMe(action) {
  try {
    yield put(startLoading({ loaderType: LoaderTypes.UPDATE_ME }));
    let response;
    response = yield httpRequest(ApiService.updateMe, action.payload);
    // eslint-disable-next-line no-underscore-dangle
    if (!response.data || !response.data._id) {
      response = yield httpRequest(ApiService.fetchMe, action.payload);
    }
    yield put(setUser(response.data));
    yield put(stopLoading({ loaderType: LoaderTypes.UPDATE_ME }));
  } catch (error) {
    yield put(stopLoading({ loaderType: LoaderTypes.UPDATE_ME, error }));
  }
}

export default updateMe;
