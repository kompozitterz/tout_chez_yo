import axios from 'axios';

// Création de l'instance Axios
const apiClient = axios.create({
  baseURL: 'http://localhost:8080', // Adresse de votre backend
  headers: {
    'Content-Type': 'application/json',
  },
});

// Intercepteur de requêtes (pour ajouter des tokens ou logs)
apiClient.interceptors.request.use(
  (config) => {
    // Ajouter un token d'authentification si présent
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    // eslint-disable-next-line no-console
    console.error('Erreur lors de la requête :', error);
    return Promise.reject(error);
  },
);

// Intercepteur de réponses (pour gérer les erreurs globales)
apiClient.interceptors.response.use(
  (response) => response, // Réponse réussie
  (error) => {
    if (error.response) {
      // Erreur reçue depuis le serveur
      // eslint-disable-next-line no-console
      console.error('Erreur du serveur :', error.response);
      if (error.response.status === 401) {
        // eslint-disable-next-line no-console
        console.warn('Non autorisé, redirection vers la page de connexion...');
        localStorage.removeItem('token'); // Supprimer le token
        window.location.href = '/login'; // Rediriger
      }
    } else if (error.request) {
      // Pas de réponse du serveur
      // eslint-disable-next-line no-console
      console.error('Pas de réponse du serveur :', error.request);
    } else {
      // Autre erreur
      // eslint-disable-next-line no-console
      console.error('Erreur Axios :', error.message);
    }
    return Promise.reject(error); // Propager l'erreur pour un traitement local
  },
);

export default apiClient;
