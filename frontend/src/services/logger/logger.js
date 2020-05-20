/* eslint-disable no-console */
// import { notification } from 'antd';
const isProd = process.env.NODE_ENV === 'production';

const logger = {
  error: (message, data) => {
    if (isProd) {
      // TODO: is a good place to send error to some logger service
    } else {
      console.error(message, data);
    //   notification.open({
    //     message: 'Dev Mode Error Notification',
    //     description: message + '\n' + (data ? JSON.stringify(data) : '')
    //   });
    }
  },
  info: (message) => {
    if (isProd) {
      // TODO: is a good place to send error to some logger service
    } else {
      console.info(message);
    }
  },
  warn: (message, data) => {
    if (isProd) {
      // TODO: is a good place to send error to some logger service
    } else {
      console.warn(message, data);
    //   notification.open({
    //     message: 'Dev Mode Warning Notification',
    //     description: message + '\n' + (data ? JSON.stringify(data) : '')
    //   });
    }
  },
  debug: (message, data) => {
    if (!isProd) {
      console.debug(`#debug -${message}`, data || '');
    }
  },
};

export default logger;
