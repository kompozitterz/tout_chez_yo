<template>
  <q-layout>
    <q-page-container>
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
import { useRouter } from 'vue-router'; // Importer le routeur Vue
import { Notify } from 'quasar'; // Importer Notify pour les notifications

export default {
  setup() {
    const email = ref('');
    const password = ref('');
    const router = useRouter();

    const login = async () => {
      try {
        // Log des données envoyées au backend
        // eslint-disable-next-line no-console
        console.log('Données envoyées au backend :', {
          email: email.value,
          password: password.value,
        });

        const response = await apiClient.post('/login', {
          email: email.value,
          password: password.value,
        });
        // eslint-disable-next-line no-console
        console.log('Réponse complète du backend :', response);
        // eslint-disable-next-line no-console
        console.log('Réponse du backend :', response.data);

        if (!response.data.token) {
          throw new Error('Token non reçu !');
        }

        // Stocker le token dans localStorage
        localStorage.setItem('token', response.data.token);

        // Notification de succès
        Notify.create({
          type: 'positive',
          message: "Connexion réussie ! Redirection vers la page d'accueil...",
          timeout: 2000, // Notification visible pendant 2 secondes
        });

        // Redirection vers la page d'accueil
        setTimeout(() => {
          router.push('/'); // Redirige vers l'accueil
        }, 2000);
      } catch (error) {
        // Notification d'échec
        Notify.create({
          type: 'negative',
          message:
            error.response?.data || 'Échec de la connexion. Vérifiez vos informations.',
        });
        // eslint-disable-next-line no-console
        console.error('Erreur lors de la connexion:', error);
      }
    };

    return {
      email,
      password,
      login,
    };
  },
};
</script>

<style src="../css/quasar.variables.scss"></style>
<style src="../css/authentification.scss" scoped></style>
