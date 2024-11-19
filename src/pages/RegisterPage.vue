<template>
  <q-layout>
    <q-page-container>
      <q-page class="flex flex-center center-card">
        <q-card class="q-pa-md" style="width: 400px; max-width: 90%;">
          <q-card-section>
            <div class="text-h6 text-center">Inscription</div>
          </q-card-section>

          <q-card-section>
            <q-form @submit.prevent="register">
              <q-input
                filled
                v-model="username"
                label="Nom d'utilisateur"
                :rules="[val => !!val || 'Le nom d’utilisateur est requis']"
              />
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
                label="S’inscrire"
                type="submit"
                color="primary"
                class="full-width q-mt-md"
              />
            </q-form>
          </q-card-section>

          <q-card-actions align="center">
            <q-btn
              flat
              label="Déjà inscrit ? Se connecter"
              @click="$router.push('/login')"
              color="primary"
            />
          </q-card-actions>
        </q-card>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
import { ref } from 'vue';
import apiClient from 'src/services/apiClient';
import { useRouter } from 'vue-router'; // Importer le routeur
import { Notify } from 'quasar';

export default {
  setup() {
    const username = ref('');
    const email = ref('');
    const password = ref('');
    const successMessage = ref('');
    const errorMessage = ref('');
    const router = useRouter(); // Obtenez une instance du routeur

    const validateForm = () => {
      if (!username.value.trim()) {
        Notify.create({ type: 'warning', message: "Le nom d'utilisateur est requis." });
        return false;
      }
      if (!email.value.trim()) {
        Notify.create({ type: 'warning', message: "L'email est requis." });
        return false;
      }
      if (!password.value.trim()) {
        Notify.create({ type: 'warning', message: 'Le mot de passe est requis.' });
        return false;
      }
      return true;
    };

    const register = async () => {
      if (!validateForm()) return;

      try {
        const response = await apiClient.post('/register', {
          username: username.value,
          email: email.value,
          password: password.value,
        });
        // eslint-disable-next-line no-console
        console.log('Réponse du serveur :', response.data);

        Notify.create({
          type: 'positive',
          message: 'Inscription réussie. Redirection vers la page de connexion...',
          timeout: 2000,
        });

        setTimeout(() => {
          router.push('/login');
        }, 2000);
      } catch (error) {
        const status = error.response?.status;

        if (status === 409) {
          Notify.create({
            type: 'negative',
            message: 'Cet email est déjà utilisé. Veuillez en essayer un autre.',
          });
        } else {
          Notify.create({
            type: 'negative',
            message: 'Erreur lors de l’inscription. Veuillez réessayer.',
          });
        }
        // eslint-disable-next-line no-console
        console.error('Erreur lors de l’inscription :', error.response || error.message);
      }
    };

    return {
      username,
      email,
      password,
      register,
      successMessage,
      errorMessage,
    };
  },
};
</script>

<style src="../css/quasar.variables.scss"></style>
<style src="../css/authentification.scss" scoped></style>
