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
		Actors:      dbMovie.Actors,
		CreatedAt:   dbMovie.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   dbMovie.UpdatedAt.Format(time.RFC3339),
	}
}

func dbShowToGraphQL(dbShow *database.Show) *Show {
	return &Show{
		ID:          strconv.FormatUint(uint64(dbShow.ID), 10),
		Title:       dbShow.Title,
		Rate:        dbShow.Rate,
		Year:        dbShow.Year,
		Description: dbShow.Description,
		Genres:      dbShow.Genres,
		ImageURL:    dbShow.ImageURL,
		Actors:      dbShow.Actors,
		CreatedAt:   dbShow.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   dbShow.UpdatedAt.Format(time.RFC3339),
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

func dbShowsToGraphQL(dbShows []*database.Show) []*Show {
	var gqlShows []*Show
	for _, dbShow := range dbShows {
		gqlShow := dbShowToGraphQL(dbShow)
		gqlShows = append(gqlShows, gqlShow)
	}
	return gqlShows
}

func (r *queryResolver) SearchMoviesByKeyword(ctx context.Context, keyword string) (*SearchResults, error) {
	dbMovies, dbShows, err := r.MovieService.SearchMoviesByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	return &SearchResults{
		Movies: dbMoviesToGraphQL(dbMovies),
		Shows:  dbShowsToGraphQL(dbShows),
	}, nil
}

func (r *queryResolver) GetMovies(ctx context.Context, limit *int, offset *int, genreRange []string, year *int, rating *float64) (*GetMoviesResponse, error) {
	dbMovies, totalCount, err := r.MovieService.GetMovies(limit, offset, genreRange, year, rating)
	if err != nil {
		return nil, err
	}

	return &GetMoviesResponse{
		Movies:     dbMoviesToGraphQL(dbMovies),
		TotalCount: totalCount,
	}, nil
}

func (r *queryResolver) GetShows(ctx context.Context, limit *int, offset *int, genreRange []string, year *int, rating *float64) (*GetShowsResponse, error) {
	dbShows, totalCount, err := r.MovieService.GetShows(limit, offset, genreRange, year, rating)
	if err != nil {
		return nil, err
	}

	return &GetShowsResponse{
		Shows:      dbShowsToGraphQL(dbShows),
		TotalCount: totalCount,
	}, nil
}

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

func (r *queryResolver) GetShowByID(ctx context.Context, id string) (*Show, error) {
	showID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	show, err := r.MovieService.GetShowByID(uint(showID))
	if err != nil {
		return nil, err
	}

	return dbShowToGraphQL(show), nil
}

func (r *queryResolver) GetRandomMovies(ctx context.Context, count *int, genreRange []string, year *int, rating *float64) ([]*Movie, error) {
	dbMovies, err := r.MovieService.GetRandomMovies(count, genreRange, year, rating)
	if err != nil {
		return nil, err
	}

	return dbMoviesToGraphQL(dbMovies), nil
}

func (r *queryResolver) GetRandomShows(ctx context.Context, count *int, genreRange []string, year *int, rating *float64) ([]*Show, error) {
	dbShows, err := r.MovieService.GetRandomShows(count, genreRange, year, rating)
	if err != nil {
		return nil, err
	}

	return dbShowsToGraphQL(dbShows), nil
}

func (r *queryResolver) GetHomePageData(ctx context.Context) (*MoviesOverview, error) {
	latestMovies, err := r.MovieService.GetLatestMovies()
	if err != nil {
		return nil, err
	}

	featuredMovies, err := r.MovieService.GetFeaturedMovies()
	if err != nil {
		return nil, err
	}

	movieOfTheDay, err := r.MovieService.GetMovieOfTheDay()
	if err != nil {
		return nil, err
	}

	return &MoviesOverview{
		LatestMovies:   dbMoviesToGraphQL(latestMovies),
		FeaturedMovies: dbMoviesToGraphQL(featuredMovies),
		MovieOfTheDay:  dbMovieToGraphQL(movieOfTheDay),
	}, nil
}

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
