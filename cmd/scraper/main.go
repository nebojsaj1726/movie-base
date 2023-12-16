package main

import (
	"log"

	"github.com/nebojsaj1726/movie-base/internal/database"
	"github.com/nebojsaj1726/movie-base/internal/scraper"
)

func main() {
	repo, err := database.NewRepository()
	if err != nil {
		log.Fatal("Error initializing the database repository:", err)
	}
	defer repo.Close()

	movies, err := scraper.ScrapeMovies()
	if err != nil {
		log.Fatal(err)
	}

	var scraperMovies []scraper.Movie
	for _, dbMovie := range movies {
		scraperMovie := scraper.Movie{
			Title:       dbMovie.Title,
			Rate:        dbMovie.Rate,
			Year:        dbMovie.Year,
			Description: dbMovie.Description,
			Genres:      dbMovie.Genres,
			Duration:    dbMovie.Duration,
		}
		scraperMovies = append(scraperMovies, scraperMovie)
	}

    if err := repo.CreateMovies(scraperMovies); err != nil {
		log.Fatalf("Error creating movies in the database: %v", err)
	}

    log.Printf("Scraping and database insertion completed. Total Movies: %d\n", len(scraperMovies))
}