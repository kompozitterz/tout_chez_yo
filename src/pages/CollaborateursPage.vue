<template>
  <q-page padding>
    <q-card flat bordered class="q-pa-md">
      <q-toolbar>
        <q-toolbar-title>Liste des Collaborateurs</q-toolbar-title>
        <!-- Bouton "Nouveau Collaborateur" pour les administrateurs -->
        <q-btn
          v-if="isAdmin"
          label="Nouveau collaborateur"
          color="primary"
          @click="showAddCollaboratorModal = true"
          icon="add" />
      </q-toolbar>

      <!-- Liste des collaborateurs -->
      <q-list bordered>
        <q-item v-for="collaborateur in collaborateurs"
        :key="collaborateur.id"
        clickable @click="viewDetails(collaborateur)">
          <q-item-section>
            <q-item-label>{{ collaborateur.nom_complet }}</q-item-label>
            <q-item-label caption>{{ collaborateur.poste }}</q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-card>

    <!-- Modal d'ajout de collaborateur (réservé aux administrateurs) -->
    <q-dialog v-model="showAddCollaboratorModal" persistent>
      <q-card style="min-width: 400px;">
        <q-card-section>
          <div class="text-h6">Ajouter un Collaborateur</div>
        </q-card-section>

        <q-card-section>
          <q-form @submit="addCollaborator">
            <!-- Formulaire d'informations personnelles -->
            <q-input filled v-model="newCollaborator.nom_complet" label="Nom complet" required />
            <q-input
            filled v-model="newCollaborator.date_naissance"
            label="Date de naissance"
            type="date" />
            <q-input filled v-model="newCollaborator.adresse" label="Adresse" />
            <q-input filled v-model="newCollaborator.nationalite" label="Nationalité" />

            <!-- Autres champs requis, exemples pour le contact -->
            <q-input filled v-model="newCollaborator.telephone" label="Téléphone" />
            <q-input filled v-model="newCollaborator.email" label="Email" type="email" />

            <!-- Soumettre -->
            <q-btn label="Ajouter" type="submit" color="primary" class="q-mt-md" />
            <q-btn
            flat label="Annuler" color="negative"
            @click="showAddCollaboratorModal = false" />
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script>
export default {
  data() {
    return {
      isAdmin: true,
      showAddCollaboratorModal: false,
      collaborateurs: [
        // Exemple de données de collaborateurs (à remplacer par des données réelles)
        { id: 1, nom_complet: 'Jean Dupont', poste: 'Développeur Fullstack' },
        { id: 2, nom_complet: 'Marie Curie', poste: 'Responsable RH' },
      ],
      newCollaborator: {
        nom_complet: '',
        date_naissance: '',
        adresse: '',
        nationalite: '',
        telephone: '',
        email: '',
      },
    };
  },
  methods: {
    viewDetails(collaborateur) {
      // Afficher les détails du collaborateur
      // eslint-disable-next-line
      console.log('Affichage des détails pour', collaborateur);
    },
    addCollaborator() {
      // Ajouter le collaborateur (à implémenter avec votre backend)
      // eslint-disable-next-line
      console.log('Collaborateur ajouté :', this.newCollaborator);
      this.showAddCollaboratorModal = false;
      this.clearForm();
    },
    clearForm() {
      this.newCollaborator = {
        nom_complet: '',
        date_naissance: '',
        adresse: '',
        nationalite: '',
        telephone: '',
        email: '',
      };
    },
  },
};
</script>

<style scoped>
.q-page {
  max-width: 800px;
  margin: auto;
}
</style>

<style src="../css/quasar.variables.scss" ></style>
<style src="../css/collaborateurs.scss" scoped></style>
