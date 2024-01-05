package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nebojsaj1726/movie-base/config"
	"github.com/nebojsaj1726/movie-base/internal/api/gql"
	"github.com/nebojsaj1726/movie-base/internal/database"
	"github.com/nebojsaj1726/movie-base/internal/database/services"
	"github.com/nebojsaj1726/movie-base/pkg/router"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

	repo, err := database.NewRepository()
	if err != nil {
		log.Fatalf("Error initializing database repository: %v", err)
	}

	movieService := services.NewMovieService(repo)

	resolver := gql.NewResolver(movieService)

    r := router.NewRouter(resolver)

    serverAddr := fmt.Sprintf(":%d", cfg.ServerPort)
    log.Printf("Server is running on http://localhost%s", serverAddr)
    if err := http.ListenAndServe(serverAddr, r); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
