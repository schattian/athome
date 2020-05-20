import {
  put,
} from 'redux-saga/effects';
import ApiService, { httpRequest } from 'src/services/api';
import { sendNotification } from 'src/redux/global/global.actions';
import { startLoading, stopLoading, LoaderTypes } from '../../loaders';

function* sendVerifyAccount(action) {
  const { values, notificationMessage } = action.payload;
  try {
    yield put(startLoading({ loaderType: LoaderTypes.VERIFY_ACCOUNT }));
    yield httpRequest(ApiService.resendVerifyAccount, values.email, values.mobile);
    yield put(stopLoading({ loaderType: LoaderTypes.VERIFY_ACCOUNT }));
    yield put(sendNotification(notificationMessage, 'success'));
  } catch (error) {
    yield put(stopLoading({ loaderType: LoaderTypes.VERIFY_ACCOUNT, error }));
  }
}

export default sendVerifyAccount;
