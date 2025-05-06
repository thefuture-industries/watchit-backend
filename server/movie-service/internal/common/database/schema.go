package database

import "gorm.io/gorm"

type Recommendations struct {
	gorm.Model
	UUID  string `gorm:"unique;size:255;not null" json:"uuid"`
	Title string `gorm:"size:255;not null" json:"title"`
	Genre string `gorm:"size:255;not null" json:"genre"`
}

type Limiter struct {
	gorm.Model
	UUID         string `gorm:"unique;size:255;not null" json:"uuid"`
	TextLimit    uint   `gorm:"default:3;not null;" json:"text_limit"`
	YoutubeLimit uint   `gorm:"default:2;not null" json:"youtube_limit"`
}
type Favourites struct {
	gorm.Model
	UUID        string `gorm:"unique;size:255;not null" json:"uuid"`
	MovieID     uint   `gorm:"not null" json:"movie_id"`
	MoviePoster string `gorm:"size:255;not null" json:"movie_poster"`
}
