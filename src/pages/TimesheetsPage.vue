<template>
  <q-page padding>
    <q-card>
      <q-card-section>
        <div class="text-h6">Feuille de Temps</div>
      </q-card-section>

      <q-card-section>
        <q-form @submit.prevent="submitTimesheet">
          <!-- Informations du collaborateur -->
          <div class="text-subtitle1">Informations du collaborateur</div>
          <q-input filled v-model="collaborateur.nom" label="Nom du collaborateur" readonly />
          <q-input filled v-model="collaborateur.id" label="ID du collaborateur" readonly />

          <!-- Date de travail -->
          <q-input filled v-model="date" label="Date de travail" type="date" readonly />

          <!-- Heures de travail -->
          <div class="text-subtitle1">Heures de travail</div>
          <q-input filled v-model="heureArrivee" label="Heure d'arrivée" readonly />
          <div v-for="(pause, index) in pauses" :key="index" class="q-mb-sm">
            <q-input filled v-model="pause.debut"
            :label="'Début de pause ' + (index + 1)" readonly />
            <q-input filled v-model="pause.fin"
            :label="'Fin de pause ' + (index + 1)" readonly />
          </div>
          <q-input filled v-model="heureDepart" label="Heure de départ" readonly />

          <!-- Boutons pour enregistrer l'heure d'arrivée, départ, et les pauses -->
          <q-btn
            :label="heureArrivee ? 'Enregistrer Heure de Départ' : 'Enregistrer Heure d\'Arrivée'"
            color="primary"
            @click="recordTime('main')"
          />

          <!-- Bouton dédié aux pauses -->
          <q-btn
            :label="pauseActionLabel"
            color="secondary"
            @click="recordTime('pause')"
          />

          <!-- Calcul du temps travaillé -->
          <q-input filled v-model="heuresTravaillees" label="Heures travaillées" readonly />

          <!-- Soumettre la feuille de temps -->
          <q-btn label="Enregistrer Feuille de Temps" color="primary" type="submit" />
        </q-form>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup>
import { ref, computed } from 'vue';

const collaborateur = ref({
  nom: 'John Doe', // Exemple, pourrait être récupéré dynamiquement
  id: '12345',
});
const date = ref(new Date().toISOString().slice(0, 10)); // Date actuelle par défaut
const heureArrivee = ref('');
const heureDepart = ref('');
const pauses = ref([
  { debut: '', fin: '' },
  { debut: '', fin: '' },
  { debut: '', fin: '' },
]);
const heuresTravaillees = ref('0 h'); // Stockage réactif des heures travaillées

// Libellé du bouton de pause qui alterne entre "Début de Pause" et "Fin de Pause"
const pauseActionLabel = computed(() => {
  const currentPause = pauses.value.find((pause) => !pause.debut || !pause.fin);
  return currentPause && currentPause.debut && !currentPause.fin
    ? 'Fin de Pause'
    : 'Début de Pause';
});

// Fonction pour calculer les heures travaillées
function calculateHoursWorked() {
  if (heureArrivee.value && heureDepart.value) {
    const arrivee = new Date(`1970-01-01T${heureArrivee.value}:00`);
    const depart = new Date(`1970-01-01T${heureDepart.value}:00`);
    let totalWorked = (depart - arrivee) / (1000 * 60 * 60); // Temps total entre arrivée et départ

    pauses.value.forEach((pause) => {
      if (pause.debut && pause.fin) {
        const pauseStart = new Date(`1970-01-01T${pause.debut}:00`);
        const pauseEnd = new Date(`1970-01-01T${pause.fin}:00`);
        totalWorked -= (pauseEnd - pauseStart) / (1000 * 60 * 60);
      }
    });
    heuresTravaillees.value = totalWorked > 0 ? `${totalWorked.toFixed(2)} h` : '0 h';
  } else {
    heuresTravaillees.value = '0 h';
  }
}

function recordTime(type) {
  const currentTime = new Date().toTimeString().slice(0, 5); // Format HH:MM

  if (type === 'main') {
    if (!heureArrivee.value) {
      heureArrivee.value = currentTime; // Enregistre l'heure d'arrivée
      // eslint-disable-next-line no-console
      console.log(`Heure d'arrivée enregistrée: ${heureArrivee.value}`);
    } else if (!heureDepart.value) {
      heureDepart.value = currentTime; // Enregistre l'heure de départ
      // eslint-disable-next-line no-console
      console.log(`Heure de départ enregistrée: ${heureDepart.value}`);
      calculateHoursWorked();
    }
  } else if (type === 'pause') {
    const currentPauseIndex = pauses.value.findIndex((pause) => !pause.debut || !pause.fin);

    if (currentPauseIndex !== -1) {
      if (!pauses.value[currentPauseIndex].debut) {
        pauses.value[currentPauseIndex].debut = currentTime; // Enregistre le début de pause
        // eslint-disable-next-line no-console
        console.log(`Début de pause enregistré: ${pauses.value[currentPauseIndex].debut}`);
      } else if (!pauses.value[currentPauseIndex].fin) {
        pauses.value[currentPauseIndex].fin = currentTime;
        // eslint-disable-next-line no-console
        console.log(`Fin de pause enregistré: ${pauses.value[currentPauseIndex].fin}`);
      }

      pauses.value = [...pauses.value];
      // eslint-disable-next-line no-console
      console.log('État actuel des pauses :', pauses.value);
    }
  }
}

function submitTimesheet() {
  // eslint-disable-next-line no-console
  console.log('Soumission de la feuille de temps :', {
    collaborateur: collaborateur.value,
    date: date.value,
    heureArrivee: heureArrivee.value,
    pauses: pauses.value,
    heureDepart: heureDepart.value,
    heuresTravaillees: heuresTravaillees.value,
  });
}
</script>

<style src="../css/quasar.variables.scss" ></style>
<style src="../css/timesheets.scss" scoped></style>
