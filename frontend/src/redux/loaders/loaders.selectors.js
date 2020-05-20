import { LoaderStatus } from './loaderTypesEnum';

export const getLoadersState = (state) => state.loaders;

export const getIsLoadingByType = (state, loaderType) => {
  const isLoading = getLoadersState(state)[loaderType].status === LoaderStatus.LOADING;
  return isLoading;
};

export const getIsEndByType = (state, loaderType) => {
  const isLoading = getLoadersState(state)[loaderType].status === LoaderStatus.END;
  return isLoading;
};

export const getStatusByType = (state, loaderType) => getLoadersState(state)[loaderType].status;

export const getErrorByType = (state, loaderType) => getLoadersState(state)[loaderType].error;
