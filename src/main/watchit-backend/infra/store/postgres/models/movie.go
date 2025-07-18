package models

import "time"

type Movie struct {
	ID               int64      `db:"id" json:"id"`
	CreatedAt        time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at" json:"updated_at"`
	Title            string     `db:"title" json:"title"`
	OriginalTitle    *string    `db:"original_title" json:"original_title"`
	Overview         *string    `db:"overview" json:"overview"`
	ReleaseDate      *time.Time `db:"release_date" json:"release_date"`
	OriginalLanguage string     `db:"original_language" json:"original_language"`
	Popularity       *float32   `db:"popularity" json:"popularity"`
	VoteAverage      *float32   `db:"vote_average" json:"vote_average"`
	VoteCount        *int       `db:"vote_count" json:"vote_count"`
	PosterPath       *string    `db:"poster_path" json:"poster_path"`
	BackdropPath     *string    `db:"backdrop_path" json:"backdrop_path"`
	Video            bool       `db:"video" json:"video"`
	Adult            bool       `db:"adult" json:"adult"`
}

type MovieWithGenres struct {
	Movie
	Genres []string `json:"genres"`
}
