package database

import (
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nebojsaj1726/movie-base/config"
	"github.com/nebojsaj1726/movie-base/internal/scraper"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository() (*Repository, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open("postgres", cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	db.LogMode(false)

	return &Repository{DB: db}, nil
}

func (r *Repository) CreateMovies(movies []scraper.Movie) error {
	for _, movie := range movies {

		newMovie := Movie{
			Title:       movie.Title,
			Rate:        movie.Rate,
			Year:        movie.Year,
			Description: movie.Description,
			Duration:    movie.Duration,
			ImageURL:    movie.ImageURL,
            Genres:      strings.Join(movie.Genres, ", "),
		}

		if err := r.DB.Create(&newMovie).Error; err != nil {
			return fmt.Errorf("error creating movie %s: %v", movie.Title, err)
		}
	}

	return nil
}

func (r *Repository) CreateShows(shows []scraper.Movie) error {
    for _, show := range shows {
        newShow := Show{
            Title:       show.Title,
            Rate:        show.Rate,
            Year:        show.Year,
            Description: show.Description,
            ImageURL:    show.ImageURL,
            Genres:      strings.Join(show.Genres, ", "),
        }

        if err := r.DB.Create(&newShow).Error; err != nil {
            return fmt.Errorf("error creating show %s: %v", show.Title, err)
        }
    }

    return nil
}


func (r *Repository) Close() {
	if err := r.DB.Close(); err != nil {
		log.Println("Error closing the database connection:", err)
	}
}
