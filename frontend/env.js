const common = {
  userTokenKey: 'someUniqueTokenKey',
};

const dev = {
  ...common,
  serverUrl: 'http://localhost:3030',
};
const qa = {
  ...common,
  serverUrl: 'http://localhost:3030',
};
const prod = {
  ...common,
  serverUrl: 'http://localhost:3030',
};
const local = {
  ...common,
  serverUrl: 'http://localhost:3030',
};

const configByEnv = {
  development: dev,
  production: prod,
  local,
  qa,
};

const config = function config(env) {
  return configByEnv[env];
};

module.exports = config;
