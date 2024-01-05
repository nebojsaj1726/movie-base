package gql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"strconv"
	"time"

	"github.com/nebojsaj1726/movie-base/internal/database/services"
)

type Resolver struct{
	MovieService *services.MovieService
}

func NewResolver(movieService *services.MovieService) *Resolver {
	return &Resolver{
		MovieService: movieService,
	}
}

// SearchMoviesByKeyword is the resolver for the searchMoviesByKeyword field.
func (r *queryResolver) SearchMoviesByKeyword(ctx context.Context, keyword string) ([]*Movie, error) {
	dbMovies, err := r.MovieService.SearchMoviesByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	var gqlMovies []*Movie
	for _, dbMovie := range dbMovies {
		gqlMovie := &Movie{
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
		gqlMovies = append(gqlMovies, gqlMovie)
	}

	return gqlMovies, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
