<template>
  <q-page padding>
    <q-card flat bordered class="q-pa-md">
      <q-toolbar>
        <q-toolbar-title>Gestion des Congés</q-toolbar-title>
        <!-- Bouton pour ajouter une nouvelle demande de congé -->
        <q-btn
          label="Nouvelle demande de congé"
          color="primary"
          @click="showAddLeaveModal = true"
          icon="add" />
      </q-toolbar>

      <!-- Liste des demandes de congés -->
      <q-table
        :rows="conges"
        :columns="columns"
        row-key="id"
        flat
        bordered
        :rows-per-page-options="[5, 10, 20]">
        <template v-slot:body-cell-status="props">
          <q-chip
            :color="getStatusColor(props.row.status)"
            class="text-white"
            :label="props.row.status"
          />
        </template>
      </q-table>
    </q-card>

    <!-- Modal pour la demande de congé -->
    <q-dialog v-model="showAddLeaveModal" persistent>
      <q-card style="min-width: 400px;">
        <q-card-section>
          <div class="text-h6">Nouvelle demande de congé</div>
        </q-card-section>

        <q-card-section>
          <q-form @submit="addLeave">
            <!-- Type de congé et dates -->
            <q-select
              filled
              v-model="newLeave.type"
              :options="leaveTypes"
              label="Type de congé"
              required
            />
            <q-input filled v-model="newLeave.startDate"
            label="Date de début"
            type="date" required />
            <q-input filled v-model="newLeave.endDate" label="Date de fin" type="date" required />

            <!-- Motif (optionnel) -->
            <q-input filled v-model="newLeave.reason" label="Motif" type="textarea" />

            <!-- Soumettre -->
            <q-btn label="Soumettre" type="submit" color="primary" class="q-mt-md" />
            <q-btn flat label="Annuler" color="negative" @click="showAddLeaveModal = false" />
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
      showAddLeaveModal: false,
      conges: [
        // Exemple de données de congés (à remplacer par des données réelles)
        {
          id: 1, type: 'Congé payé', startDate: '2024-12-01', endDate: '2024-12-10', status: 'en attente',
        },
        {
          id: 2, type: 'Congé maladie', startDate: '2024-11-15', endDate: '2024-11-20', status: 'acceptée',
        },
      ],
      newLeave: {
        type: '',
        startDate: '',
        endDate: '',
        reason: '',
        status: 'en attente', // Par défaut, chaque demande est en attente
      },
      leaveTypes: [
        { label: 'Congé payé', value: 'Congé payé' },
        { label: 'RTT', value: 'RTT' },
        { label: 'Congé sans solde', value: 'Congé sans solde' },
        { label: 'Congé maladie', value: 'Congé maladie' },
        { label: 'Congé maternité/paternité', value: 'Congé maternité/paternité' },
      ],
      columns: [
        {
          name: 'type', label: 'Type de Congé', align: 'left', field: (row) => row.type,
        },
        {
          name: 'startDate', label: 'Date de Début', align: 'left', field: (row) => row.startDate,
        },
        {
          name: 'endDate', label: 'Date de Fin', align: 'left', field: (row) => row.endDate,
        },
        {
          name: 'status', label: 'Statut', align: 'left', field: (row) => row.status,
        },
      ],
    };
  },
  methods: {
    addLeave() {
      // Ajouter la nouvelle demande de congé à la liste (à implémenter avec votre backend)
      this.conges.push({
        ...this.newLeave,
        id: Date.now(),
      });
      this.showAddLeaveModal = false;
      this.clearForm();
    },
    clearForm() {
      this.newLeave = {
        type: '',
        startDate: '',
        endDate: '',
        reason: '',
        status: 'en attente',
      };
    },
    getStatusColor(status) {
      switch (status) {
        case 'acceptée':
          return 'green';
        case 'en attente':
          return 'orange';
        case 'refusée':
          return 'red';
        default:
          return 'grey';
      }
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
<style src="../css/conges.scss" scoped></style>
