<template>
  <q-page padding>
    <q-card>
      <q-card-section>
        <div class="text-h6">Créer une Facture</div>
      </q-card-section>

      <q-card-section>
        <q-form @submit="submitFacture">
          <!-- Informations de l'entreprise -->
          <div class="text-subtitle1">Informations de l'entreprise</div>
          <q-input filled v-model="entreprise.nom" label="Nom de l'entreprise" />
          <q-input filled v-model="entreprise.adresse" label="Adresse complète" />
          <q-input filled v-model="entreprise.siret" label="Numéro de SIRET ou SIREN" />
          <q-input filled v-model="entreprise.tva" label="Numéro de TVA intracommunautaire" />
          <q-input filled v-model="entreprise.telephone" label="Téléphone" type="tel" />
          <q-input filled v-model="entreprise.email" label="Email" type="email" />
          <q-input filled v-model="entreprise.statut" label="Statut juridique (SARL, SAS...)" />

          <!-- Informations du client -->
          <div class="text-subtitle1">Informations du Client</div>
          <q-input filled v-model="client.nom" label="Nom du client" />
          <q-input filled v-model="client.adresse" label="Adresse complète" />
          <q-input filled v-model="client.tva" label="Numéro de TVA intracommunautaire du client" />
          <q-input filled v-model="client.telephone" label="Téléphone du client" type="tel" />
          <q-input filled v-model="client.email" label="Email du client" type="email" />

          <!-- Numéro de facture -->
          <q-input filled v-model="numeroFacture" label="Numéro de facture (ex: 2024-001)" />

          <!-- Date de la facture -->
          <q-input filled v-model="dateFacture" label="Date de la facture" type="date" />

          <!-- Description des produits/services -->
          <div class="text-subtitle1">Description des Produits/Services</div>
          <q-input filled v-model="produit.description" label="Description" type="textarea" />
          <q-input filled v-model="produit.quantite" label="Quantité" type="number" />
          <q-input
          filled v-model="produit.tarifUnitaire"
          label="Tarif unitaire (€)"
          type="number" />

          <!-- Montants -->
          <q-input filled v-model="montantHT" label="Sous-total HT (€)" type="number" />
          <q-input filled v-model="tauxTVA" label="Taux de TVA (%)" type="number" />
          <q-input filled :value="calculateTTC" label="Montant Total TTC (€)" readonly />

          <!-- Conditions de paiement -->
          <div class="text-subtitle1">Conditions de Paiement</div>
          <q-input filled v-model="modePaiement" label="Mode de paiement" />
          <q-input filled v-model="delaiPaiement" label="Délai de paiement (ex : 30 jours)" />
          <q-input filled v-model="penalites" label="Pénalités de retard" />

          <!-- Mention Facture -->
          <div class="text-subtitle1">Mention</div>
          <q-checkbox v-model="mentionFacture" label="Facture" />

          <!-- Informations légales -->
          <q-input filled v-model="cgv" label="Conditions Générales de Vente" type="textarea" />

          <!-- Mentions additionnelles -->
          <q-input filled v-model="echeance" label="Date d'échéance pour le paiement" type="date" />
          <q-input filled v-model="numeroCommande" label="Numéro de commande ou référence" />
          <q-checkbox v-model="signature" label="Signature/Cachet de l'entreprise" />

          <q-btn label="Créer Facture" color="primary" type="submit" />
        </q-form>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup>
import { ref, computed } from 'vue';

const entreprise = ref({
  nom: '',
  adresse: '',
  siret: '',
  tva: '',
  telephone: '',
  email: '',
  statut: '',
});
const client = ref({
  nom: '',
  adresse: '',
  tva: '',
  telephone: '',
  email: '',
});
const numeroFacture = ref('');
const dateFacture = ref('');
const produit = ref({
  description: '',
  quantite: 0,
  tarifUnitaire: 0,
});
const montantHT = ref(0);
const tauxTVA = ref(20); // Par défaut à 20%
const modePaiement = ref('');
const delaiPaiement = ref('');
const penalites = ref('');
const mentionFacture = ref(true);
const cgv = ref('');
const echeance = ref('');
const numeroCommande = ref('');
const signature = ref(false);

const calculateTTC = computed(() => montantHT.value + ((montantHT.value * tauxTVA.value) / 100));

function submitFacture() {
  // eslint-disable-next-line
  console.log({
    entreprise: entreprise.value,
    client: client.value,
    numeroFacture: numeroFacture.value,
    dateFacture: dateFacture.value,
    produit: produit.value,
    montantHT: montantHT.value,
    tauxTVA: tauxTVA.value,
    montantTTC: calculateTTC.value,
    modePaiement: modePaiement.value,
    delaiPaiement: delaiPaiement.value,
    penalites: penalites.value,
    mentionFacture: mentionFacture.value,
    cgv: cgv.value,
    echeance: echeance.value,
    numeroCommande: numeroCommande.value,
    signature: signature.value,
  });
}
</script>

<style src="../css/quasar.variables.scss" ></style>
<style src="../css/facturation.scss" scoped></style>
