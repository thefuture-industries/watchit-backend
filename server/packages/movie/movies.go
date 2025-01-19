package movie

import (
	"compress/gzip"
	"encoding/json"
	"flicksfi/internal/types"
	"flicksfi/packages"
	"flicksfi/packages/errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/xrash/smetrics"
)

var genres map[string]int = map[string]int{
	"Action":          28,
	"Adventure":       12,
	"Animation":       16,
	"Comedy":          35,
	"Crime":           80,
	"Documentary":     99,
	"Drama":           18,
	"Family":          10751,
	"Fantasy":         14,
	"History":         36,
	"Horror":          27,
	"Music":           10402,
	"Mystery":         9648,
	"Romance":         10749,
	"Science Fiction": 878,
	"TV Movie":        10770,
	"Thriller":        53,
	"War":             10752,
	"Western":         37,
}

// ----------------
// Получение фильмов
// ----------------
func GetMovies(parametrs map[string]string) ([]types.Movie, error) {
	// Определяем жанр
	var genre_ids int = genres[parametrs["genre_id"]]
	// Количество возвращяемых фильмов
	// var limit int = 10

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

	// Чек для проверки нашел ли пользователь фильм по Title
	isTitle := false
	// Чек для проверки нашел ли пользователь фильм по Дате
	isDate := false

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, fmt.Errorf("error decode movies %s", err)
	}

	// Пробегаемся по массиву фильмов
	for _, movie := range movies {
		// Пробегаемся по массиву данных в page
		for _, movieItem := range movie.Results {

			// Существует ли genre_id
			if parametrs["genre_id"] != "" {
				// Пробегаемся по массиву genres_ids в деталях фильма
				for _, indx := range movieItem.GenreIds {
					if indx == genre_ids {
						response = append(response, movieItem)
					}
				}
			}

			// Поиск даты фильма
			if parametrs["release_date"] != "" {
				if strings.Contains(movieItem.ReleaseDate, parametrs["release_date"]) {
					response = append(response, movieItem)
					isDate = true
				}
			}

			// Поиск фильма по Title и Overview
			if parametrs["search"] != "" {
				similarity := smetrics.JaroWinkler(strings.ToLower(parametrs["search"]), strings.ToLower(movieItem.Title), 0.1, 4)
				if similarity >= 0.7 {
					response = append(response, movieItem)
					isTitle = true
				}
			}
		}
	}

	// Если фильм не найден по Title, то ошибка
	if err := errors.SearchNotFound(parametrs, isTitle); err != nil {
		return nil, err
	}
	// Если фильм не найден по ReleaseDate, то ошибка
	if err := errors.DateNotFound(parametrs, isDate); err != nil {
		return nil, err
	}

	return pkg.ShuffleArray(response), nil
}

// ------------------------------
// Получение деталий фильмов по ID
// ------------------------------
func MovieDetails(id int) (types.Movie, error) {
	// Читаем файл (gzip)
	file, err := os.Open("pkg/movie/db/movies.json.gz")
	if err != nil {
		return types.Movie{}, fmt.Errorf("error open file")
	}
	defer file.Close()

	// Создать декомпрессор gzip
	zr, err := gzip.NewReader(file)
	if err != nil {
		return types.Movie{}, fmt.Errorf("error decompress file to movie")
	}
	defer zr.Close()

	// Читаем массив байтов
	data, err := io.ReadAll(zr)
	if err != nil {
		return types.Movie{}, fmt.Errorf("error read data movies")
	}

	// Создаем переменную для фильмов
	var movies []types.JsonMovies
	var response types.Movie

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return types.Movie{}, fmt.Errorf("error decode movies %s", err)
	}

	// Пробегаемся по массиву и ищем Id == id(переданный)
	for _, movie_data := range movies {
		for _, movieItem := range movie_data.Results {
			if movieItem.Id == id {
				response = movieItem
				break
			}
		}
	}

	return response, nil
}

func GetMovieByGenre(genre int, limit int) ([]types.Movie, error) {
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
	var count int = 0

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, fmt.Errorf("error decode movies %s", err)
	}

	rand.Seed(time.Now().UnixNano())

	indices := make([]int, len(movies))
	for i := range indices {
		indices[i] = i
	}
	rand.Shuffle(len(indices), func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	for _, i := range indices {
		movie := movies[i]

		for _, movieItem := range movie.Results {
			for _, movieGenre := range movieItem.GenreIds {
				if movieGenre == genre && count <= limit {
					response = append(response, movieItem)
					count++
				}
			}

			if count > limit {
				break
			}
		}

		if count > limit {
			break
		}
	}

	return response, nil
}
