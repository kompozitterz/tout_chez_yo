package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// Clé secrète utilisée pour signer et vérifier les tokens JWT (à garder confidentielle)
var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

// Middleware pour vérifier l'authentification par token JWT
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Récupère le header Authorization
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header missing", http.StatusUnauthorized)
            return
        }

        // Vérifie le format du token (Bearer <token>)
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            http.Error(w, "Invalid token format", http.StatusUnauthorized)
            return
        }

        // Parse et valide le token JWT
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, http.ErrAbortHandler
            }
            return SECRET_KEY, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Récupère les claims du token (par exemple l'ID de l'utilisateur)
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }

        // Ajoute l'ID de l'utilisateur au contexte pour l'utiliser dans les handlers suivants
        userID := claims["user_id"].(string)
        ctx := context.WithValue(r.Context(), "user_id", userID)
        r = r.WithContext(ctx)

        // Passe au prochain handler
        next.ServeHTTP(w, r)
    })
}

func EnableCORS(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      log.Println("Requête reçue :", r.Method, r.URL.Path) // Ajout de logs
      w.Header().Set("Access-Control-Allow-Origin", "*")
      w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
      w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

      if r.Method == http.MethodOptions {
          w.WriteHeader(http.StatusOK)
          return
      }

      next.ServeHTTP(w, r)
  })
}
