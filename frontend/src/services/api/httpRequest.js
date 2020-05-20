import { call } from 'redux-saga/effects';
import getLocalMessages from 'src/i18n/serverMessages';
import logger from '../logger';

const getErrorMessage = (error, errorText, defaultErrorMessage) => {
  if (errorText) return errorText;
  if (error && typeof error === 'string') return error;
  if (error
    && error.response
    && error.response.data
    && error.response.data.message) return error.response.data.message;
  if (error && error.data && error.data.message) return error.data.message;
  if (error && error.message) return error.message;
  if (defaultErrorMessage) return defaultErrorMessage;
  return 'oops something went wrong';
};

export default function* httpRequest(...params) {
  try {
    const res = yield call(...params);
    return res;
  } catch (error) {
    logger.error('Network request failed', error);
    const errorMessage = getLocalMessages(getErrorMessage(error));
    throw new Error(errorMessage);
  }
}
