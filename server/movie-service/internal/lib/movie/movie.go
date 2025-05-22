package movie

import (
	"encoding/json"
	"fmt"
	"go-movie-service/internal/common/constants"
	"go-movie-service/internal/lib"
	"go-movie-service/internal/types"
	"io"
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

	// var moviesJson []types.Movies
	var moviesJson types.Movies
	// var movies []types.Movie

	// if err := json.NewDecoder(movieFile).Decode(&moviesJson); err != nil {
	// 	m.logger.Error(err.Error())
	// 	return types.Movies{}, fmt.Errorf("error get list movies!")
	// }

	offset := PIDX(randomPage)

	_, err := movieFile.Seek(offset, io.SeekStart)
	if err != nil {
		return 
	}

	for _, movies := range moviesJson {
		if movies.Page == randomPage {
			return movies, nil // RETURNS MOVIES
		}
	}

	return types.Movies{}, fmt.Errorf("we didn't find any movies.")
}

func (m *Movie) GetDetailsMovies(id uint32) (types.Movie, error) {
	movieFile, err := os.Open(constants.MOVIE_JSON_PATH_READ)
	if err != nil {
		m.logger.Error(err.Error())
		return types.Movie{}, err
	}
	defer movieFile.Close()

	var moviesJson []types.Movies

	if err := json.NewDecoder(movieFile).Decode(&moviesJson); err != nil {
		m.logger.Error(err.Error())
		return types.Movie{}, fmt.Errorf("error get list movies!")
	}

	for _, movies := range moviesJson {
		for _, movie := range movies.Results {
			if movie.Id == id {
				return movie, nil
			}
		}
	}

	return types.Movie{}, fmt.Errorf("we didn't find any movies with id: %d", id)
}
