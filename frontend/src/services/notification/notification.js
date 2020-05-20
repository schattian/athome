import { sendNotification as sendNotificationAction } from 'src/redux/global/global.actions';
import logger from 'src/services/logger';

let dispatch;

export const setDispatch = (dispatchRef) => {
  dispatch = dispatchRef;
};

export const sendNotification = (message, type, id) => {
  if (dispatch) {
    dispatch(sendNotificationAction(message, type, id));
  } else {
    logger.error('src/services/notification/notification.js missing dispatch');
  }
};
