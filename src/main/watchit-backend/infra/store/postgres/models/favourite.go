package models

import "time"

type Favourite struct {
	ID          int64     `db:"id" json:"id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	UserUUID    string    `json:"user_uuid"`
	MovieId     int       `json:"movie_id"`
	MoviePoster string    `json:"movie_poster"`
}
