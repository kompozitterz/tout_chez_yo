<template>
  <q-layout>
    <q-page-container>
      <q-page>
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
    </q-page-container>
  </q-layout>
</template>

<script>
import axios from 'axios';
import { ref } from 'vue';

export default {
  setup() {
    const email = ref('');
    const password = ref('');

    const login = async () => {
      try {
        const response = await axios.post('http://localhost:8080/api/login', {
          email: email.value,
          password: password.value,
        });

        // Stocke le token dans le localStorage ou Vuex
        localStorage.setItem('token', response.data.token);
        // eslint-disable-next-line no-console
        console.log('Connexion réussie !');
        this.$router.push('/dashboard'); // Redirection après connexion
      } catch (error) {
        // eslint-disable-next-line no-console
        console.error('Erreur lors de la connexion:', error);
        // eslint-disable-next-line no-console
        console.log('Échec de la connexion. Vérifiez vos informations.');
      }
    };

    return { email, password, login };
  },
};
</script>

<style src="../css/quasar.variables.scss" ></style>
<style src="../css/authentification.scss" scoped></style>
