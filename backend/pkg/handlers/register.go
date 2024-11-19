package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// Structure des requêtes
// type UserRequest struct {
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

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
func RegisterHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	// Hasher le mot de passe
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Erreur lors du hachage du mot de passe", http.StatusInternalServerError)
		return
	}

	// Insertion dans la base de données
	_, err = db.Exec(`INSERT INTO users (username, email, password) VALUES (?, ?, ?)`,
		user.Username, user.Email, hashedPassword)
	if err != nil {
		// Vérifier si c'est une erreur liée à une contrainte UNIQUE
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			http.Error(w, "Cet email est déjà utilisé", http.StatusConflict)
			return
		}
		http.Error(w, "Erreur lors de l'insertion : "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Utilisateur créé avec succès"))
}


