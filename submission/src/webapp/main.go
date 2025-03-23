package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/shinroo/fp/src/webapp/internal/api"
	"github.com/shinroo/fp/src/webapp/internal/app"
	"github.com/shinroo/fp/src/webapp/internal/auth"
	"github.com/shinroo/fp/src/webapp/internal/middleware"
	"github.com/shinroo/fp/src/webapp/internal/repository"
	"github.com/shinroo/fp/src/webapp/pkg/db"
)

func main() {
	dbConn, err := db.Open(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	accountRepo := repository.NewAccount(dbConn)
	sessionRepo := repository.NewSession(dbConn)
	profileRepo := repository.NewProfile(dbConn)
	spcaRepo := repository.NewSPCA(dbConn)
	dogRepo := repository.NewDog(dbConn)

	appRouter := http.NewServeMux()
	apiRouter := http.NewServeMux()
	authRouter := http.NewServeMux()

	appServer := app.NewServer(profileRepo)
	apiServer := api.NewServer(
		accountRepo,
		sessionRepo,
		profileRepo,
		spcaRepo,
		dogRepo,
	)
	authServer := auth.NewServer()

	// login
	authRouter.HandleFunc("/auth/login", authServer.Login)
	apiRouter.HandleFunc("/api/login", apiServer.Login)

	// signup
	authRouter.HandleFunc("/auth/signup", authServer.Signup)
	apiRouter.HandleFunc("/api/signup", apiServer.Signup)

	// home
	appRouter.HandleFunc("/app/home", appServer.Home)

	// profile
	appRouter.HandleFunc("/app/profile", appServer.Profile)
	apiRouter.HandleFunc("/api/profile", apiServer.Profile)

	// search
	appRouter.HandleFunc("/app/search", appServer.Search)
	apiRouter.HandleFunc("/api/search", apiServer.Search)

	authenticatedAppRouter := middleware.WrapWithRedirectAuth(appRouter, sessionRepo, slog.Default())

	combinedRouter := http.NewServeMux()
	combinedRouter.Handle("/app/", authenticatedAppRouter)
	combinedRouter.Handle("/api/", apiRouter)
	combinedRouter.Handle("/auth/", authRouter)

	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", combinedRouter))
}
