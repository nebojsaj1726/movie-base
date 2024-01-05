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
