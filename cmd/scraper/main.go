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

	movies, err := scraper.ScrapeMedia("movies")
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
			ImageURL:    dbMovie.ImageURL,
			Actors:      dbMovie.Actors,
		}
		scraperMovies = append(scraperMovies, scraperMovie)
	}

	if err := repo.CreateMovies(scraperMovies); err != nil {
		log.Fatalf("Error creating movies in the database: %v", err)
	}

	shows, err := scraper.ScrapeMedia("shows")
	if err != nil {
		log.Fatal(err)
	}

	var scraperShows []scraper.Movie
	for _, dbMovie := range shows {
		scraperMovie := scraper.Movie{
			Title:       dbMovie.Title,
			Rate:        dbMovie.Rate,
			Year:        dbMovie.Year,
			Description: dbMovie.Description,
			Genres:      dbMovie.Genres,
			ImageURL:    dbMovie.ImageURL,
			Actors:      dbMovie.Actors,
		}
		scraperShows = append(scraperShows, scraperMovie)
	}

	if err := repo.CreateShows(scraperShows); err != nil {
		log.Fatalf("Error creating shows in the database: %v", err)
	}

	log.Printf("Scraping and database insertion completed. Total movies: %d, Total shows: %d\n", len(scraperMovies), len(scraperShows))
}
