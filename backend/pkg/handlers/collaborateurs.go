package handlers

import (
	"backend/pkg/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

// Ajouter un collaborateur
func AddCollaborator(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var collaborateur models.Collaborateur
	if err := json.NewDecoder(r.Body).Decode(&collaborateur); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	query := `
	INSERT INTO collaborateurs (nom, prenom, profession, date_naissance, date_entree, adresse, nationalite, telephone, email)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, collaborateur.Nom, collaborateur.Prenom, collaborateur.Profession, collaborateur.DateNaissance,
		collaborateur.DateEntree, collaborateur.Adresse, collaborateur.Nationalite, collaborateur.Telephone, collaborateur.Email)
	if err != nil {
		http.Error(w, "Erreur lors de l'insertion", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	collaborateur.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(collaborateur)
}
