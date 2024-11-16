package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// Structure des requêtes
type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Handler pour l'enregistrement
func RegisterHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := RegisterUser(db, req.Username, req.Email, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Utilisateur enregistré avec succès"))
}

// Handler pour la connexion
func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := LoginUser(db, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
