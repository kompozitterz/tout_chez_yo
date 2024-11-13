<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title>
          TOUT CHEZ YO
        </q-toolbar-title>

        <!-- Sélecteur de langue avec drapeaux -->
        <q-btn-dropdown
          flat
          dense
          color="primary"
          icon="language"
          label="语言"
        >
          <q-list>
            <q-item clickable v-ripple @click="setLanguage('fr')">
              <q-item-section avatar>
                <q-icon name="flag_fr" />
              </q-item-section>
              <q-item-section>🇨🇵</q-item-section>
            </q-item>
            <q-item clickable v-ripple @click="setLanguage('en')">
              <q-item-section avatar>
                <q-icon name="flag_us" />
              </q-item-section>
              <q-item-section>🇬🇧</q-item-section>
            </q-item>
          </q-list>
        </q-btn-dropdown>

        <!-- Boutons Register et Login -->
        <q-space />
        <q-btn flat dense label="Register" @click="$router.push('/register')" />
        <q-btn flat dense label="Login" @click="$router.push('/login')" />
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

const menuList = [
  {
    title: 'Rédaction de devis',
    icon: 'description',
    link: '/devis',
  },
  {
    title: 'Facturation',
    icon: 'attach_money',
    link: '/facturation',
  },
  {
    title: 'Comptabilité',
    icon: 'calculate',
    link: '/comptabilite',
  },
  {
    title: 'Messagerie instantanée',
    icon: 'chat',
    link: '/messagerie',
  },
  {
    title: 'Timesheets',
    icon: 'watch',
    link: '/timesheets',
  },
  {
    title: 'Inventaire',
    icon: 'inventory',
    link: '/inventaire',
  },
  {
    title: 'Collaborateurs',
    icon: 'groups',
    link: '/collaborateurs',
  },
  {
    title: 'Congés',
    icon: 'beach_access',
    link: '/conges',
  },
];

const leftDrawerOpen = ref(false);

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}

</script>
