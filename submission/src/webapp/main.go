package main

import (
	"log"
	"net/http"

	"github.com/shinroo/fp/src/webapp/internal/api"
	"github.com/shinroo/fp/src/webapp/internal/app"
)

func main() {
	unauthenticatedRouter := http.NewServeMux()

	appServer := app.NewServer()
	apiServer := api.NewServer()

	// login
	unauthenticatedRouter.HandleFunc("/app/login", appServer.Login)
	unauthenticatedRouter.HandleFunc("/api/login", apiServer.Login)

	// signup
	unauthenticatedRouter.HandleFunc("/app/signup", appServer.Signup)
	unauthenticatedRouter.HandleFunc("/api/signup", apiServer.Signup)

	// home
	unauthenticatedRouter.HandleFunc("/app/home", appServer.Home)

	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", unauthenticatedRouter))
}
