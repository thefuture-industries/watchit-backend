package movie

import (
	"encoding/json"
	"fmt"
	"go-movie-service/internal/common/constants"
	"go-movie-service/internal/common/utils"
	"go-movie-service/internal/lib"
	"go-movie-service/internal/lib/machinelearning"
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
	logger        *lib.Logger
	lsaBuilder    *machinelearning.LSABuilder
	moviesByCache []types.Movies
}

func NewMovie() *Movie {
	return &Movie{
		logger:        lib.NewLogger(),
		lsaBuilder:    machinelearning.NewLSABuilder(),
		moviesByCache: nil,
	}
}

func (m *Movie) GetMovies() ([]types.Movie, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomPage := uint16(r.Intn(int(MAX_COUNT_MOVIES)))

	movieFile, err := os.Open(constants.MOVIE_JSON_PATH_READ)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}
	defer movieFile.Close()

	offset := PIDX(randomPage)

	_, errSeek := movieFile.Seek(int64(offset), io.SeekStart)
	if errSeek != nil {
		return nil, fmt.Errorf("error getting movies by %d", randomPage)
	}

	var moviesJson types.Movies
	decoder := utils.JSON.NewDecoder(movieFile)
	if err := decoder.Decode(&moviesJson); err != nil {
		return nil, fmt.Errorf("we didn't find any movies.")
	}

	return moviesJson.Results, nil
}

func (m *Movie) GetDetailsMovies(id uint32) (types.Movie, error) {
	if m.moviesByCache == nil {
		movieFile, err := os.Open(constants.MOVIE_JSON_PATH_READ)
		if err != nil {
			m.logger.Error(err.Error())
			return types.Movie{}, err
		}
		defer movieFile.Close()

		if err := json.NewDecoder(movieFile).Decode(&m.moviesByCache); err != nil {
			m.logger.Error(err.Error())
			return types.Movie{}, fmt.Errorf("error get list movies!")
		}
	}

	for _, movies := range m.moviesByCache {
		for _, movie := range movies.Results {
			if movie.Id == id {
				return movie, nil
			}
		}
	}

	return types.Movie{}, fmt.Errorf("we didn't find any movies with id: %d", id)
}

func (m *Movie) GetMoviesByText(textInput string) ([]types.Movie, error) {
	if m.moviesByCache == nil {
		movieFile, err := os.Open(constants.MOVIE_JSON_PATH_READ)
		if err != nil {
			m.logger.Error(err.Error())
			return nil, err
		}
		defer movieFile.Close()

		if err := json.NewDecoder(movieFile).Decode(&m.moviesByCache); err != nil {
			m.logger.Error(err.Error())
			return nil, fmt.Errorf("error get list movies!")
		}
	}

	var movieList []types.Movie

	for _, movies := range m.moviesByCache {
		for _, movie := range movies.Results {
			movieList = append(movieList, movie)
		}
	}

	matrix, docs := m.lsaBuilder.AnalyzeByMovie(movieList, textInput)
	if matrix == nil {
		return nil, fmt.Errorf("we didn't find any movies.")
	}

	rows, _ := matrix.Dims()
	inputVec := matrix.RawRowView(rows - 1)

	sims := make([]docSim, 0, rows-1)
}
