package services

import (
	"vade_microservice/models"
)

func CreateDocument(doc models.Document) {
	models.Documents[doc.ID] = doc
}

func GetAllDocuments() []models.Document {
	docs := make([]models.Document, 0, len(models.Documents))
	for _, doc := range models.Documents {
		docs = append(docs, doc)
	}
	return docs
}

func GetDocumentByID(id string) (models.Document, bool) {
	doc, found := models.Documents[id]
	return doc, found
}

func DeleteDocumentByID(id string) bool {
	if _, found := models.Documents[id]; found {
		delete(models.Documents, id)
		return true
	}
	return false
}
