package handlers

import (
	"auth/pkg/database"
	"encoding/json"
	"net/http"
)

// Получение документов только для авторизованных
func GetDocuments(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, title, content FROM documents")
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var documents []database.Document
	for rows.Next() {
		var doc database.Document
		if err := rows.Scan(&doc.ID, &doc.Title, &doc.Content); err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		documents = append(documents, doc)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(documents)
}