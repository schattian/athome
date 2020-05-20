const express = require('express');
const next = require('next');
const nextI18NextMiddleware = require('next-i18next/middleware').default;
const nextI18next = require('./src/i18n');

const port = process.env.PORT || 3000;
const isDev = process.env.NODE_ENV !== 'production'
  && (process.env.CONFIG_ENV === 'development'
    || process.env.CONFIG_ENV === 'local');

const app = next({ dev: isDev });
const handle = app.getRequestHandler();
(async () => {
  await app.prepare();
  const server = express();

  server.use(nextI18NextMiddleware(nextI18next));
  server.get('*', (req, res) => handle(req, res));

  await server.listen(port);
  console.log(`> Ready on http://localhost:${port}`); // eslint-disable-line no-console
})();
