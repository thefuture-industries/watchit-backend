package movie

import (
	"compress/gzip"
	"encoding/json"
	"flicksfi/internal/types"
	"flicksfi/pkg"
	"fmt"
	"io"
	"os"
)

func Similar(movie_data map[string]interface{}) ([]types.Movie, error) {
	// Читаем файл (gzip)
	file, err := os.Open("pkg/movie/db/movies.json.gz")
	if err != nil {
		return nil, fmt.Errorf("error open file")
	}
	defer file.Close()

	// Создать декомпрессор gzip
	zr, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("error decompress file to movie")
	}
	defer zr.Close()

	// Читаем массив байтов
	data, err := io.ReadAll(zr)
	if err != nil {
		return nil, fmt.Errorf("error read data movies")
	}

	// Создаем переменную для фильмов
	var movies []types.JsonMovies
	var response []types.Movie

	// Из json в массив фильмов
	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, fmt.Errorf("error decode movies %s", err)
	}

	for _, movie := range movies {
		for _, movieItem := range movie.Results {
			tfidf := pkg.TF_IDF_MOVIE(movieItem.Overview, movie_data["overview"].(string), 0.189)
			fmt.Printf("tfidf для %s: %.4f\n", movieItem.Title, tfidf)

			if tfidf >= 0.54 {
				response = append(response, movieItem)
			}
		}
	}

	return pkg.TruncateArrayMovies(response), nil
}
