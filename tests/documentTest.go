package tests

import (
	"testing"
	"vade_microservice/models"
	"vade_microservice/services"
)

func TestCreateAndGetDocument(t *testing.T) {
	doc := models.Document{ID: "1", Name: "Doc1", Description: "Description1"}
	services.CreateDocument(doc)

	retDoc, found := services.GetDocumentByID("1")
	if !found || retDoc.Name != "Doc1" {
		t.Errorf("Expected document not found or mismatch")
	}
}

func TestDeleteDocument(t *testing.T) {
	doc := models.Document{ID: "2", Name: "Doc2", Description: "Description2"}
	services.CreateDocument(doc)

	if !services.DeleteDocumentByID("2") {
		t.Errorf("Failed to delete document")
	}
	if _, found := services.GetDocumentByID("2"); found {
		t.Errorf("Document was not deleted")
	}
}
