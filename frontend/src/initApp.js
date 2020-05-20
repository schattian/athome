import {
  setApiInstance,
  setDefaultIdKey,
  setDefaultUpdateMethod,
  setDispatch,
  setErrorHandler,
} from 'net-provider';
import ApiService from 'src/services/api';
import logger from 'src/services/logger';
import { setDispatch as setNotificationDispatch } from 'src/services/notification/notification';

export default (props) => {
  setDefaultUpdateMethod('patch');
  setDefaultIdKey('_id');
  setErrorHandler((err) => logger.error(err));
  setDispatch(props.store.dispatch);
  setApiInstance(ApiService.getAxios());
  setNotificationDispatch(props.store.dispatch);
};
