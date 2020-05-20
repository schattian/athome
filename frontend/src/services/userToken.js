
import store from 'store';
import ApiService from './api';

const TOKEN_KEY = process.env.userTokenKey;

export function getToken() {
  const token = store.get(TOKEN_KEY);
  if (token) {
    const axiosInstance = ApiService.getAxios();
    axiosInstance.defaults.headers.common.Authorization = token;
  }
  return token;
}
export function setToken(token) {
  const axiosInstance = ApiService.getAxios();
  axiosInstance.defaults.headers.common.Authorization = token;
  return store.set(TOKEN_KEY, token);
}
export function removeToken() {
  const axiosInstance = ApiService.getAxios();
  delete axiosInstance.defaults.headers.common.Authorization;
  return store.remove(TOKEN_KEY);
}
