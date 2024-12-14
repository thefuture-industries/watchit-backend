package pkg

import (
	"regexp"
	"sort"
	"strings"
)

// NLP обработка естественного языка
// ---------------------------------
func Stemming(str string) []string {
	var stop_words map[string]bool = map[string]bool{
		"a":     true,
		"on":    true,
		"is":    true,
		"at":    true,
		"the":   true,
		"be":    true,
		"are":   true,
		"in":    true,
		"of":    true,
		"to":    true,
		"and":   true,
		"an":    true,
		"he":    true,
		"she":   true,
		"you":   true,
		"they":  true,
		"hi":    true,
		"that":  true,
		"than":  true,
		"thats": true,
		"his":   true,
		"him":   true,
		"with":  true,
		"for":   true,
		"where": true,
		"who":   true,
		"when":  true,
		"no":    true,
		"yes":   true,
		"all":   true,
		"her":   true,
		"has":   true,
		"out":   true,
		"'":     true,
	}
	stop_numbers := regexp.MustCompile(`\d`)
	var str_result []string

	// Если пустая строка
	if len(str) == 0 {
		return nil
	}

	// Разделение предложение на массив слов
	re := regexp.MustCompile(`[\s.,!?;:]+`)
	text := re.ReplaceAllString(str, " ")
	text = strings.TrimSpace(text)
	words := strings.Fields(text)

	// Работы с текстом
	for _, word := range words {
		lower := strings.ToLower(word)
		if _, ok := stop_words[lower]; !ok && !stop_numbers.MatchString(lower) {
			if strings.HasSuffix(lower, "ing") {
				lower = lower[:len(lower)-3]
			} else if strings.HasSuffix(lower, "ed") {
				lower = lower[:len(lower)-2]
			} else if strings.HasSuffix(lower, "s") {
				lower = lower[:len(lower)-1]
			}

			str_result = append(str_result, lower)
		}
	}

	return str_result
}

// Получение косинусойдного вектора предложения
// --------------------------------------------
func CosineVector(str []string) string {
	// Структура
	type WordCount struct {
		word  string
		count int
	}

	// Переменные
	word_counts := make(map[string]int)
	var result []WordCount

	// Счет повторяющих слов
	for _, word := range str {
		word_counts[word]++
	}

	// Конвертация из map в массив
	for key, value := range word_counts {
		result = append(result, WordCount{key, value})
	}

	// Сортировка массива
	sort.Slice(result, func(i, j int) bool {
		return result[i].count > result[j].count
	})

	return result[0].word
}

// Сравнение и получение от 0-1 схожесть двух строк
// ------------------------------------------------
// str1 - по какому предложению искать : str2 - с помощью какого
func TF_IDF_MOVIE(str1, str2 string, add_size float64) float64 {
	var genres []string = []string{"Action", "Adventure", "Animation", "Comedy", "Crime", "Documentary", "Drama", "Family", "Fantasy", "History", "Horror", "Music", "Mystery", "Romance", "Science Fiction", "Thriller", "War", "Western"}
	var findsWords int

	// Если строки пустые
	if str1 == "" || str2 == "" {
		return 0.0
	}

	// Стеминг слов
	stemmSTR_1 := Stemming(str1)
	stemmSTR_2 := Stemming(str2)

	// Получение смысла предложения
	cos_vect := CosineVector(stemmSTR_2)

	for _, stemm1 := range stemmSTR_1 {
		for _, stemm2 := range stemmSTR_2 {
			if stemm1 == stemm2 {
				findsWords++
			}
		}

		// Ищем слово отображающее смысл
		if stemm1 == cos_vect {
			findsWords++
		}
	}

	// Ищем жанр в тексте
	for _, stemm2 := range stemmSTR_2 {
		for _, genre := range genres {
			if strings.ToLower(genre) == stemm2 {
				for _, stemm1 := range stemmSTR_1 {
					if stemm1 == strings.ToLower(genre) {
						findsWords++
						break
					}
				}
			}
		}
	}

	// Если вышли за пределы, то фильм абсолютно похож
	if len(stemmSTR_1) < findsWords {
		return 1.0
	}

	tfidf := float64(float64(findsWords)/float64(len(stemmSTR_1)) + add_size)
	return tfidf
}
