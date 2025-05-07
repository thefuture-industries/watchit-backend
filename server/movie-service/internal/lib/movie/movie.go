package movie

import (
	"encoding/json"
	"fmt"
	"go-movie-service/internal/common/constants"
	"go-movie-service/internal/lib"
	"go-movie-service/internal/types"
	"math/rand"
	"os"
	"time"
)

type Movie struct {
	logger *lib.Logger
}

func NewMovie() *Movie {
	return &Movie{
		logger: lib.NewLogger(),
	}
}

func (m *Movie) GetMovies() (types.Movies, error) {
	index := LoadIDX()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var pages []uint
	for p := range index {
		if p <= 500 {
			pages = append(pages, uint(p))
		}
	}

	if len(pages) == 0 {
		return types.Movies{}, fmt.Errorf("No movies found!")
	}

	randomPage := pages[r.Intn(len(pages))]
	offset := index[uint32(randomPage)]

	file, _ := os.Open(constants.MOVIE_JSON_PATH)
	defer file.Close()
	if _, err := file.Seek(int64(offset), 0); err != nil {
		m.logger.Error(err.Error())
		return types.Movies{}, err
	}

	decoder := json.NewDecoder(file)
	var movies types.Movies
	if err := decoder.Decode(&movies); err != nil {
		m.logger.Error(err.Error())
		return types.Movies{}, err
	}

	return movies, nil
}
