package text

import (
	"compress/gzip"
	"encoding/json"
	"flick_finder/internal/types"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Стоп-слова
var stopWords []string = []string{"a", "on", "is", "the", "are", "in", "of", "to", "and", "an"}

func OverviewText(text string) ([]types.Movie, error) { // This is word -> ["This", "is", "word"] -> ["This", "word"] -> ["this", "word"]
	// Чтение файла
	file, err := os.Open("pkg/movie/db/movies.json.gz")
	if err != nil {
		return nil, fmt.Errorf("error reading movies")
	}
	defer file.Close()

	// Открытие zip файл
	zr, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("error reading movies")
	}
	defer zr.Close()

	// Чтение файла zip
	data, err := io.ReadAll(zr)
	if err != nil {
		return nil, fmt.Errorf("error reading movies")
	}

	// Конвертация данных файла в json
	var movies []types.JsonMovies
	var response []types.Movie

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, fmt.Errorf("error convert data to json")
	}

	re := regexp.MustCompile(`\b\w+\b`)
	array_text := re.FindAllString(text, -1) // ["This", "is", "word"]

	filterArrayText := []string{} // ["this", "word"]
	for _, txt := range array_text {
		txt = strings.ToLower(txt)
		isStop := false

		for _, word := range stopWords {
			if txt == word {
				isStop = true
				break
			}

			if !isStop {
				filterArrayText = append(filterArrayText, txt)
			}
		}
	}

	for _, movie := range movies {
		for _, movieItem := range movie.Results {
			for _, txt := range filterArrayText {
				if movieItem.Overview == txt {
					response = append(response, movieItem)
					break
				}
			}
		}
	}

	return response, nil
}
