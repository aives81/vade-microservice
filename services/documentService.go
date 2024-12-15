package services

import (
	"errors"
	"vade_microservice/models"
)

func CreateDocument(doc models.Document) error {
	// Fields validation
	if doc.ID == "" || doc.Name == "" || doc.Description == "" {
		return errors.New("invalid document: ID, name, and description are required")
	}

	// Check if the document already exists
	if _, exists := models.Documents[doc.ID]; exists {
		return errors.New("document with the same ID already exists")
	}

	// Add document
	models.Documents[doc.ID] = doc
	return nil
}

func GetAllDocuments() []models.Document {
	docs := make([]models.Document, 0, len(models.Documents))
	for _, doc := range models.Documents {
		docs = append(docs, doc)
	}
	return docs
}

func GetDocumentByID(id string) (models.Document, error) {
	if _, exists := models.Documents[id]; !exists {
		return models.Document{}, errors.New("document not found")
	}

	return models.Documents[id], nil
}

func DeleteDocumentByID(id string) error {
	// Check if the document exists
	if _, exists := models.Documents[id]; !exists {
		return errors.New("document not found")
	}

	// Delete the document
	delete(models.Documents, id)
	return nil
}
