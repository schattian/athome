const NextI18Next = require('next-i18next').default;
const config = require('./config');

module.exports = new NextI18Next({
  localePath: typeof window === 'undefined' ? 'public/static/locales' : 'static/locales',
  defaultLanguage: config.defaultLanguage,
  otherLanguages: config.available.map(({ lang }) => lang),
  detection: {
    lookupCookie: config.lookupCookie,
    cookieMinutes: config.cookieMinutes,
  },
  browserLanguageDetection: true,
  serverLanguageDetection: true,
});

/* Optionally, export class methods as named exports */
