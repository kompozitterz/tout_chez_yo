import MainLayout from 'layouts/MainLayout.vue';
import IndexPage from 'pages/IndexPage.vue';
import InventairePage from 'pages/InventairePage.vue';
import ComptabilitePage from 'pages/ComptabilitePage.vue';
import FacturationPage from 'pages/FacturationPage.vue';
import DevisPage from 'pages/DevisPage.vue';
import MessageriePage from 'pages/MessageriePage.vue';
import TimesheetsPage from 'pages/TimesheetsPage.vue';
import CollaborateursPage from 'pages/CollaborateursPage.vue';
import CongesPage from 'pages/CongesPage.vue';
import ErrorNotFound from 'pages/ErrorNotFound.vue';
import LoginPage from '../pages/LoginPage.vue';
import RegisterPage from '../pages/RegisterPage.vue';

const routes = [
  // {
  //   path: '/login',
  //   name: 'Login',
  //   component: LoginPage,
  //   meta: { requiresAuth: false },
  // },
  // {
  //   path: '/register',
  //   name: 'Register',
  //   component: RegisterPage,
  //   meta: { requiresAuth: false },
  // },

  {
    path: '/',
    component: MainLayout,
    children: [
      { path: '', component: IndexPage },
      { path: 'inventaire', component: InventairePage },
      { path: 'comptabilite', component: ComptabilitePage },
      { path: 'facturation', component: FacturationPage },
      { path: 'devis', component: DevisPage },
      { path: 'messagerie', component: MessageriePage },
      { path: 'timesheets', component: TimesheetsPage },
      { path: 'collaborateurs', component: CollaborateursPage },
      { path: 'conges', component: CongesPage },
      { path: '/login', name: 'Login', component: LoginPage },
      { path: '/register', name: 'Register', component: RegisterPage },
    ],
  },

  {
    path: '/:catchAll(.*)*',
    name: 'ErrorNotFound',
    component: ErrorNotFound,
  },
];

export default routes;
