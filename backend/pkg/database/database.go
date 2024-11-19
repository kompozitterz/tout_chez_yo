package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initialise la base de données et retourne une instance *sql.DB
func InitDB() *sql.DB {
	// Obtenir le répertoire courant
	baseDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Erreur lors de la récupération du répertoire courant : %v", err)
	}

	// Construire le chemin absolu vers le dossier de la base de données
	databasePath := filepath.Join(baseDir)
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		err = os.MkdirAll(databasePath, 0755)
		if err != nil {
			log.Fatalf("Erreur lors de la création du dossier '%s': %v", databasePath, err)
		}
		log.Printf("Dossier '%s' créé avec succès.\n", databasePath)
	}

	// Chemin complet vers auth.db
	dbFilePath := filepath.Join(databasePath, "auth.db")

	// Ouvrir ou créer la base de données auth.db
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la base de données '%s': %v", dbFilePath, err)
	}
	log.Printf("Base de données '%s' ouverte avec succès.\n", dbFilePath)

	// Créer les tables nécessaires
	createUsersTable(db)
	createCollaboratorsTable(db)

	return db
}

// createUsersTable vérifie et crée la table des utilisateurs
func createUsersTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		username TEXT NOT NULL,
		email TEXT UNIQUE,
		password TEXT NOT NULL
	);`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Erreur lors de la création de la table 'users': %v", err)
	}
	log.Println("Table 'users' vérifiée ou créée avec succès.")
}

// createCollaboratorsTable vérifie et crée la table des collaborateurs
func createCollaboratorsTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS collaborateurs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nom TEXT NOT NULL,
		prenom TEXT NOT NULL,
		profession TEXT NOT NULL,
		date_naissance TEXT,
		date_entree TEXT NOT NULL,
		adresse TEXT,
		nationalite TEXT,
		telephone TEXT,
		email TEXT
	);`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Erreur lors de la création de la table 'collaborateurs': %v", err)
	}
	log.Println("Table 'collaborateurs' vérifiée ou créée avec succès.")
}
