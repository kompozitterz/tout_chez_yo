package handlers

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Clé secrète pour signer le token JWT
var SECRET_KEY []byte

// HashPassword hache le mot de passe en utilisant bcrypt
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// CheckPasswordHash compare un mot de passe en clair avec son haché
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// GenerateJWT génère un token JWT pour l'utilisateur
func GenerateJWT(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Expiration dans 24 heures
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func RegisterUser(db *sql.DB, username, email, password string) error {
    hashedPassword, err := HashPassword(password)
    if err != nil {
        return err
    }

    _, err = db.Exec(`INSERT INTO users (id, username, email, password) VALUES (?, ?, ?, ?)`,
        uuid.New().String(), username, email, hashedPassword)
    return err
}

func LoginUser(db *sql.DB, email, password string) (string, error) {
    var hashedPassword string
    var userID string

    row := db.QueryRow(`SELECT id, password FROM users WHERE email = ?`, email)
    err := row.Scan(&userID, &hashedPassword)
    if err != nil {
        return "", err
    }

    if !CheckPasswordHash(password, hashedPassword) {
        return "", errors.New("invalid password")
    }

    return GenerateJWT(userID)
}
