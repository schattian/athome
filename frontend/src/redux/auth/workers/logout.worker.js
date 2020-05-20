import {
  put,
} from 'redux-saga/effects';
import { removeToken } from 'src/services/userToken';
import { removeUser } from 'src/redux/user/user.actions';
import { onLogoutEnd } from '../auth.actions';

function* logout() {
  yield put(removeUser());
  yield put(onLogoutEnd());
  removeToken();
}

export default logout;
