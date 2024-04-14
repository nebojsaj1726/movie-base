package services

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nebojsaj1726/movie-base/internal/database"
)

const (
	latestMoviesCacheKey   = "latest_movies"
	featuredMoviesCacheKey = "featured_movies"
	movieOfTheDayCacheKey  = "movie_of_the_day"
	cacheExpiration        = 24 * time.Hour
)

type MovieService struct {
	Repo       *database.Repository
	RedisCache *database.RedisClient
}

func NewMovieService(repo *database.Repository) *MovieService {
	redisClient, err := database.NewRedisClient()
	if err != nil {
		log.Fatalf("Error initializing Redis client: %v", err)
	}

	return &MovieService{Repo: repo, RedisCache: redisClient}
}

func (s *MovieService) SearchMoviesByKeyword(keyword string) ([]*database.Movie, error) {
	return s.Repo.SearchMoviesByKeyword(keyword)
}

func (s *MovieService) GetMovies(limit, offset *int, genreRange []string, year *int, rating *float64) ([]*database.Movie, int, error) {
	return s.Repo.GetMovies(limit, offset, genreRange, year, rating)
}

func (s *MovieService) GetMovieByID(id uint) (*database.Movie, error) {
	return s.Repo.GetMovieByID(id)
}

func (s *MovieService) GetRandomMovies(count *int, genreRange []string, year *int, rating *float64) ([]*database.Movie, error) {
	return s.Repo.GetRandomMovies(count, genreRange, year, rating)
}

func (s *MovieService) GetLatestMovies() ([]*database.Movie, error) {
	var cachedMovies []*database.Movie
	if result, err := s.RedisCache.GetCache(context.Background(), latestMoviesCacheKey, &cachedMovies); err == nil {
		return *result.(*[]*database.Movie), nil
	}

	limit := 30
	offset := 0
	movies, _, err := s.Repo.GetMovies(&limit, &offset, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	shuffleMovies(movies)

	err = s.RedisCache.SetCache(context.Background(), latestMoviesCacheKey, movies[:5], cacheExpiration)
	if err != nil {
		return nil, fmt.Errorf("error caching latest movies: %v", err)
	}

	return movies[:5], nil
}

func (s *MovieService) GetFeaturedMovies() ([]*database.Movie, error) {
	var cachedMovies []*database.Movie
	if result, err := s.RedisCache.GetCache(context.Background(), featuredMoviesCacheKey, &cachedMovies); err == nil {
		return *result.(*[]*database.Movie), nil
	}

	limit := 300
	offset := 0
	rating := 7.0
	movies, _, err := s.Repo.GetMovies(&limit, &offset, nil, nil, &rating)
	if err != nil {
		return nil, err
	}

	shuffleMovies(movies)

	err = s.RedisCache.SetCache(context.Background(), featuredMoviesCacheKey, movies[:5], cacheExpiration)
	if err != nil {
		return nil, fmt.Errorf("error caching featured movies: %v", err)
	}

	return movies[:5], nil
}

func (s *MovieService) GetMovieOfTheDay() (*database.Movie, error) {
	var cachedMovie *database.Movie
	if result, err := s.RedisCache.GetCache(context.Background(), movieOfTheDayCacheKey, &cachedMovie); err == nil {
		return *result.(**database.Movie), nil
	}

	count := 1
	rating := 7.0
	movies, err := s.Repo.GetRandomMovies(&count, nil, nil, &rating)
	if err != nil {
		return nil, err
	}
	if len(movies) == 0 {
		return nil, fmt.Errorf("no movie found")
	}

	movieOfTheDay := movies[0]

	err = s.RedisCache.SetCache(context.Background(), movieOfTheDayCacheKey, movieOfTheDay, cacheExpiration)
	if err != nil {
		return nil, fmt.Errorf("error caching movie of the day: %v", err)
	}

	return movieOfTheDay, nil
}

func shuffleMovies(movies []*database.Movie) {
	rand.Shuffle(len(movies), func(i, j int) {
		movies[i], movies[j] = movies[j], movies[i]
	})
}
