package movie

import (
	"compress/gzip"
	"encoding/json"
	"flick_finder/internal/types"
	"flick_finder/pkg"
	"fmt"
	"io"
	"os"
)

// -------------------------------------
// Функиция получения популярных фильмов
// -------------------------------------
func PopularMovie(page int) ([]types.Movie, error) {
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

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, fmt.Errorf("error decode movies %s", err)
	}

	// Пробегаемся по массиву и ищем page == page(переданный)
	for _, movie := range movies {
		if movie.Page == page {
			response = append(response, movie.Results...)
			break
		}
	}

	// Если страница не найдена
	if len(response) == 0 {
		return nil, fmt.Errorf("page %d not found", page)
	}

	return pkg.TruncateArrayMovies(response), nil
}
