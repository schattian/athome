import {
  put,
} from 'redux-saga/effects';
import ApiService, { httpRequest } from 'src/services/api';
import Router from 'next/router';
import { setToken } from 'src/services/userToken';
import { setUser } from 'src/redux/user/user.actions';
import { onRegisterEnd } from '../auth.actions';
import { startLoading, stopLoading, LoaderTypes } from '../../loaders';


function* register(action) {
  const {
    values,
    nextRoute,
  } = action.payload;
  const {
    email, password, firstName, lastName,
  } = values;
  try {
    yield put(startLoading({ loaderType: LoaderTypes.REGISTER }));
    let response = yield httpRequest(
      ApiService.createUser,
      email.trim().toLowerCase(),
      password.trim(),
      firstName,
      lastName,
    );
    if (
      (response.data.verifiedRequired && !response.data.isVerified)
      || (
        response.data.user
        && response.data.user.verifiedRequired
        && !response.data.user.isVerified
      )
    ) {
      Router.replace({ pathname: '/verify-account', query: { email } });
      return;
    }

    if (!response.data.accessToken) {
      // When server is not auto login after sign up
      response = yield httpRequest(ApiService.login, email, password);
    }
    const { accessToken, user } = response.data;
    setToken(accessToken);
    yield put(setUser(user));
    yield put(stopLoading({ loaderType: LoaderTypes.LOGIN }));
    yield put(onRegisterEnd(accessToken));

    yield put(stopLoading({ loaderType: LoaderTypes.REGISTER }));
    if (nextRoute && nextRoute.length) {
      Router.replace({ pathname: nextRoute });
    }
  } catch (error) {
    yield put(stopLoading({ loaderType: LoaderTypes.REGISTER, error }));
  }
}

export default register;
