package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initialise la base de données et retourne une instance *sql.DB
func InitDB() *sql.DB {
	// Vérifier que le dossier backend/database existe
	if _, err := os.Stat("./backend/database"); os.IsNotExist(err) {
		err = os.MkdirAll("./backend/database", 0755)
		if err != nil {
			log.Fatal("Erreur lors de la création du dossier database:", err)
		}
		log.Println("Dossier './backend/database' créé avec succès.")
	}

	// Ouvrir ou créer la base de données auth.db
	db, err := sql.Open("sqlite3", "./backend/database/auth.db")
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base de données:", err)
	}
	log.Println("Base de données './backend/database/auth.db' ouverte avec succès.")

	// Créer la table des utilisateurs si elle n'existe pas
	createUsersTable(db)

	// Créer la table des collaborateurs si elle n'existe pas
	createCollaboratorsTable(db)

	return db
}

// createUsersTable vérifie et crée la table des utilisateurs
func createUsersTable(db *sql.DB) {
	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        username TEXT NOT NULL,
        email TEXT UNIQUE,
        password TEXT NOT NULL
    );`)
	if err != nil {
		log.Fatal("Erreur lors de la préparation de la création de la table 'users':", err)
	}

	if _, err = statement.Exec(); err != nil {
		log.Fatal("Erreur lors de l'exécution de la création de la table 'users':", err)
	}
	log.Println("Table 'users' vérifiée ou créée avec succès.")
}

// createCollaboratorsTable vérifie et crée la table des collaborateurs
func createCollaboratorsTable(db *sql.DB) {
	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS collaborateurs (
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
	);`)
	if err != nil {
		log.Fatal("Erreur lors de la préparation de la création de la table 'collaborateurs':", err)
	}

	if _, err = statement.Exec(); err != nil {
		log.Fatal("Erreur lors de l'exécution de la création de la table 'collaborateurs':", err)
	}
	log.Println("Table 'collaborateurs' vérifiée ou créée avec succès.")
}
