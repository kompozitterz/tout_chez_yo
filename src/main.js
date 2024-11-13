import { createApp } from 'vue';
// import { createI18n } from 'vue-i18n';
import { Quasar } from 'quasar';
import { createRouter, createWebHistory } from 'vue-router';
import routes from './router';
import App from './App.vue';
// import messages from './i18n';

// Créez une instance de Vue Router
const router = createRouter({
  history: createWebHistory(),
  routes,
});

const app = createApp(App);
app.use(Quasar, { plugins: {} }); // Configurez Quasar avec des options vides (obligatoire)
// app.use(i18n);
app.use(router);
app.mount('#app');
