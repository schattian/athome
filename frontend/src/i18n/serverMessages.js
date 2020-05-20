import kebabCase from 'lodash/kebabCase';
import { i18n } from './index';

// Add new keys in kebabCase https://lodash.com/docs/4.17.15#kebabCase
const locales = {
  es: {
    'user-not-found': 'Usuario no encontrado',
  },
};

function getLocalMessages(message = '') {
  const lang = i18n.language;
  const localMessage = lang && locales[lang] ? locales[lang][kebabCase(message)] : null;
  return localMessage || message;
}

export default getLocalMessages;
