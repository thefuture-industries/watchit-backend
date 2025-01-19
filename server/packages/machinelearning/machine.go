package machinelearning

import (
	"sort"
	"strings"
)

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
func TF_IDF(str1, str2 string, add_size float64) float64 {
	var genres []string = []string{"Action", "Adventure", "Animation", "Comedy", "Crime", "Documentary", "Drama", "Family", "Fantasy", "History", "Horror", "Music", "Mystery", "Romance", "Science Fiction", "Thriller", "War", "Western"}
	var genres_map = make(map[string]bool)
	for _, genre := range genres {
		genres_map[strings.ToLower(genre)] = true
	}

	var weight float64 = 0

	// Если строки пустые
	if str1 == "" || str2 == "" {
		return 0.0
	}

	// Стеминг слов
	stemmSTR_1 := Stemming(str1)
	stemmSTR_2 := Stemming(str2)

	// Из массива слов в map
	stemm1Map := make(map[string]bool)
	for _, s := range stemmSTR_1 {
		stemm1Map[s] = true
	}

	// Получение смысла предложения
	cos_vect := CosineVector(stemmSTR_2)

	// Анализ эмоционального тона
	tone1 := AnalyzeTone(str1)
	tone2 := AnalyzeTone(str2)
	similarity := CalculateToneSimilarity(tone1, tone2)

	if similarity >= 0.63 {
		weight += 1.0
	}

	for _, stemm2 := range stemmSTR_2 {
		// Проверяем совпадение слов
		if stemm1Map[stemm2] {
			weight++
		}

		// Проверяем жанры
		if genres_map[stemm2] && stemm1Map[stemm2] {
			weight++
		}
	}

	// Ищем слово отображающее смысл
	if stemm1Map[cos_vect] {
		weight++
	}

	// Если вышли за пределы, то фильм абсолютно похож
	if float64(len(stemmSTR_1)) < weight {
		return 1.0
	}

	tfidf := float64(float64(weight)/float64(len(stemmSTR_1)) + add_size)
	return tfidf
}
