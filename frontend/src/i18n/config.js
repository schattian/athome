module.exports = {
  defaultLanguage: 'en',
  // All the languages that are available must be defined here
  available: [
    {
      lang: 'es',
      locale: 'es_AR',
      name: 'Castellano',
    },
    {
      lang: 'en',
      locale: 'en_EN',
      name: 'English',
    },
  ],
  // The name of the cookie used to store the user language
  lookupCookie: 'lang',

  // The life of the cookie in minutes
  cookieMinutes: 120,
};
