package main

import (
	// "encoding/json"
	"fmt"
	"log"

	"github.com/nebojsaj1726/movie-base/internal/scraper"
)

func main() {
	movies, err := scraper.ScrapeMovies()
	if err != nil {
		log.Fatal(err)
	}

	for i, movie := range movies {
		fmt.Printf("Movie %d:\n", i+1)
		fmt.Printf("Title: %s\n", movie.Title)
		fmt.Printf("Rate: %s\n", movie.Rate)
		fmt.Printf("Year: %s\n", movie.Year)
		fmt.Printf("Genres: %v\n", movie.Genres)
		fmt.Printf("Duration: %s\n", movie.Duration)
		fmt.Printf("Description: %s\n", movie.Description)
		fmt.Println("--------------------------")
	}

	// moviesJSON, err := json.MarshalIndent(movies, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}