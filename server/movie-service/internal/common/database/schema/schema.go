package schema

import (
	"database/sql"
	"time"
)

type Users struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	UUID         string         `gorm:"unique;type:varchar(255);not null" json:"uuid"`
	Username     string         `gorm:"unique;type:varchar(20);not null" json:"username"`
	Email        sql.NullString `gorm:"unique;type:varchar(100);default:null" json:"email"`
	Password     string         `gorm:"type:varchar(50);not null" json:"password"`
	IPAddress    string         `gorm:"type:varchar(15);not null" json:"ip_address"`
	Country      string         `gorm:"type:varchar(50);not null" json:"country"`
	Subscription bool           `gorm:"default:false;not null" json:"subscription"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type Profiles struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UUID      string    `gorm:"unique;type:varchar(255);not null" json:"uuid"`
	Bio       string    `gorm:"type:varchar(255)" json:"bio"`
	Avatar    byte      `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Payments struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UUID        string    `gorm:"unique;type:varchar(255);not null" json:"uuid"`
	Amount      float64   `gorm:"not null" json:"amount"`
	PaymentDate time.Time `gorm:"not null" json:"payment_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaymentCards struct {
	ID               uint      `gorm:"primarykey" json:"id"`
	UUID             string    `gorm:"unique;type:varchar(255);not null" json:"uuid"`
	CardNumber       string    `gorm:"type:varchar(16);not null" json:"card_number"`
	ExpirationMounth int16     `gorm:"not null;check:expiration_month >= 1 AND expiration_month <= 12" json:"expiration_mounth"`
	ExpirationYear   int16     `gorm:"not null" json:"expiration_year"`
	CVVCode          string    `gorm:"type:varchar(3);not null" json:"cvv_code"`
	CardHolderName   string    `gorm:"type:varchar(100);not null" json:"cardholder_name"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Limiter struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UUID      string    `gorm:"unique;type:varchar(255);not null" json:"uuid"`
	TextLimit uint      `gorm:"default:3;not null;" json:"text_limit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Recommendations struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UUID      string    `gorm:"type:varchar(255);not null" json:"uuid"`
	GenreID   uint      `gorm:"not null" json:"genre_id"`
	Count     uint      `gorm:"not null" json:"count"`
	CreatedAt time.Time `json:"created_at"`
}

type Favourites struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UUID        string    `gorm:"unique;type:varchar(255);not null" json:"uuid"`
	MovieID     uint      `gorm:"not null" json:"movie_id"`
	MoviePoster string    `gorm:"type:varchar(255);not null" json:"movie_poster"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Genres struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	GenreID   uint   `gorm:"unique;not null" json:"genre_id"`
	GenreName string `gorm:"unique;not null" json:"genre_name"`
}
