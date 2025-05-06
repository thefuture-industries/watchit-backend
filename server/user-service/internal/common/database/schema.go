package database

import (
	"database/sql"
	"time"
)

type Users struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	UUID         string         `gorm:"unique;size:255;not null" json:"uuid"`
	Username     string         `gorm:"unique;size:20;not null" json:"username"`
	Email        sql.NullString `gorm:"unique;type:varchar(100);default:null" json:"email"`
	Password     string         `gorm:"size:50;not null" json:"password"`
	IPAddress    string         `gorm:"size:15;not null" json:"ip_address"`
	Country      string         `gorm:"size:50;not null" json:"country"`
	Subscription bool           `gorm:"default:false;not null" json:"subscription"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type Recommendations struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UUID      string    `gorm:"unique;size:255;not null" json:"uuid"`
	GenreID   uint      `gorm:"not null" json:"genre_id"`
	Count     uint      `gorm:"not null" json:"count"`
	CreatedAt time.Time `json:"created_at"`
}

type Favourites struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UUID        string    `gorm:"unique;size:255;not null" json:"uuid"`
	MovieID     uint      `gorm:"not null" json:"movie_id"`
	MoviePoster string    `gorm:"size:255;not null" json:"movie_poster"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Limiter struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	UUID         string    `gorm:"unique;size:255;not null" json:"uuid"`
	TextLimit    uint      `gorm:"default:3;not null;" json:"text_limit"`
	YoutubeLimit uint      `gorm:"default:2;not null" json:"youtube_limit"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Comments struct{}

type Payments struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UUID      string    `gorm:"unique;size:255;not null" json:"uuid"`
	Email     string    `gorm:"size:100;not null" json:"email"`
	Card      string    `gorm:"size:50;not null" json:"card"`
	CardEnd   string    `gorm:"size:3;not null" json:"card_end"`
	Total     float64   `gorm:"size:5;not null" json:"total"`
	PaymentAt string    `gorm:"size:255;not null" json:"payment_at"`
	EndingAt  string    `gorm:"size:255;not null" json:"ending_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
