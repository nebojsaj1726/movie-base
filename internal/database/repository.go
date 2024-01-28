package database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nebojsaj1726/movie-base/config"
	"github.com/nebojsaj1726/movie-base/internal/scraper"
)

const cacheExpiration = 3 * time.Hour

type Repository struct {
	DB          *gorm.DB
	RedisClient *RedisClient
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

	redisClient, err := NewRedisClient()
	if err != nil {
		return nil, fmt.Errorf("error initializing Redis client: %v", err)
	}

	repo := &Repository{DB: db, RedisClient: redisClient}
	return repo, nil
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

func (r *Repository) SearchMoviesByKeyword(keyword string) ([]*Movie, error) {
	var movies []*Movie
	if err := r.DB.Where("title ILIKE ?", "%"+keyword+"%").Limit(15).Find(&movies).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, fmt.Errorf("error searching movies by keyword: %v", err)
		}
	}

	return movies, nil
}

func (r *Repository) GetMovies(limit, offset *int, genreRange []string, year *int, rating *float64) ([]*Movie, error) {
	cacheKey := fmt.Sprintf("GetMovies_%v_%v_%v_%v_%v", *limit, *offset, genreRange, nilCheck(year), nilCheck(rating))

	var cachedMovies []*Movie
	if result, err := r.RedisClient.GetCache(context.Background(), cacheKey, &cachedMovies); err == nil {
		return *result.(*[]*Movie), nil
	}

	var movies []*Movie
	query := r.DB.Limit(*limit).Offset(*offset)

	if len(genreRange) > 0 {
		query = query.Where("genres IN (?)", genreRange)
	}
	if year != nil {
		query = query.Where("year = ?", *year)
	}
	if rating != nil {
		query = query.Where("rate >= ?", *rating)
	}
	if err := query.Find(&movies).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, fmt.Errorf("error retrieving movies: %v", err)
		}
	}
	if err := r.RedisClient.SetCache(context.Background(), cacheKey, movies, cacheExpiration); err != nil {
		log.Println("Error caching result:", err)
	}

	return movies, nil
}

func (r *Repository) GetMovieByID(id uint) (*Movie, error) {
	var movie Movie

	if err := r.DB.First(&movie, id).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, fmt.Errorf("error retrieving movie by ID: %v", err)
		}
		return nil, nil
	}

	return &movie, nil
}

func (r *Repository) GetRandomMovies(count *int, genreRange []string, year *int, rating *float64) ([]*Movie, error) {
	var movies []*Movie
	query := r.DB

	if len(genreRange) > 0 {
		query = query.Where("genres IN (?)", genreRange)
	}
	if year != nil {
		query = query.Where("year = ?", *year)
	}
	if rating != nil {
		query = query.Where("rate >= ?", *rating)
	}

	if err := query.Order("RANDOM()").Limit(*count).Find(&movies).Error; err != nil {
		return nil, fmt.Errorf("error retrieving random movies: %v", err)
	}

	return movies, nil
}

func nilCheck(v interface{}) interface{} {
	switch val := v.(type) {
	case *int:
		if val != nil {
			return *val
		}
	case *float64:
		if val != nil {
			return *val
		}
	}
	return nil
}

func (r *Repository) Close() error {
	if err := r.DB.Close(); err != nil {
		return fmt.Errorf("error closing the database connection: %v", err)
	}
	return nil
}
