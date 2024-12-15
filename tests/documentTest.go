package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"vade_microservice/models"
	"vade_microservice/services"
)

func TestCreateDocument(t *testing.T) {

	t.Run("Create valid document", func(t *testing.T) {
		doc := models.Document{ID: "1", Name: "Doc 1", Description: "Description 1"}
		err := services.CreateDocument(doc)
		assert.NoError(t, err, "Should not return an error for valid document")
	})

	t.Run("Create document with duplicate ID", func(t *testing.T) {
		doc := models.Document{ID: "1", Name: "Duplicate", Description: "Duplicate"}
		err := services.CreateDocument(doc)
		assert.Error(t, err, "Should return an error for duplicate ID")
	})
}

func TestGetAllDocuments(t *testing.T) {

	t.Run("Get all documents when none exist", func(t *testing.T) {
		docs := services.GetAllDocuments()
		assert.Empty(t, docs, "Should return an empty list when no documents exist")
	})

	t.Run("Get all documents after creation", func(t *testing.T) {
		err := services.CreateDocument(models.Document{ID: "1", Name: "Doc 1", Description: "Description 1"})
		if err != nil {
			return
		}
		err1 := services.CreateDocument(models.Document{ID: "2", Name: "Doc 2", Description: "Description 2"})
		if err1 != nil {
			return
		}
		docs := services.GetAllDocuments()
		assert.Len(t, docs, 2, "Should return a list of 2 documents")
	})
}

func TestGetDocumentByID(t *testing.T) {
	err := services.CreateDocument(models.Document{ID: "1", Name: "Doc 1", Description: "Description 1"})
	if err != nil {
		return
	}

	t.Run("Get existing document by ID", func(t *testing.T) {
		doc, err := services.GetDocumentByID("1")
		assert.NoError(t, err, "Should not return an error for existing document")
		assert.Equal(t, "Doc 1", doc.Name, "Title should match")
	})

	t.Run("Get non-existing document by ID", func(t *testing.T) {
		_, err := services.GetDocumentByID("2")
		assert.Error(t, err, "Should return an error for non-existing document")
	})
}

func TestDeleteDocumentByID(t *testing.T) {
	err := services.CreateDocument(models.Document{ID: "1", Name: "Doc 1", Description: "Description 1"})
	if err != nil {
		return
	}

	t.Run("Delete existing document", func(t *testing.T) {
		err := services.DeleteDocumentByID("1")
		assert.NoError(t, err, "Should not return an error for deleting existing document")

		// Verify deletion
		_, err = services.GetDocumentByID("1")
		assert.Error(t, err, "Deleted document should not be found")
	})

	t.Run("Delete non-existing document", func(t *testing.T) {
		err := services.DeleteDocumentByID("2")
		assert.Error(t, err, "Should return an error for deleting non-existing document")
	})
}
