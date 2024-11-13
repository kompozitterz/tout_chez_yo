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
        <EssentialLink v-for="link in menuList" :key="link.title" v-bind="link" />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref } from 'vue';
import EssentialLink from 'components/EssentialLink.vue';

defineOptions({
  name: 'MainLayout',
});

const menuList = [
  {
    title: 'Rédaction de devis',
    // caption: 'quasar.dev',
    icon: 'description',
    // link: 'https://quasar.dev',
  },
  {
    title: 'Facturation',
    // caption: 'github.com/quasarframework',
    icon: 'attach_money',
    // link: 'https://github.com/quasarframework',
  },
  {
    title: 'Comptabilité',
    // caption: 'chat.quasar.dev',
    icon: 'calculate',
    // link: 'https://chat.quasar.dev',
  },
  {
    title: 'Messagerie instantanée',
    // caption: 'forum.quasar.dev',
    icon: 'chat',
    // link: 'https://forum.quasar.dev',
  },
  {
    title: 'Timesheets',
    // caption: '@quasarframework',
    icon: 'watch',
    // link: 'https://twitter.quasar.dev',
  },
  {
    title: 'Inventaire',
    // caption: '@QuasarFramework',
    icon: 'inventory',
    // link: 'https://facebook.quasar.dev',
  },
  {
    title: 'Collaborateurs',
    // caption: 'Community Quasar projects',
    icon: 'groups',
    // link: 'https://awesome.quasar.dev',
  },
  {
    title: 'Congés',
    // caption: 'Community Quasar projects',
    icon: 'beach_access',
    // link: 'https://awesome.quasar.dev',
  },
];

const leftDrawerOpen = ref(false);

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}

</script>
