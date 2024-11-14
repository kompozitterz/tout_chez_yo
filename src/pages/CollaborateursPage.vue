<template>
  <q-page padding>
    <q-card flat bordered class="q-pa-md">
      <q-toolbar>
        <q-toolbar-title>Liste des Collaborateurs</q-toolbar-title>
        <q-btn
          v-if="isAdmin"
          label="Nouveau collaborateur"
          color="primary"
          @click="showAddCollaboratorModal = true"
          icon="add"
        />
      </q-toolbar>

      <q-list bordered>
        <q-item
          v-for="collaborateur in collaborateurs"
          :key="collaborateur.id"
          clickable
          @click="viewDetails(collaborateur)"
        >
          <q-item-section>
            <q-item-label>{{ collaborateur.prenom }} {{ collaborateur.nom }}</q-item-label>
            <q-item-label caption>{{ collaborateur.profession }}</q-item-label>
            <q-item-label caption>{{ collaborateur.poste }}</q-item-label>
            <q-item-label caption>
              Entrée le : {{ collaborateur.date_entree }}
            </q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-card>

    <q-dialog v-model="showAddCollaboratorModal" persistent>
      <q-card style="min-width: 400px;">
        <q-card-section>
          <div class="text-h6">Ajouter un Collaborateur</div>
        </q-card-section>

        <q-card-section>
          <q-form @submit="addCollaborator">
            <q-input filled v-model="newCollaborator.nom" label="Nom" required />
            <q-input filled v-model="newCollaborator.prenom" label="Prénom" required />
            <q-input filled v-model="newCollaborator.profession" label="Profession" required />
            <q-input
              filled
              v-model="newCollaborator.date_naissance"
              label="Date de naissance"
              type="date"
            />
            <q-input
            filled v-model="newCollaborator.date_entree"
            label="Date d'entrée" type="date" required />
            <q-input filled v-model="newCollaborator.adresse" label="Adresse" />
            <q-input filled v-model="newCollaborator.nationalite" label="Nationalité" />
            <q-input filled v-model="newCollaborator.telephone" label="Téléphone" />
            <q-input filled v-model="newCollaborator.email" label="Email" type="email" />

            <q-btn label="Ajouter" type="submit" color="primary" class="q-mt-md" />
            <q-btn
            flat label="Annuler" color="negative" @click="showAddCollaboratorModal = false" />
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
        {
          id: 1, nom: 'Dupont', prenom: 'Jean', profession: 'Développeur', poste: 'Fullstack', date_entree: '2024-01-15',
        },
        {
          id: 2, nom: 'Curie', prenom: 'Marie', profession: 'Scientifique', poste: 'Responsable RH', date_entree: '2023-06-10',
        },
      ],
      newCollaborator: {
        nom: '',
        prenom: '',
        profession: '',
        date_naissance: '',
        date_entree: '',
        adresse: '',
        nationalite: '',
        telephone: '',
        email: '',
      },
    };
  },
  methods: {
    viewDetails(collaborateur) {
      // eslint-disable-next-line no-console
      console.log('Affichage des détails pour', collaborateur);
    },
    addCollaborator() {
      // eslint-disable-next-line no-console
      console.log('Collaborateur ajouté :', this.newCollaborator);
      this.collaborateurs.push({ ...this.newCollaborator, id: Date.now() });
      this.showAddCollaboratorModal = false;
      this.clearForm();
    },
    clearForm() {
      this.newCollaborator = {
        nom: '',
        prenom: '',
        profession: '',
        date_naissance: '',
        date_entree: '',
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
