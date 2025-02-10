package main

import (
	"log"
	"net/http"

	"github.com/shinroo/fp/src/webapp/internal/handlers"
)

func main() {
	unauthenticatedRouter := http.NewServeMux()

	server := handlers.NewServer()

	unauthenticatedRouter.HandleFunc("/app/login", server.Login)

	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", unauthenticatedRouter))
}
