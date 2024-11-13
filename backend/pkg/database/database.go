package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
    // Vérifier que le dossier backend/database existe, sinon le créer
    if _, err := os.Stat("./backend/database"); os.IsNotExist(err) {
        err = os.MkdirAll("./backend/database", 0755)
        if err != nil {
            log.Fatal("Erreur lors de la création du dossier database:", err)
        }
    }

    // Ouvrir ou créer la base de données auth.db
    db, err := sql.Open("sqlite3", "./backend/database/auth.db")
    if err != nil {
        log.Fatal("Erreur lors de l'ouverture de la base de données:", err)
    }

    // Créer la table des utilisateurs si elle n'existe pas
    statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        username TEXT NOT NULL,
        email TEXT UNIQUE,
        password TEXT NOT NULL
    );`)
    if err != nil {
        log.Fatal("Erreur lors de la préparation de la création de table:", err)
    }
    statement.Exec()

    return db
}
