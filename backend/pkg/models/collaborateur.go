package models

type Collaborateur struct {
	ID            int64  `json:"id"`
	Nom           string `json:"nom"`
	Prenom        string `json:"prenom"`
	Profession    string `json:"profession"`
	DateNaissance string `json:"date_naissance"`
	DateEntree    string `json:"date_entree"`
	Adresse       string `json:"adresse"`
	Nationalite   string `json:"nationalite"`
	Telephone     string `json:"telephone"`
	Email         string `json:"email"`
}
