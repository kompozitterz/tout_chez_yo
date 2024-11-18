package main

import (
	"backend/pkg/database"
	"backend/pkg/handlers"
	"backend/pkg/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Récupérer le port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialiser la base de données
	db := database.InitDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Erreur lors de la fermeture de la base de données : %v", err)
		}
	}()

	// Configurer le routeur
	router := mux.NewRouter()
	router.Use(middleware.EnableCORS)

	// Route pour l'inscription
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, db)
	}).Methods("POST")

	// Route pour la connexion
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		handlers.LoginHandler(w, r, db)
	}).Methods("POST", "OPTIONS")

	// Route pour ajouter un collaborateur
	router.HandleFunc("/addCollaborator", func(w http.ResponseWriter, r *http.Request) {
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
