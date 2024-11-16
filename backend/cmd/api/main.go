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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := database.InitDB()
	defer db.Close()

	router := mux.NewRouter()
  router.Use(middleware.EnableCORS)

  router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodOptions {
      log.Println("Requête OPTIONS capturée directement")
      w.Header().Set("Access-Control-Allow-Origin", "*")
      w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
      w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
      w.WriteHeader(http.StatusOK)
      return
    }
    handlers.RegisterHandler(w, r, db)
  }).Methods("OPTIONS", "POST")


	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, db)
	}).Methods("POST")

	log.Printf("Serveur démarré sur http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Erreur au démarrage du serveur : %v", err)
	}
}
