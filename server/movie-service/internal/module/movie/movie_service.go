package movie

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"go-movie-service/internal/types"
	"io"
	"os"
	"path/filepath"
)

var MovieDetails = movieDetails

// movieDetails получает детали фильма по ID
func movieDetails(id int) (types.Movie, error) {

	filePath := filepath.Join("internal", "data", "movies.json.gz")

	file, err := os.Open(filePath)
	if err != nil {
		return types.Movie{}, fmt.Errorf("error opening file %s: %v", filePath, err)
	}
	defer file.Close()

	zr, err := gzip.NewReader(file)
	if err != nil {
		return types.Movie{}, fmt.Errorf("error creating gzip reader: %v", err)
	}
	defer zr.Close()

	data, err := io.ReadAll(zr)
	if err != nil {
		return types.Movie{}, fmt.Errorf("error reading compressed data: %v", err)
	}
	var movieData types.Movies
	var response types.Movie

	err = json.Unmarshal(data, &movieData)
	if err != nil {
		return types.Movie{}, fmt.Errorf("error decoding movies: %v", err)
	}

	for _, movieItem := range movieData.Results {
		if movieItem.Id == id {
			response = movieItem
			return response, nil
		}
	}

	return types.Movie{}, nil
}
