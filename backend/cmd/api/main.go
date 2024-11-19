package main

import (
	"backend/pkg/database"
	"backend/pkg/handlers"
	"backend/pkg/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var SECRET_KEY []byte

func main() {
	// Chemin absolu vers le fichier .env
	envPath := "/Users/cailloux/Desktop/OPotable/tout_chez_yo/backend/scripts/.env"

	// Charger le fichier .env
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}
	log.Println("Fichier .env chargé avec succès")

	// Récupérer la clé secrète
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("Erreur : SECRET_KEY n'est pas défini dans les variables d'environnement")
	}
	SECRET_KEY = []byte(secretKey)
	log.Printf("SECRET_KEY chargé : %s", string(SECRET_KEY[:len(SECRET_KEY)/2])+"...")

	// Afficher le répertoire de travail courant
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Erreur lors de la récupération du répertoire courant : %v", err)
	}
	log.Printf("Répertoire courant : %s", dir)

	// Récupérer le port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valeur par défaut
	}
	log.Printf("Port défini : %s", port)

	// Initialiser la base de données
	db := database.InitDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Erreur lors de la fermeture de la base de données : %v", err)
		}
	}()
	log.Println("Base de données initialisée avec succès")

	// Configurer le routeur
	router := mux.NewRouter()
	router.Use(middleware.EnableCORS)

	// Routes
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		handlers.RegisterHandler(w, r, db)
	}).Methods("POST", "OPTIONS")

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		handlers.LoginHandler(w, r, db)
	}).Methods("POST", "OPTIONS")

	router.HandleFunc("/collaborateur", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		handlers.AddCollaborator(w, r, db)
	}).Methods("POST", "OPTIONS")

	// Démarrer le serveur
	log.Printf("Serveur démarré sur http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Erreur au démarrage du serveur : %v", err)
	}
}
