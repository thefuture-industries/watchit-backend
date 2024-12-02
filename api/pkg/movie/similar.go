package movie

import (
	"compress/gzip"
	"encoding/json"
	"flick_finder/internal/types"
	"flick_finder/pkg"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var genreMap map[int]string = map[int]string{
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

// Стоп-слова
var stopWords []string = []string{"a", "on", "is", "the", "are", "in", "of", "to", "and", "an"}
var stopNumbers []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

// Похожие фильмы
// --------------
func Similar(movie_data map[string]interface{}) ([]types.Movie, error) {
	// Получаем элементы из interface{}
	// genre_ids := movie_data["genre_id"].(string)

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

	// Конвертация из string в int
	// genreINT, err := strconv.Atoi(movie_data["genre_id"])
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Разбиваем на слова title и overview
	re := regexp.MustCompile(`\b\w+\b`)
	overviewWORDS := re.FindAllString(movie_data["overview"].(string), -1)
	titleWORDS := re.FindAllString(movie_data["title"].(string), -1)

	// Удаляем стоп-слова
	// Переменная отфильтрованная
	overviewFILTER := []string{}
	titleFILTER := []string{}

	// Фильтрование предложения OVERVIEW
	for _, overviewWORD := range overviewWORDS {
		overviewWORD = strings.ToLower(overviewWORD)
		isStopWord := false
		for _, stopWord := range stopWords {
			if overviewWORD == stopWord {
				isStopWord = true
				break
			}
		}

		if !isStopWord {
			overviewFILTER = append(overviewFILTER, overviewWORD)
		}
	}

	// Фильтрование предложения TITLE
	for _, titleWORD := range titleWORDS {
		titleWORD = strings.ToLower(titleWORD)
		isStopWord := false
		for _, stopNumber := range stopNumbers {
			if titleWORD == stopNumber {
				isStopWord = true
				break
			}
		}

		if !isStopWord {
			titleFILTER = append(titleFILTER, titleWORD)
		}
	}

	// Поиск схожих слов из overview и часто-популярных слов в JSON
	overviewSimilars, err := findSimilarWords(movie_data["genre_id"].([]int), overviewFILTER)
	if err != nil {
		return nil, err
	}

	// Поиск схожих слов из overview и часто-популярных слов в JSON
	titleSimilars, err := findSimilarWords(movie_data["genre_id"].([]int), titleFILTER)
	if err != nil {
		return nil, err
	}

	for _, movie := range movies {
		for _, movieItem := range movie.Results {
			found := false // Флаг, указывающий на наличие совпадения
			// Похожие по overview
			for _, overviewSimilar := range overviewSimilars {
				if strings.Contains(strings.ToLower(movieItem.Overview), overviewSimilar) {
					found = true
					break
				}
			}

			if found {
				response = append(response, movieItem)
				continue
			}

			// Похожие по title
			for _, title := range titleFILTER {
				if strings.ToLower(movieItem.Title) == title {
					response = append(response, movieItem)
					break
				}
			}

			// Похожие по ключевым словам words.json
			for _, titleSimilar := range titleSimilars {
				if strings.Contains(strings.ToLower(movieItem.Title), titleSimilar) {
					response = append(response, movieItem)
					break
				}
			}
		}
	}

	return pkg.TruncateArrayMovies(response), nil
}

// Поиск ключевых слов
// ------------------- // genreID int
func findSimilarWords(genreIDs []int, wordsToFind []string) ([]string, error) {
	var keywordsList types.KeywordList

	// Читаем файл
	file, err := os.Open("pkg/movie/db/temp/words.json")
	if err != nil {
		return nil, fmt.Errorf("error open file")
	}
	defer file.Close()

	// Читаем массив байтов
	words, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error read data")
	}

	// Из json в массив слов
	err = json.Unmarshal(words, &keywordsList)
	if err != nil {
		return nil, fmt.Errorf("error decode data %s", err)
	}

	// Инициализируем пустой срез
	foundWords := make([]string, 0)

	// Получаем карту категорий и их слов ВСЕХ
	keywordMap, err := findPopularWordsMap(keywordsList)
	if err != nil {
		return nil, err
	}

	// Ищем название категории по её ID
	var genreNames = make([]string, len(genreIDs))
	for i, genreID := range genreIDs {
		genreNames[i] = genreMap[genreID]
	}

	// Проверяет, существует ли категория с найденным именем в карте keywordMap
	var genreWords []string
	for _, genreName := range genreNames {
		keywords, ok := keywordMap[genreName]
		if ok && len(keywords) > 0 {
			genreWords = append(genreWords, keywords...)
		} else {
			genreWords = append(genreWords, "")
		}
	}

	// Перебираем слова из списка wordsToFind и слова из найденной категории.
	// Внутренний цикл сравнивает каждое слово из wordsToFind с каждым словом из categoryWords
	for _, word := range wordsToFind {
		for _, catWord := range genreWords {
			if strings.ToLower(word) == catWord {
				foundWords = append(foundWords, word)
				break
			}
		}
	}

	return foundWords, nil
}

// Функция для получение списка слов для любой категории
// -----------------------------------------------------
func findPopularWordsMap(keywords types.KeywordList) (map[string][]string, error) {
	// Инициализация карты
	keywordMap := make(map[string][]string)

	// Итерация по категориям
	for _, category := range keywords.Keywords {
		// Проверка ID жанра
		categoryName, ok := genreMap[category.GenreID]
		if !ok {
			return nil, fmt.Errorf("unknown genre ID: %d", category.GenreID)
		}

		// Запись в карту
		keywordMap[categoryName] = category.Words
	}

	return keywordMap, nil
}
