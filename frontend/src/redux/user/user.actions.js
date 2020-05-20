const NAME_SPACE = 'user';

export const actionsType = {
  SET_USER: `${NAME_SPACE}/SET_USER`,
  REMOVE_USER: `${NAME_SPACE}/REMOVE_USER`,
  FETCH_ME: `${NAME_SPACE}/FETCH_ME`,
  UPDATE_ME: `${NAME_SPACE}/UPDATE_ME`,
};

export const setUser = (payload) => ({
  type: actionsType.SET_USER,
  payload,
});

export const removeUser = (payload) => ({
  type: actionsType.REMOVE_USER,
  payload,
});

export const fetchMe = (payload) => ({
  type: actionsType.FETCH_ME,
  payload,
});
export const updateMe = (payload) => ({
  type: actionsType.UPDATE_ME,
  payload,
});
