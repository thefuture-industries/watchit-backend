package movie

import (
	"encoding/json"
	"fmt"
	"go-movie-service/internal/common/constants"
	"go-movie-service/internal/types"
	"math/rand"
	"os"
	"time"
)

type Movie struct{}

func NewMovie() *Movie {
	return &Movie{}
}

func (m *Movie) GetMovies() (types.Movies, error) {
	index := LoadIDX()

	rand.Seed(time.Now().UnixNano())
	var pages []uint
	for p := range index {
		if p <= 500 {
			pages = append(pages, uint(p))
		}
	}

	if len(pages) == 0 {
		return types.Movies{}, fmt.Errorf("No movies found!")
	}

	randomPage := pages[rand.Uint32()%uint32(len(pages))]
	offset := index[uint32(randomPage)]

	file, _ := os.Open(constants.MOVIE_JSON_PATH)
	defer file.Close()
	file.Seek(int64(offset), 0)

	decoder := json.NewDecoder(file)
	var movies types.Movies
	decoder.Decode(&movies)

	return movies, nil
}
