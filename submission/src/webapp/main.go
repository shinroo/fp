package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shinroo/fp/src/webapp/internal/api"
	"github.com/shinroo/fp/src/webapp/internal/app"
	"github.com/shinroo/fp/src/webapp/pkg/db"
)

func main() {
	dbConn, err := db.Open(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

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
