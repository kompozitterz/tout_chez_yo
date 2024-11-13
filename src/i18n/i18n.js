import { createI18n } from 'vue-i18n';
import en from './en';
import fr from './fr';

const messages = {
  en,
  fr,
};

const i18n = createI18n({
  locale: 'fr', // Langue par défaut
  fallbackLocale: 'en',
  messages,
});

export default i18n;
