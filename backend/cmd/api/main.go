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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := database.InitDB()
	defer db.Close()

	router := mux.NewRouter()
  router.Use(handlers.EnableCORS)

  router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
      handlers.RegisterHandler(w, r, db)
  }).Methods("POST")

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, db)
	}).Methods("POST")

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, db)
	}).Methods("POST")

	log.Printf("Serveur démarré sur http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Erreur au démarrage du serveur : %v", err)
	}
}
