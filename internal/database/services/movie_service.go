package services

import (
	"github.com/nebojsaj1726/movie-base/internal/database"
)

type MovieService struct {
	Repo *database.Repository
}

func NewMovieService(repo *database.Repository) *MovieService {
	return &MovieService{Repo: repo}
}

func (s *MovieService) SearchMoviesByKeyword(keyword string) ([]*database.Movie, error) {
	return s.Repo.SearchMoviesByKeyword(keyword)
}

func (s *MovieService) GetMovies(limit, offset *int, genreRange []string, year *int, rating *float64) ([]*database.Movie, error) {
	return s.Repo.GetMovies(limit, offset, genreRange, year, rating)
}

func (s *MovieService) GetMovieByID(id uint) (*database.Movie, error) {
	return s.Repo.GetMovieByID(id)
}

func (s *MovieService) GetRandomMovies(count *int, genreRange []string, year *int, rating *float64) ([]*database.Movie, error) {
	return s.Repo.GetRandomMovies(count, genreRange, year, rating)
}
