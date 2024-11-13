package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

// Fonction pour générer une clé secrète aléatoire
func generateSecretKey() (string, error) {
    key := make([]byte, 32) // 32 bytes = 256 bits
    if _, err := rand.Read(key); err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(key), nil
}

// Fonction pour mettre à jour la clé secrète dans le fichier .env
func updateEnvFile(secretKey string) error {
    envFile := ".env"

    // Vérifie si le fichier existe, sinon le crée
    if _, err := os.Stat(envFile); os.IsNotExist(err) {
        file, err := os.Create(envFile)
        if err != nil {
            return err
        }
        defer file.Close()
    }

    // Lire le contenu du fichier .env
    content, err := os.ReadFile(envFile)
    if err != nil {
        return err
    }

    // Convertir le contenu en une chaîne et diviser par lignes
    lines := strings.Split(string(content), "\n")

    // Chercher et remplacer la ligne contenant "SECRET_KEY="
    keyExists := false
    for i, line := range lines {
        if strings.HasPrefix(line, "SECRET_KEY=") {
            lines[i] = "SECRET_KEY=" + secretKey
            keyExists = true
            break
        }
    }

    // Ajouter la clé si elle n'existe pas encore
    if !keyExists {
        lines = append(lines, "SECRET_KEY="+secretKey)
    }

    // Joindre les lignes et réécrire le fichier .env
    output := strings.Join(lines, "\n")
    err = os.WriteFile(envFile, []byte(output), 0644)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    // Générer une nouvelle clé secrète
    secretKey, err := generateSecretKey()
    if err != nil {
        fmt.Println("Erreur lors de la génération de la clé :", err)
        return
    }

    // Mettre à jour le fichier .env
    err = updateEnvFile(secretKey)
    if err != nil {
        fmt.Println("Erreur lors de la mise à jour de .env :", err)
    } else {
        fmt.Println("La clé secrète a été mise à jour avec succès dans .env")
    }
}
