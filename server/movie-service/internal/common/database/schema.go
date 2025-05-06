package database

import "time"

type Genres struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	GenreID   uint   `gorm:"unique;not null" json:"genre_id"`
	GenreName string `gorm:"unique;not null" json:"genre_name"`
}

type Recommendations struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UUID      string    `gorm:"unique;size:255;not null" json:"uuid"`
	GenreID   uint      `gorm:"not null" json:"genre_id"`
	Count     uint      `gorm:"not null" json:"count"`
	CreatedAt time.Time `json:"created_at"`
}
