package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"vade_microservice/models"
	"vade_microservice/services"
)

func HandleDocuments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var doc models.Document
		err := json.NewDecoder(r.Body).Decode(&doc)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		err1 := services.CreateDocument(doc)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte(`{"message": "Document created successfully"}`))
		if err != nil {
			return
		}
	case "GET":
		docs := services.GetAllDocuments()
		if len(docs) == 0 {
			http.Error(w, "Nothing to display", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(docs)
		if err != nil {
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleDocumentByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/documents/")
	if id == "" {
		http.Error(w, "Missing document ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		doc, _ := services.GetDocumentByID(id)
		if doc == (models.Document{}) {
			http.Error(w, "Document not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(doc)
	case "DELETE":
		if services.DeleteDocumentByID(id) == nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(w, "Document not found", http.StatusNotFound)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
