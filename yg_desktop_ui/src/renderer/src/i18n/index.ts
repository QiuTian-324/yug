import { createI18n } from 'vue-i18n';
import en from './locales/en.json';
import zh from './locales/zh.json';

const messages = {
  en: en,
  zh: zh,
};

const i18n = createI18n({
  legacy: false,
  locale: 'zh', // 默认语言
  // locale: 'en', // 默认语言，
  // fallbackLocale: 'en', // 回退语言
  globalInjection: true,
  messages,
});

export default i18n;