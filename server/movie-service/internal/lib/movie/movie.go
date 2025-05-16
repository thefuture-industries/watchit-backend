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

const (
	MAX_COUNT_MOVIES uint16 = 500
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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomPage := uint16(r.Intn(int(MAX_COUNT_MOVIES)))

	movieFile, err := os.Open(constants.MOVIE_JSON_PATH_READ)
	if err != nil {
		m.logger.Error(err.Error())
		return types.Movies{}, err
	}
	defer movieFile.Close()

	var moviesJson []types.Movies
	// var movies []types.Movie

	if err := json.NewDecoder(movieFile).Decode(&moviesJson); err != nil {
		m.logger.Error(err.Error())
		return types.Movies{}, fmt.Errorf("Error get list movies!")
	}

	for _, movies := range moviesJson {
		if movies.Page == randomPage {
			return movies, nil // RETURNS MOVIES
		}
	}

	return types.Movies{}, fmt.Errorf("No movies found!")
}
