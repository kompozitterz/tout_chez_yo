<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <!-- Bouton Menu -->
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <!-- Bouton Accueil -->
        <q-btn
          flat
          dense
          round
          icon="home"
          aria-label="Accueil"
          @click="$router.push('/')"
        />

        <!-- Logo -->
        <img src="/public/Tout_chez_yo.svg" alt="Logo" style="height: 40px; margin-left: 8px;" />

        <!-- Titre -->
        <q-toolbar-title>
          TOUT CHEZ YO
        </q-toolbar-title>

        <!-- Voyant de connexion -->
        <div class="row items-center">
          <q-icon
            v-if="isConnected"
            name="check_circle"
            color="green"
            size="24px"
            title="Connecté"
            class="q-mr-md"
          />
          <q-icon
            v-else
            name="highlight_off"
            color="red"
            size="24px"
            title="Déconnecté"
            class="q-mr-md"
          />
          <span v-if="isConnected">Connecté</span>
          <span v-else>Déconnecté</span>
        </div>

        <!-- Icône Notifications avec menu déroulant -->
        <q-btn-dropdown
          flat
          dense
          round
          icon="notifications"
          color="primary"
        >
          <q-badge
            v-if="unreadNotifications > 0"
            color="red"
            floating
            style="top: -10px; right: -10px;"
          >
            {{ unreadNotifications }}
          </q-badge>

          <!-- Liste des notifications -->
          <q-list separator>
            <q-item v-if="notifications.length === 0">
              <q-item-section>
                <q-icon name="info" color="grey" />
                <span>Aucune notification</span>
              </q-item-section>
            </q-item>
            <q-item
              v-for="(notification, index) in notifications"
              :key="index"
              clickable
              v-ripple
              @click="markAsRead(index)"
            >
              <q-item-section avatar>
                <q-icon :name="notification.icon || 'notification_important'" />
              </q-item-section>
              <q-item-section>
                <div class="text-bold">{{ notification.title }}</div>
                <div class="text-grey text-caption">{{ notification.message }}</div>
              </q-item-section>
            </q-item>
          </q-list>
        </q-btn-dropdown>

        <!-- Sélecteur de langue -->
        <q-btn-dropdown
          flat
          dense
          color="primary"
          icon="language"
          label="Langue"
        >
          <q-list>
            <q-item clickable v-ripple @click="setLanguage('fr')">
              <q-item-section avatar>
                <q-icon name="flag_fr" />
              </q-item-section>
              <q-item-section>🇨🇵 Français</q-item-section>
            </q-item>
            <q-item clickable v-ripple @click="setLanguage('en')">
              <q-item-section avatar>
                <q-icon name="flag_us" />
              </q-item-section>
              <q-item-section>🇬🇧 Anglais</q-item-section>
            </q-item>
          </q-list>
        </q-btn-dropdown>

        <!-- Boutons Register, Login, Logout -->
        <q-space />
        <q-btn flat dense label="Register" v-if="!isConnected" @click="$router.push('/register')" />
        <q-btn flat dense label="Login" v-if="!isConnected" @click="$router.push('/login')" />
        <q-btn flat dense label="Logout" v-if="isConnected" @click="logout" />
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <q-list>
        <q-item-label header>Menu</q-item-label>
        <q-item
          v-for="link in menuList"
          :key="link.title"
          clickable
          v-ripple
          :to="link.link"
        >
          <q-item-section avatar>
            <q-icon :name="link.icon" />
          </q-item-section>
          <q-item-section>{{ link.title }}</q-item-section>
        </q-item>
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref } from 'vue';

defineOptions({
  name: 'MainLayout',
});

// Menu de navigation
const menuList = [
  { title: 'Rédaction de devis', icon: 'description', link: '/devis' },
  { title: 'Facturation', icon: 'attach_money', link: '/facturation' },
  { title: 'Comptabilité', icon: 'calculate', link: '/comptabilite' },
  { title: 'Messagerie instantanée', icon: 'chat', link: '/messagerie' },
  { title: 'Timesheets', icon: 'watch', link: '/timesheets' },
  { title: 'Inventaire', icon: 'inventory', link: '/inventaire' },
  { title: 'Collaborateurs', icon: 'groups', link: '/collaborateurs' },
  { title: 'Congés', icon: 'beach_access', link: '/conges' },
];

// État de l'application
const leftDrawerOpen = ref(false);
const isConnected = ref(false); // État de connexion
const unreadNotifications = ref(0); // Compteur de notifications non lues
const notifications = ref([]); // Liste des notifications
// const audio = new Audio('/public/notification.mp3');

// Ajouter une nouvelle notification
function addNotification(title, message, icon) {
  notifications.value.push({ title, message, icon });
  unreadNotifications.value++;
  // audio.play();
}

// Marquer une notification comme lue
function markAsRead(index) {
  notifications.value.splice(index, 1);
  if (unreadNotifications.value > 0) unreadNotifications.value--;
}

// Changer la langue (simulé)
function setLanguage(lang) {
  // eslint-disable-next-line no-console
  console.log(`Langue changée en : ${lang}`);
}

// Simuler une notification
function simulateNotification() {
  setTimeout(() => {
    addNotification(
      'Nouvelle notification',
      'Vous avez reçu un message !',
      'mail',
    );
    simulateNotification();
  }, Math.random() * 10000 + 5000); // Notifications toutes les 5 à 15 secondes
}

// Vérifier la connexion
function checkConnection() {
  const token = localStorage.getItem('token');
  isConnected.value = !!token;
}

// Déconnexion
function logout() {
  localStorage.removeItem('token');
  isConnected.value = false;
  window.location.reload();
}

// Initialisation
checkConnection();
simulateNotification(); // Démarrer la simulation de notifications

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}
</script>
