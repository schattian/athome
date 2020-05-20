import { all, call } from 'redux-saga/effects';
import { crudReduxSaga } from 'net-provider';
import example from 'src/example/redux/example.saga';


import global from './global/global.saga';
import auth from './auth/auth.saga';
import user from './user/user.saga';

// Add your saga here
const rootSaga = function* rootSaga() {
  yield all([
    call(crudReduxSaga, 'crudReduxSaga'),
    call(example, 'exampleWatcher'),
    call(global, 'global'),
    call(auth, 'authWatcher'),
    call(user, 'userWatcher'),
  ]);
};

export default rootSaga;
