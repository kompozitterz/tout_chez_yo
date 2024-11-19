package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// Clé secrète utilisée pour signer et vérifier les tokens JWT
var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))
// var SECRET_KEY []byte

// func init() {
//   // Chargez le fichier .env
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
// 	}
//   log.Println(".env chargé avec succès")

// 	secret := os.Getenv("SECRET_KEY")
// 	if secret == "" {
// 		log.Fatal("Erreur : SECRET_KEY n'est pas défini dans les variables d'environnement")
// 	}
// 	SECRET_KEY = []byte(secret)
// }

// Middleware pour vérifier l'authentification par token JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Récupère le header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Authorization header manquant")
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Vérifie le format du token (Bearer <token>)
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			log.Println("Format du token JWT invalide")
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Parse et valide le token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Méthode de signature JWT invalide")
				return nil, http.ErrAbortHandler
			}
			return SECRET_KEY, nil
		})

		if err != nil {
			log.Printf("Erreur de parsing JWT : %v\n", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Récupère les claims du token (par exemple l'ID de l'utilisateur)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Println("Claims JWT invalides")
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			log.Println("Claim 'user_id' manquant ou invalide dans le token JWT")
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Ajoute l'ID de l'utilisateur au contexte pour l'utiliser dans les handlers suivants
		ctx := context.WithValue(r.Context(), "user_id", userID)
		r = r.WithContext(ctx)

		// Passe au prochain handler
		next.ServeHTTP(w, r)
	})
}

// Middleware pour activer CORS
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Requête reçue :", r.Method, r.URL.Path) // Ajout de logs
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
