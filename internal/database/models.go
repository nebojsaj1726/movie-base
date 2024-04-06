package database

import "time"

type Movie struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Rate        string    `json:"rate"`
	Year        string    `json:"year"`
	Description string    `json:"description"`
	Genres      string    `json:"genres"`
	Duration    string    `json:"duration"`
	ImageURL    string    `json:"image_url"`
	Actors      string    `json:"actors"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Movie) TableName() string {
	return "movies"
}

type Show struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Rate        string    `json:"rate"`
	Year        string    `json:"year"`
	Description string    `json:"description"`
	Genres      string    `json:"genres"`
	ImageURL    string    `json:"image_url"`
	Actors      string    `json:"actors"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Show) TableName() string {
	return "shows"
}
