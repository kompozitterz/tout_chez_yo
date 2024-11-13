<template>
  <q-page padding>
    <q-card>
      <q-card-section>
        <div class="text-h6">Messagerie instantanée</div>
      </q-card-section>

      <q-card-section>
        <div class="message-container">
          <q-input
            filled
            v-model="message"
            label="Écrivez votre message"
            type="textarea"
            autogrow
          />
          <q-btn icon="send" color="primary" @click="sendMessage" />
        </div>

        <div class="record-container">
          <q-btn icon="mic" color="secondary" @click="toggleAudioRecording">
            <q-tooltip>Enregistrer Audio</q-tooltip>
          </q-btn>
          <q-btn icon="videocam" color="secondary" @click="toggleVideoRecording">
            <q-tooltip>Enregistrer Vidéo</q-tooltip>
          </q-btn>
        </div>

        <!-- Affichage des messages -->
        <div class="messages-display">
          <div v-for="(msg, index) in messages" :key="index" class="message-item">
            <p>{{ msg.text }}</p>
            <audio v-if="msg.audio" :src="msg.audio" controls></audio>
            <video v-if="msg.video" :src="msg.video" controls width="200"></video>
          </div>
        </div>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup>
import { ref } from 'vue';

const message = ref('');
const messages = ref([]);
let audioRecorder; let
  videoRecorder;
const audioChunks = [];
const videoChunks = [];

// Envoi d'un message texte
function sendMessage() {
  if (message.value) {
    messages.value.push({ text: message.value });
    message.value = '';
  }
}

// Fonction d'enregistrement audio
async function toggleAudioRecording() {
  if (audioRecorder && audioRecorder.state === 'recording') {
    audioRecorder.stop();
  } else {
    const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
    audioRecorder = new MediaRecorder(stream);

    audioRecorder.ondataavailable = (e) => audioChunks.push(e.data);
    audioRecorder.onstop = () => {
      const audioBlob = new Blob(audioChunks, { type: 'audio/mp3' });
      const audioUrl = URL.createObjectURL(audioBlob);
      messages.value.push({ audio: audioUrl });
    };

    audioRecorder.start();
  }
}

// Fonction d'enregistrement vidéo
async function toggleVideoRecording() {
  if (videoRecorder && videoRecorder.state === 'recording') {
    videoRecorder.stop();
  } else {
    const stream = await navigator.mediaDevices.getUserMedia({ video: true });
    videoRecorder = new MediaRecorder(stream);

    videoRecorder.ondataavailable = (e) => videoChunks.push(e.data);
    videoRecorder.onstop = () => {
      const videoBlob = new Blob(videoChunks, { type: 'video/mp4' });
      const videoUrl = URL.createObjectURL(videoBlob);
      messages.value.push({ video: videoUrl });
    };

    videoRecorder.start();
  }
}
</script>

<style scoped>
.message-container {
  display: flex;
  gap: 10px;
  align-items: center;
  margin-bottom: 20px;
}
.record-container {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}
.messages-display {
  margin-top: 20px;
}
.message-item {
  padding: 10px;
  border-bottom: 1px solid #ccc;
}
.q-field__inner {
  width: 1000px;
}
</style>

<style src="../css/quasar.variables.scss" ></style>
<style src="../css/messagerie.scss" scoped></style>
