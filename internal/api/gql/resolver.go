package gql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"strconv"
	"time"

	"github.com/nebojsaj1726/movie-base/internal/database"
	"github.com/nebojsaj1726/movie-base/internal/database/services"
)

type Resolver struct {
	MovieService *services.MovieService
}

func NewResolver(movieService *services.MovieService) *Resolver {
	return &Resolver{
		MovieService: movieService,
	}
}

func dbMovieToGraphQL(dbMovie *database.Movie) *Movie {
	return &Movie{
		ID:          strconv.FormatUint(uint64(dbMovie.ID), 10),
		Title:       dbMovie.Title,
		Rate:        dbMovie.Rate,
		Year:        dbMovie.Year,
		Description: dbMovie.Description,
		Genres:      dbMovie.Genres,
		Duration:    dbMovie.Duration,
		ImageURL:    dbMovie.ImageURL,
		CreatedAt:   dbMovie.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   dbMovie.UpdatedAt.Format(time.RFC3339),
	}
}

func dbMoviesToGraphQL(dbMovies []*database.Movie) []*Movie {
	var gqlMovies []*Movie
	for _, dbMovie := range dbMovies {
		gqlMovie := dbMovieToGraphQL(dbMovie)
		gqlMovies = append(gqlMovies, gqlMovie)
	}
	return gqlMovies
}

// SearchMoviesByKeyword is the resolver for the searchMoviesByKeyword field.
func (r *queryResolver) SearchMoviesByKeyword(ctx context.Context, keyword string) ([]*Movie, error) {
	dbMovies, err := r.MovieService.SearchMoviesByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	return dbMoviesToGraphQL(dbMovies), nil
}

// GetMovies is the resolver for the getMovies field.
func (r *queryResolver) GetMovies(ctx context.Context, limit *int, offset *int, genreRange []string, year *int, rating *float64) ([]*Movie, error) {
	dbMovies, err := r.MovieService.GetMovies(limit, offset, genreRange, year, rating)
	if err != nil {
		return nil, err
	}

	return dbMoviesToGraphQL(dbMovies), nil
}

// GetMovieByID is the resolver for the getMovieByID field.
func (r *queryResolver) GetMovieByID(ctx context.Context, id string) (*Movie, error) {
	movieID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	movie, err := r.MovieService.GetMovieByID(uint(movieID))
	if err != nil {
		return nil, err
	}

	return dbMovieToGraphQL(movie), nil
}

func (r *queryResolver) GetRandomMovies(ctx context.Context, count *int, genreRange []string, year *int, rating *float64) ([]*Movie, error) {
	dbMovies, err := r.MovieService.GetRandomMovies(count, genreRange, year, rating)
	if err != nil {
		return nil, err
	}

	return dbMoviesToGraphQL(dbMovies), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
