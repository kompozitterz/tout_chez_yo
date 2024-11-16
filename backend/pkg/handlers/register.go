package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Structure des requêtes
type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CheckPasswordHash compare un mot de passe en clair avec son haché
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword hache le mot de passe
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("erreur lors du hashage : %w", err)
	}
	return string(bytes), nil
}

// RegisterUser enregistre un utilisateur
func RegisterUser(db *sql.DB, username, email, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return fmt.Errorf("erreur lors du hashage : %w", err)
	}

	_, err = db.Exec(`INSERT INTO users (id, username, email, password) VALUES (?, ?, ?, ?)`,
		uuid.New().String(), username, email, hashedPassword)
	if err != nil {
		return fmt.Errorf("erreur lors de l'insertion : %w", err)
	}

	return nil
}

// Handler pour l'enregistrement
func RegisterHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    log.Println("Erreur de décodage JSON :", err)
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	if err := RegisterUser(db, req.Username, req.Email, req.Password); err != nil {
    log.Println("Erreur d’enregistrement :", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Utilisateur enregistré avec succès"))
}
