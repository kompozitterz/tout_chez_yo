<template>
  <q-page class="flex flex-center q-gutter-md">
    <div class="row q-gutter-md wrap justify-center">
      <router-link to="/devis" class="no-decoration">
        <q-badge class="column items-center" align="center">
          <q-icon name="description" size="lg" />
          <div>Rédaction de devis</div>
        </q-badge>
      </router-link>

      <router-link to="/facturation" class="no-decoration">
        <q-badge class="column items-center" align="center">
          <q-icon name="attach_money" size="lg" />
          <div>Facturation</div>
        </q-badge>
      </router-link>

      <router-link to="/comptabilite" class="no-decoration">
        <q-badge class="column items-center" align="center">
          <q-icon name="calculate" size="lg" />
          <div>Comptabilité</div>
        </q-badge>
      </router-link>

      <router-link to="/messagerie" class="no-decoration">
        <q-badge class="column items-center" align="center">
          <q-icon name="chat" size="lg" />
          <div>Messagerie instantanée</div>
        </q-badge>
      </router-link>

      <router-link to="/timesheets" class="no-decoration">
        <q-badge class="column items-center" align="center">
          <q-icon name="watch" size="lg" />
          <div>Timesheets</div>
        </q-badge>
      </router-link>

      <router-link to="/inventaire" class="no-decoration">
        <q-badge class="column items-center" align="center">
          <q-icon name="inventory" size="lg" />
          <div>Inventaire</div>
        </q-badge>
      </router-link>

      <router-link to="/collaborateurs" class="no-decoration">
        <q-badge class="column items-center" align="center">
          <q-icon name="groups" size="lg" />
          <div>Collaborateurs</div>
        </q-badge>
      </router-link>

      <router-link to="/conges" class="no-decoration">
        <q-badge class="column items-center" align="center">
          <q-icon name="beach_access" size="lg" />
          <div>Congés</div>
        </q-badge>
      </router-link>
    </div>
  </q-page>
</template>

<script setup>
defineOptions({
  name: 'IndexPage',
});
</script>

<style>
.no-decoration {
  text-decoration: none;
}
</style>
