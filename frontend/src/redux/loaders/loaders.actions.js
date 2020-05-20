
const NAMESPACE = 'LOADERS';

export const actionsType = {
  START_LOADING: `${NAMESPACE}/START_LOADING`,
  STOP_LOADING: `${NAMESPACE}/STOP_LOADING`,
};

export const startLoading = (payload) => ({ type: actionsType.START_LOADING, payload });
export const stopLoading = (payload) => ({ type: actionsType.STOP_LOADING, payload });
