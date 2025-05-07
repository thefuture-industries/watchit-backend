package movie

import (
	"math/rand"
	"time"
)

type Movie struct{}

func NewMovie() *Movie {
	return &Movie{}
}

func (m *Movie) GetMovies() {
	index := LoadIDX()

	rand.Seed(time.Now().UnixNano())
	var pages []uint
	for p := range index {
		if p <= 500 {
			pages = append(pages, uint(p))
		}
	}
}
