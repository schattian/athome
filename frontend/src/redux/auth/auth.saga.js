import {
  takeEvery,
} from 'redux-saga/effects';
import { actionsType } from './auth.actions';
import login from './workers/login.worker';
import logout from './workers/logout.worker';
import register from './workers/register.worker';
import sendVerifyAccount from './workers/sendVerifyAccount.worker';
import sendForgotPassword from './workers/sendForgotPassword.worker';

export default function* authWatcher() {
  yield takeEvery(actionsType.LOGIN, login);
  yield takeEvery(actionsType.LOGOUT, logout);
  yield takeEvery(actionsType.REGISTER, register);
  yield takeEvery(actionsType.SEND_VERIFY_ACCOUNT, sendVerifyAccount);
  yield takeEvery(actionsType.SEND_FORGOT_PASSWORD, sendForgotPassword);
}
