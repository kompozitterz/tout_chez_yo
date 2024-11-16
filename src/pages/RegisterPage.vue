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

    const register = async () => {
      try {
        const response = await apiClient.post('/register', {
          username: username.value,
          email: email.value,
          password: password.value,
        });

        // Utiliser la réponse pour afficher un message ou effectuer une vérification
        // eslint-disable-next-line no-console
        console.log('Réponse du serveur :', response.data);

        successMessage.value = 'Inscription réussie !';
        errorMessage.value = ''; // Réinitialiser les erreurs

        // Afficher une notification
        Notify.create({
          type: 'positive',
          message: 'Inscription réussie. Redirection vers la page de connexion...',
          timeout: 2000, // Notification pendant 2 secondes
        });

        // Rediriger après la notification
        setTimeout(() => {
          router.push('/login'); // Utiliser l'instance de `router` au lieu de `this`
        }, 2000);
      } catch (error) {
        Notify.create({
          type: 'negative',
          message: 'Erreur lors de l’inscription. Veuillez réessayer.',
        });

        // eslint-disable-next-line no-console
        console.error('Erreur lors de l’inscription :', error.response || error.message);
        successMessage.value = '';
        errorMessage.value = error.response?.data || 'Erreur interne. Veuillez réessayer.';
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
