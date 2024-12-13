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
		services.CreateDocument(doc)
		w.WriteHeader(http.StatusCreated)
	case "GET":
		docs := services.GetAllDocuments()
		json.NewEncoder(w).Encode(docs)
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
		doc, found := services.GetDocumentByID(id)
		if !found {
			http.Error(w, "Document not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(doc)
	case "DELETE":
		if services.DeleteDocumentByID(id) {
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(w, "Document not found", http.StatusNotFound)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
