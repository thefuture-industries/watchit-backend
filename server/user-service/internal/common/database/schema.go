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
