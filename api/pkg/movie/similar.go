package movie

import (
	"compress/gzip"
	"encoding/json"
	"flicksfi/internal/types"
	"flicksfi/pkg"
	"flicksfi/pkg/machinelearning"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
)

var genres_by_id map[int]string = map[int]string{
	28:    "Action",
	12:    "Adventure",
	16:    "Animation",
	35:    "Comedy",
	80:    "Crime",
	99:    "Documentary",
	18:    "Drama",
	10751: "Family",
	14:    "Fantasy",
	36:    "History",
	27:    "Horror",
	10402: "Music",
	9648:  "Mystery",
	10749: "Romance",
	878:   "Science Fiction",
	10770: "TV Movie",
	53:    "Thriller",
	10752: "War",
	37:    "Western",
}

func Similar(movie_data map[string]interface{}) ([]types.Movie, error) {
	nums_worker := runtime.NumCPU()
	moviesChan := make(chan types.Movie)
	resultsChan := make(chan types.Movie)
	var wg sync.WaitGroup

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

	// Из json в массив фильмов
	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, fmt.Errorf("error decode movies %s", err)
	}

	// Запуск горутин для обработки фильмов
	for i := 0; i < nums_worker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var genres_names []string
			for _, genre := range movie_data["genre_id"].([]int) {
				if genre, exists := genres_by_id[genre]; exists {
					genres_names = append(genres_names, genre)
				}
			}

			genres_names_str := strings.Join(genres_names, " ")

			// Вычисление схожести фильмов
			for movie := range moviesChan {
				similarity := machinelearning.TF_IDF(movie.Overview, movie_data["overview"].(string)+genres_names_str, 0.189)
				// fmt.Printf("tfidf для %s: %.4f\n", movie.Title, similarity)

				if similarity >= float64(0.59) {
					resultsChan <- movie
				}
			}
		}()
	}

	// Отправка фильмов в канал
	go func() {
		for _, movie := range movies {
			for _, movieItem := range movie.Results {
				moviesChan <- movieItem
			}
		}
		close(moviesChan)
	}()

	// Ожидания завершения всех горутин
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Сбор результатов
	var results []types.Movie
	for movie := range resultsChan {
		results = append(results, movie)
	}

	return pkg.ShuffleArray(results), nil
}
