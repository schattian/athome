import ApiService from 'src/services/api';
import logger from 'src/services/logger';
import { setUser } from 'src/redux/user/user.actions';
import { removeToken } from 'src/services/userToken';

const NAME_SPACE = 'auth';

export const actionsType = {
  SET_TOKEN: `${NAME_SPACE}/SET_TOKEN`,
  ON_CHECK_TOKEN_START: `${NAME_SPACE}/ON_CHECK_TOKEN_START`,
  ON_CHECK_TOKEN_END: `${NAME_SPACE}/ON_CHECK_TOKEN_END`,
  ON_CHECK_TOKEN_FAILED: `${NAME_SPACE}/ON_CHECK_TOKEN_FAILED`,
  ON_CHECK_TOKEN_FAILED_NETWORK_ERROR: `${NAME_SPACE}/ON_CHECK_TOKEN_FAILED_NETWORK_ERROR`,
  LOGOUT: `${NAME_SPACE}/LOGOUT`,
  ON_LOGOUT_END: `${NAME_SPACE}/ON_LOGOUT_END`,
  LOGIN: `${NAME_SPACE}/LOGIN`,
  REGISTER: `${NAME_SPACE}/REGISTER`,
  SEND_VERIFY_ACCOUNT: `${NAME_SPACE}/SEND_VERIFY_ACCOUNT`,
  SEND_FORGOT_PASSWORD: `${NAME_SPACE}/SEND_FORGOT_PASSWORD`,
  ON_LOGIN_END: `${NAME_SPACE}/ON_LOGIN_END`,
  ON_REGISTER_END: `${NAME_SPACE}/ON_REGISTER_END`,
  SET_LOADING: `${NAME_SPACE}/SET_LOADING`,
};

export const onCheckTokenStart = () => ({
  type: actionsType.ON_CHECK_TOKEN_START,
});
export const onCheckTokenEnd = () => ({
  type: actionsType.ON_CHECK_TOKEN_END,
});
export const setLoading = (payload) => ({
  type: actionsType.SET_LOADING,
  payload,
});
export const onCheckTokenFailed = () => ({
  type: actionsType.ON_CHECK_TOKEN_FAILED,
});
export const onCheckTokenFailedNetworkError = () => ({
  type: actionsType.ON_CHECK_TOKEN_FAILED_NETWORK_ERROR,
});

export const onLoginEnd = (token) => ({
  type: actionsType.ON_LOGIN_END,
  token,
});
export const onRegisterEnd = (token) => ({
  type: actionsType.ON_REGISTER_END,
  token,
});
/**
 * @function login
 * @param {object} payload
 * @param {string} payload.email
 * @param {string} payload.password
 * @param {string} payload.nextRoute next route
 *
 */
export const login = (payload) => ({
  type: actionsType.LOGIN,
  payload,
});

/**
 * @function register
 * @param {object} payload
 * @param {string} payload.email
 * @param {string} payload.password
 * @param {string} payload.nextRoute next route
 *
 */
export const register = (payload) => ({
  type: actionsType.REGISTER,
  payload,
});

/**
 * @function logout
 *
 */
export const logout = () => ({
  type: actionsType.LOGOUT,
});
export const onLogoutEnd = () => ({
  type: actionsType.ON_LOGOUT_END,
});

/**
 * @function sendVerifyAccount
 * @param {object} payload
 * @param {string} payload.email
 *
 */
export const sendVerifyAccount = (payload) => ({
  type: actionsType.SEND_VERIFY_ACCOUNT,
  payload,
});

/**
 * @function sendForgotPassword
 * @param {object} payload
 * @param {string} payload.email
 *
 */
export const sendForgotPassword = (payload) => ({
  type: actionsType.SEND_FORGOT_PASSWORD,
  payload,
});


/**
 * @function reAuthenticate
 * This is action create to let as make the action async
 * @param {*} dispatch
 * @param {*} token
 * @param {*} onRemoveToken
 */
export function reAuthenticate(dispatch, token) {
  logger.debug('reAuthenticate start');
  ApiService.reAuthenticate(token).then((response) => {
    logger.debug('Token is valid');
    dispatch(setUser(response.data));
    dispatch(onCheckTokenEnd());
  })
    .catch((error) => {
      if (error.message === 'Network Error') {
        dispatch(onCheckTokenFailedNetworkError());
        logger.debug('reAuthenticate failed, Network Error');
      } else {
        logger.debug('Token is not valid');
        removeToken();
        dispatch(onCheckTokenFailed());
      }
    })
    .finally(() => {
      logger.debug('reAuthenticate end');
    });
}
