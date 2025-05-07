package database

import (
	"database/sql"
	"time"
)

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
