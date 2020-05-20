const NAME_SPACE = 'example';

export const actionsType = {
  FETCH_FAKE_DATA: `${NAME_SPACE}/FETCH_FAKE_DATA`,
  SET_FAKE_DATA: `${NAME_SPACE}/SET_FAKE_DATA`,
};

export const fetchFakeData = (payload) => ({
  type: actionsType.FETCH_FAKE_DATA,
  payload,
});

export const setFakeData = (payload) => ({
  type: actionsType.SET_FAKE_DATA,
  payload,
});
