package main

import (
	"log"
	"net/http"
	"vade_microservice/controllers"
)

func main() {
	http.HandleFunc("/documents", controllers.HandleDocuments)
	http.HandleFunc("/documents/", controllers.HandleDocumentByID)

	log.Println("Server starting on port 8001...")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
