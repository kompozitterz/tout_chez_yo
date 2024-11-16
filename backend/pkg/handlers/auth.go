package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

func LoginUser(db *sql.DB, email, password string) (string, error) {
  var hashedPassword string
  var userID string

  row := db.QueryRow(`SELECT id, password FROM users WHERE email = ?`, email)
  if err := row.Scan(&userID, &hashedPassword); err != nil {
      return "", err
  }

  if !CheckPasswordHash(password, hashedPassword) {
      return "", errors.New("mot de passe invalide")
  }

  token, err := GenerateJWT(userID)
  if err != nil {
      return "", err
  }

  return token, nil
}

// Handler pour la connexion
func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requête invalide", http.StatusBadRequest)
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


// GenerateJWT (si non défini ailleurs)
func GenerateJWT(userID string) (string, error) {
	if len(SECRET_KEY) == 0 {
		return "", errors.New("clé secrète JWT non définie")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET_KEY)
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
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	if err := RegisterUser(db, req.Username, req.Email, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Utilisateur enregistré avec succès"))
}
