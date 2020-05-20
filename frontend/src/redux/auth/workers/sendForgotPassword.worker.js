import {
  put,
} from 'redux-saga/effects';
import ApiService, { httpRequest } from 'src/services/api';
import { sendNotification } from 'src/redux/global/global.actions';
import { startLoading, stopLoading, LoaderTypes } from '../../loaders';

function* sendForgotPassword(action) {
  const { values, notificationMessage } = action.payload;
  try {
    yield put(startLoading({ loaderType: LoaderTypes.FORGOT_PASSWORD }));
    yield httpRequest(ApiService.forgotPassword, values.email, values.mobile);
    yield put(stopLoading({ loaderType: LoaderTypes.FORGOT_PASSWORD }));
    yield put(sendNotification(notificationMessage, 'success'));
  } catch (error) {
    yield put(stopLoading({ loaderType: LoaderTypes.FORGOT_PASSWORD, error }));
  }
}

export default sendForgotPassword;
