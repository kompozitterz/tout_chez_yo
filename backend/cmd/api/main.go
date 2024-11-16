package main

import (
	"backend/pkg/database"
	"backend/pkg/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Charger la configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port par défaut
	}

	// Initialiser la base de données
	db := database.InitDB()
	defer db.Close()

	// Configurer le routeur
	router := mux.NewRouter()

	// Ajouter des routes
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, db)
	}).Methods("POST")

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, db)
	}).Methods("POST")

	// Lancer le serveur
	log.Printf("Serveur démarré sur http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Erreur au démarrage du serveur : %v", err)
	}
}
