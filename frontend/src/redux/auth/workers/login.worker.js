import {
  put,
} from 'redux-saga/effects';
import ApiService, { httpRequest } from 'src/services/api';
import Router from 'next/router';
import { setToken } from 'src/services/userToken';
import { setUser } from 'src/redux/user/user.actions';
import { startLoading, stopLoading, LoaderTypes } from '../../loaders';
import { onLoginEnd } from '../auth.actions';

function* signIn(action) {
  const {
    values, nextRoute,
  } = action.payload;
  try {
    yield put(startLoading({ loaderType: LoaderTypes.LOGIN }));
    const response = yield httpRequest(
      ApiService.login,
      values.email.trim().toLowerCase(),
      values.password.trim(),
    );
    const { accessToken, user } = response.data;
    setToken(accessToken);
    yield put(setUser(user));
    yield put(stopLoading({ loaderType: LoaderTypes.LOGIN }));
    yield put(onLoginEnd(accessToken));
    if (nextRoute && nextRoute.length) {
      Router.replace({ pathname: nextRoute });
    }
  } catch (error) {
    yield put(stopLoading({ loaderType: LoaderTypes.LOGIN, error }));
  }
}

export default signIn;
