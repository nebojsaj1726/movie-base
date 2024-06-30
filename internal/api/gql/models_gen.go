// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

type GetMoviesResponse struct {
	Movies     []*Movie `json:"movies"`
	TotalCount int      `json:"totalCount"`
}

type GetShowsResponse struct {
	Shows      []*Show `json:"shows"`
	TotalCount int     `json:"totalCount"`
}

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rate        string `json:"rate"`
	Year        string `json:"year"`
	Description string `json:"description"`
	Genres      string `json:"genres"`
	Duration    string `json:"duration"`
	ImageURL    string `json:"imageURL"`
	Actors      string `json:"actors"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type MoviesOverview struct {
	LatestMovies   []*Movie `json:"latestMovies"`
	FeaturedMovies []*Movie `json:"featuredMovies"`
	MovieOfTheDay  *Movie   `json:"movieOfTheDay"`
}

type Query struct {
}

type SearchResults struct {
	Movies []*Movie `json:"movies"`
	Shows  []*Show  `json:"shows"`
}

type Show struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rate        string `json:"rate"`
	Year        string `json:"year"`
	Description string `json:"description"`
	Genres      string `json:"genres"`
	ImageURL    string `json:"imageURL"`
	Actors      string `json:"actors"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
