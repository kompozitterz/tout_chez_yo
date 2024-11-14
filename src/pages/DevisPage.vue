<template>
  <q-page padding>
    <q-card>
      <q-card-section>
        <div class="text-h6">Créer un Devis</div>
      </q-card-section>

      <q-card-section>
        <q-form @submit="submitDevis">
          <!-- Informations de l'entreprise -->
          <div class="text-subtitle1">Informations de l'entreprise</div>
          <q-input
          filled v-model="entreprise.nom"
          label="Nom de l'entreprise"
          class="input"/>

          <q-input
          filled v-model="entreprise.adresse"
          label="Adresse"
          class="input"/>

          <q-input
          filled v-model="entreprise.telephone"
          label="Téléphone"
          type="tel"
          class="input"/>

          <q-input
          filled v-model="entreprise.email"
          label="Email"
          type="email"
          class="input"/>

          <!-- Informations du client -->
          <div class="text-subtitle1">Informations du Client</div>
          <q-input
          filled v-model="client.nom"
          label="Nom du client"
          class="input"/>

          <q-input
          filled v-model="client.adresse"
          label="Adresse"
          class="input"/>

          <q-input
          filled v-model="client.telephone"
          label="Téléphone"
          type="tel"
          class="input"/>

          <q-input
          filled v-model="client.email"
          label="Email"
          type="email"
          class="input"/>

          <!-- Date et numéro du devis -->
          <q-input
          filled v-model="date"
          label="Date d'émission"
          type="date"
          class="input"/>

          <q-input
          filled v-model="numero"
          label="Numéro de devis"
          class="input"/>

          <!-- Description des services ou produits -->
          <div class="text-subtitle1">Détails des services/produits</div>
          <q-input
          filled v-model="produit.description"
          label="Description"
          type="textarea"
          class="input"/>

          <q-input
          filled v-model="produit.quantite"
          label="Quantité"
          type="number"
          class="input"/>

          <q-input
          filled v-model="produit.tarifUnitaire"
          label="Tarif unitaire (€)"
          type="number"
          class="input"/>

          <!-- Montants HT, TVA, TTC -->
          <q-input
          filled v-model="montantHT"
          label="Montant Total HT (€)"
          type="number"
          class="input"/>

          <q-input filled v-model="tauxTVA"
          label="Taux de TVA (%)"
          type="number"
          class="input"/>

          <q-input filled :value="calculateTTC"
          label="Montant Total TTC (€)"
          readonly
          class="input"/>

          <!-- Délai et conditions de livraison -->
          <q-input
          filled v-model="delaiLivraison"
          label="Délai de livraison"
          class="input"/>

          <!-- Conditions de paiement -->
          <q-input
          filled v-model="conditionsPaiement"
          label="Conditions de paiement"
          type="textarea"
          class="input"/>

          <!-- Validité du devis -->
          <q-input
          filled v-model="validite"
          label="Validité du devis (en jours)"
          type="number"
          class="input"/>

          <!-- Conditions générales de vente -->
          <q-input
          filled
          v-model="conditionsGenerales"
          label="Conditions générales de vente"
          type="textarea"
          class="input"/>

          <!-- Signature et accord -->
          <div class="text-subtitle1">Signature</div>
          <q-checkbox v-model="accord" label="Bon pour accord" />

          <q-btn label="Créer Devis" color="primary" type="submit" />
        </q-form>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup>
import { ref, computed } from 'vue';

const entreprise = ref({
  nom: '', adresse: '', telephone: '', email: '',
});
const client = ref({
  nom: '', adresse: '', telephone: '', email: '',
});
const date = ref('');
const numero = ref('');
const produit = ref({ description: '', quantite: '', tarifUnitaire: 0 });
const montantHT = ref(0);
const tauxTVA = ref(20); // Par défaut à 20%
const delaiLivraison = ref('');
const conditionsPaiement = ref('');
const validite = ref(30); // Par défaut à 30 jours
const conditionsGenerales = ref('');
const accord = ref(false);

const calculateTTC = computed(() => montantHT.value + ((montantHT.value * tauxTVA.value) / 100));

function submitDevis() {
  // eslint-disable-next-line
  console.log({
    entreprise: entreprise.value,
    client: client.value,
    date: date.value,
    numero: numero.value,
    produit: produit.value,
    montantHT: montantHT.value,
    tauxTVA: tauxTVA.value,
    montantTTC: calculateTTC.value,
    delaiLivraison: delaiLivraison.value,
    conditionsPaiement: conditionsPaiement.value,
    validite: validite.value,
    conditionsGenerales: conditionsGenerales.value,
    accord: accord.value,
  });
}
</script>

<style src="../css/quasar.variables.scss" ></style>
<style src="../css/devis.scss" scoped></style>
