import { createApp } from 'vue';
import { createI18n } from 'vue-i18n';
import { Quasar } from 'quasar';
import { createRouter, createWebHistory } from 'vue-router';
import routes from './router';
import App from './App.vue';
import messages from './i18n';

// Créez une instance de i18n
const i18n = createI18n({
  locale: 'fr',
  fallbackLocale: 'en',
  messages,
});

// Créez une instance de Vue Router
const router = createRouter({
  history: createWebHistory(),
  routes,
});

const app = createApp(App);
app.use(Quasar);
app.use(i18n);
app.use(router); // Ajoutez le router à l’application
app.mount('#app');
