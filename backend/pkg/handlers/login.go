package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(db *sql.DB, email, password string) (string, error) {
	log.Printf("Tentative de connexion avec l'email : %s", email)

	var hashedPassword string
	var userID string

	row := db.QueryRow(`SELECT id, password FROM users WHERE email = ?`, email)
	if err := row.Scan(&userID, &hashedPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("Utilisateur non trouvé dans la base de données")
			return "", errors.New("UTILISATEUR NON TROUVÉ")
		}
		log.Printf("Erreur SQL : %v", err)
		return "", err
	}

	log.Printf("Utilisateur trouvé : ID=%s, Mot de passe haché=%s", userID, hashedPassword)

	// Vérifier le mot de passe
	if !CheckPasswordHash(password, hashedPassword) {
		log.Println("Mot de passe incorrect")
		return "", errors.New("MOT DE PASSE INVALIDE")
	}

	// Générer le token JWT
	token, err := GenerateJWT(userID)
	if err != nil {
		log.Printf("Erreur lors de la génération du JWT : %v", err)
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

	response := map[string]interface{}{
		"token": token,
		"user": map[string]string{
			"email": req.Email,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}


// GenerateJWT génère un token JWT avec une clé par défaut
func GenerateJWT(userID string) (string, error) {
	// Définir une clé par défaut pour signer le JWT
	defaultSecretKey := []byte("default_secret_key")

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(defaultSecretKey)
}

