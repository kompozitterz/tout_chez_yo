<template>
  <q-page class="flex flex-center">
    <q-card class="q-pa-md" style="width: 400px; max-width: 90%;">
      <q-card-section>
        <div class="text-h6 text-center">Connexion</div>
      </q-card-section>

      <q-card-section>
        <q-form @submit.prevent="login">
          <q-input
            filled
            v-model="email"
            label="Email"
            type="email"
            :rules="[val => !!val || 'L’email est requis']"
          />
          <q-input
            filled
            v-model="password"
            label="Mot de passe"
            type="password"
            :rules="[val => !!val || 'Le mot de passe est requis']"
          />
          <q-btn
            label="Se connecter"
            type="submit"
            color="primary"
            class="full-width q-mt-md"
          />
        </q-form>
      </q-card-section>

      <q-card-actions align="center">
        <q-btn
        flat
        label="Pas encore inscrit ? Créer un compte"
        @click="$router.push('/register')"
        color="primary" />
      </q-card-actions>
    </q-card>
  </q-page>
</template>

<script>
import { ref } from 'vue';

export default {
  setup() {
    const email = ref('');
    const password = ref('');

    const login = () => {
      // eslint-disable-next-line no-console
      console.log('Tentative de connexion avec', email.value, password.value);
      // Ajoutez ici votre logique pour appeler l’API de connexion
    };

    return { email, password, login };
  },
};
</script>

<style src="../css/quasar.variables.scss" ></style>
<style src="../css/authentification.scss" scoped></style>
