const NAME_SPACE = 'global';

export const actionsType = {
  INCREMENT: `${NAME_SPACE}/INCREMENT`,
  DECREMENT: `${NAME_SPACE}/DECREMENT`,
  SEND_NOTIFICATION: `${NAME_SPACE}/SEND_NOTIFICATION`,
  CLOSE_NOTIFICATION: `${NAME_SPACE}/CLOSE_NOTIFICATION`,
  REMOVE_NOTIFICATION: `${NAME_SPACE}/REMOVE_NOTIFICATION`,
  SET_NETWORK_ONLINE: `${NAME_SPACE}/SET_NETWORK_ONLINE`,
  SET_WINDOW_SIZE: `${NAME_SPACE}/SET_WINDOW_SIZE`,
};

// eslint-disable-next-line import/prefer-default-export
export const increment = (payload) => ({
  type: actionsType.INCREMENT,
  payload,
});

export const decrement = (payload) => ({
  type: actionsType.DECREMENT,
  payload,
});

// Notifications
export const sendNotification = (message, type, id) => {
  const key = id || new Date().getTime() + Math.random();

  return {
    type: actionsType.SEND_NOTIFICATION,
    notification: {
      message,
      options: {
        variant: type || 'info',
      },
      key,
    },
  };
};

export const closeNotification = (key) => ({
  type: actionsType.CLOSE_NOTIFICATION,
  dismissAll: !key, // dismiss all if no key has been defined
  key,
});

export const removeNotification = (key) => ({
  type: actionsType.REMOVE_NOTIFICATION,
  key,
});

// Notifications End

export const setWindowSize = (width, height) => ({
  type: actionsType.SET_WINDOW_SIZE,
  width,
  height,
});
export const setNetworkOnline = (payload) => ({
  type: actionsType.SET_NETWORK_ONLINE,
  payload,
});
