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
			Actors:      strings.Join(movie.Actors, ", "),
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
			Actors:      strings.Join(show.Actors, ", "),
		}

		if err := r.DB.Create(&newShow).Error; err != nil {
			return fmt.Errorf("error creating show %s: %v", show.Title, err)
		}
	}

	return nil
}

func (r *Repository) SearchMoviesByKeyword(keyword string) ([]*Movie, []*Show, error) {
	var movies []*Movie
	var shows []*Show

	if err := r.DB.Where("title ILIKE ?", "%"+keyword+"%").Limit(8).Find(&movies).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, nil, fmt.Errorf("error searching movies by keyword: %v", err)
		}
	}

	if err := r.DB.Where("title ILIKE ?", "%"+keyword+"%").Limit(7).Find(&shows).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, nil, fmt.Errorf("error searching shows by keyword: %v", err)
		}
	}

	return movies, shows, nil
}

func (r *Repository) GetMovies(limit, offset *int, genreRange []string, year *int, rating *float64) ([]*Movie, int, error) {
	cacheKey := fmt.Sprintf("GetMovies_%v_%v_%v_%v_%v", *limit, *offset, genreRange, nilCheck(year), nilCheck(rating))

	type CachedMovies struct {
		Movies     []*Movie
		TotalCount int
	}

	var cachedData CachedMovies
	if _, err := r.RedisClient.GetCache(context.Background(), cacheKey, &cachedData); err == nil {
		return cachedData.Movies, cachedData.TotalCount, nil
	}

	var movies []*Movie
	var totalCount int
	query := r.DB.Model(&Movie{})

	if len(genreRange) > 0 {
		genreParams := make([]interface{}, len(genreRange))
		genreQuery := "("
		for i, genre := range genreRange {
			if i > 0 {
				genreQuery += " OR "
			}
			genreQuery += "genres LIKE ?"
			genreParams[i] = "%" + genre + "%"
		}
		genreQuery += ")"
		query = query.Where(genreQuery, genreParams...)
	}
	if year != nil {
		query = query.Where("year = ?", *year)
	}
	if rating != nil {
		query = query.Where("rate >= ?", *rating)
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, fmt.Errorf("error counting movies: %v", err)
	}

	if err := query.Limit(*limit).Offset(*offset).Find(&movies).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, 0, fmt.Errorf("error retrieving movies: %v", err)
		}
	}

	cachedData = CachedMovies{Movies: movies, TotalCount: totalCount}
	if err := r.RedisClient.SetCache(context.Background(), cacheKey, cachedData, cacheExpiration); err != nil {
		log.Println("Error caching result:", err)
	}

	return movies, totalCount, nil
}

func (r *Repository) GetShows(limit, offset *int, genreRange []string, year *int, rating *float64) ([]*Show, int, error) {
	cacheKey := fmt.Sprintf("GetShows_%v_%v_%v_%v_%v", *limit, *offset, genreRange, nilCheck(year), nilCheck(rating))

	type CachedShows struct {
		Shows      []*Show
		TotalCount int
	}

	var cachedData CachedShows
	if _, err := r.RedisClient.GetCache(context.Background(), cacheKey, &cachedData); err == nil {
		return cachedData.Shows, cachedData.TotalCount, nil
	}

	var shows []*Show
	var totalCount int
	query := r.DB.Model(&Show{})

	if len(genreRange) > 0 {
		genreParams := make([]interface{}, len(genreRange))
		genreQuery := "("
		for i, genre := range genreRange {
			if i > 0 {
				genreQuery += " OR "
			}
			genreQuery += "genres LIKE ?"
			genreParams[i] = "%" + genre + "%"
		}
		genreQuery += ")"
		query = query.Where(genreQuery, genreParams...)
	}
	if year != nil {
		query = query.Where("year = ?", *year)
	}
	if rating != nil {
		query = query.Where("rate >= ?", *rating)
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, fmt.Errorf("error counting shows: %v", err)
	}

	if err := query.Limit(*limit).Offset(*offset).Find(&shows).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, 0, fmt.Errorf("error retrieving shows: %v", err)
		}
	}

	cachedData = CachedShows{Shows: shows, TotalCount: totalCount}
	if err := r.RedisClient.SetCache(context.Background(), cacheKey, cachedData, cacheExpiration); err != nil {
		log.Println("Error caching result:", err)
	}

	return shows, totalCount, nil
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

func (r *Repository) GetShowByID(id uint) (*Show, error) {
	var show Show

	if err := r.DB.First(&show, id).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, fmt.Errorf("error retrieving show by ID: %v", err)
		}
		return nil, nil
	}

	return &show, nil
}

func (r *Repository) GetRandomMovies(count *int, genreRange []string, year *int, rating *float64) ([]*Movie, error) {
	var movies []*Movie
	query := r.DB

	if len(genreRange) > 0 {
		genreParams := make([]interface{}, len(genreRange))
		genreQuery := "("
		for i, genre := range genreRange {
			if i > 0 {
				genreQuery += " OR "
			}
			genreQuery += "genres LIKE ?"
			genreParams[i] = "%" + genre + "%"
		}
		genreQuery += ")"
		query = query.Where(genreQuery, genreParams...)
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

func (r *Repository) GetRandomShows(count *int, genreRange []string, year *int, rating *float64) ([]*Show, error) {
	var shows []*Show
	query := r.DB

	if len(genreRange) > 0 {
		genreParams := make([]interface{}, len(genreRange))
		genreQuery := "("
		for i, genre := range genreRange {
			if i > 0 {
				genreQuery += " OR "
			}
			genreQuery += "genres LIKE ?"
			genreParams[i] = "%" + genre + "%"
		}
		genreQuery += ")"
		query = query.Where(genreQuery, genreParams...)
	}
	if year != nil {
		query = query.Where("year = ?", *year)
	}
	if rating != nil {
		query = query.Where("rate >= ?", *rating)
	}

	if err := query.Order("RANDOM()").Limit(*count).Find(&shows).Error; err != nil {
		return nil, fmt.Errorf("error retrieving random shows: %v", err)
	}

	return shows, nil
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
