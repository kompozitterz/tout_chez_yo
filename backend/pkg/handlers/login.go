package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
  Username string `json:"username"`
  Email    string `json:"email"`
  Password string `json:"password"`
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

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
	var req User
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
