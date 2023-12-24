package tests

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nebojsaj1726/movie-base/internal/database"
	"github.com/nebojsaj1726/movie-base/internal/scraper"
	"github.com/stretchr/testify/assert"
)

var mockMovies = []scraper.Movie{
	{
		Title:       "Mock Movie 1",
		Rate:        "8.0",
		Year:        "2021",
		Description: "A mock movie description.",
		Duration:    "2h",
		ImageURL:    "http://mock-movie-1.jpg",
		Genres:      []string{"Action", "Adventure"},
	},
	{
		Title:       "Mock Movie 2",
		Rate:        "7.5",
		Year:        "2020",
		Description: "Another mock movie description.",
		Duration:    "1h 30m",
		ImageURL:    "http://mock-movie-2.jpg",
		Genres:      []string{"Comedy", "Drama"},
	},
}

var mockShows = []scraper.Movie{
	{
		Title:       "Mock Show 1",
		Rate:        "9.0",
		Year:        "2019",
		Description: "A mock show description.",
		ImageURL:    "http://mock-show-1.jpg",
		Genres:      []string{"Sci-Fi", "Drama"},
	},
	{
		Title:       "Mock Show 2",
		Rate:        "8.5",
		Year:        "2018",
		Description: "Another mock show description.",
		ImageURL:    "http://mock-show-2.jpg",
		Genres:      []string{"Mystery", "Thriller"},
	},
}

func setupTestDatabase(t *testing.T) (*database.Repository, *gorm.DB) {
	// Use SQLite in-memory for testing
	db, err := gorm.Open("sqlite3", "file::memory:?cache=shared")
	assert.Nil(t, err, "Error initializing the database for testing")


	if err := db.AutoMigrate(&database.Movie{}).Error; err != nil {
		assert.FailNow(t, "Error migrating movies table for testing")
	}

	if err := db.AutoMigrate(&database.Show{}).Error; err != nil {
		assert.FailNow(t, "Error migrating shows table for testing")
	}

	repo := &database.Repository{DB: db}
	return repo, db
}

func TestCreateMovies(t *testing.T) {
	repo, db := setupTestDatabase(t)
	defer db.Close()

	err := repo.CreateMovies(mockMovies)
	assert.Nil(t, err, "Error creating movies in the database")

	var movies []database.Movie
	db.Find(&movies)

	assert.Len(t, movies, len(mockMovies), "Number of movies in the database doesn't match")

	for i, mockMovie := range mockMovies {
		assert.Equal(t, mockMovie.Title, movies[i].Title, "Movie title mismatch")
	}
}

func TestCreateShows(t *testing.T) {
	repo, db := setupTestDatabase(t)
	defer db.Close()

	err := repo.CreateShows(mockShows)
	assert.Nil(t, err, "Error creating shows in the database")

	var shows []database.Show
	db.Find(&shows)

	assert.Len(t, shows, len(mockShows), "Number of shows in the database doesn't match")

	for i, mockShow := range mockShows {
		assert.Equal(t, mockShow.Title, shows[i].Title, "Show title mismatch")
	}
}