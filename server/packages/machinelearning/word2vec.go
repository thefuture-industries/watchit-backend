package machinelearning

import (
	"strings"

	"github.com/james-bowman/nlp"
)

type MovieMatcher struct {
	vectorizer  *nlp.CountVectoriser
	transformer *nlp.TfidfTransformer
	threshold   float64
}

func NewMovieMatcher(threshold float64) *MovieMatcher {
	return &MovieMatcher{
		vectorizer:  nlp.NewCountVectoriser(),
		transformer: nlp.NewTfidfTransformer(),
		threshold:   threshold,
	}
}

func (mm *MovieMatcher) GetVectorizer() *nlp.CountVectoriser {
	return mm.vectorizer
}

func (mm *MovieMatcher) GetTransformer() *nlp.TfidfTransformer {
	return mm.transformer
}

func (mm *MovieMatcher) GetThreshold() float64 {
	return mm.threshold
}

// Word2VecPlot - функция для визуализации векторов слов
// -----------------------------------------------------
func Word2VecPlot(str1, str2, title string) float64 {
	var weight float64 = 0

	if str1 == "" || str2 == "" {
		return 0.0
	}

	stemmSTR_1 := Stemming(str1)
	stemmSTR_2 := Stemming(str2)

	extractText := func(str string, markers []struct{ start, end string }) string {
		for _, marker := range markers {
			if strings.Contains(str, marker.start) {
				start := strings.Index(str, marker.start) + len(marker.start)
				end := strings.LastIndex(str, marker.end)

				if start < end {
					return strings.TrimSpace(str[start:end])
				}
			}
		}

		return str
	}

	markers := []struct{ start, end string }{
		{"<?>", "<?>"},
		{"<>", "<>"},
		{"=>", "<="},
	}

	// Проверяем на совпадение названия
	for _, word := range stemmSTR_2 {
		// Фильмы с предположительным названием
		if strings.Contains(word, "<?>") {
			// Извлекаем текст между специальными символами и сравниваем с title
			clean_text := extractText(word, markers)
			if strings.EqualFold(title, clean_text) {
				weight += 10000.0
			}
		}

		// Фильмы с точными моментами
		if strings.Contains(word, "<>") {
			clean_text := extractText(word, markers)
			if strings.EqualFold(title, clean_text) || strings.EqualFold(str1, clean_text) {
				weight++
			}
		}

		// Фильмы по актерам
		if strings.Contains(word, "=>") {
			clean_text := extractText(word, markers)
			if strings.EqualFold(title, clean_text) || strings.EqualFold(str1, clean_text) {
				weight++
			}
		}
	}

	tone1 := AnalyzeTone(str1)
	tone2 := AnalyzeTone(str2)
	similarity := CalculateToneSimilarity(tone1, tone2)

	// Проверяем тональность
	if similarity >= 0.73 {
		weight += 1.0
	}

	// Создаем мапу для быстрой проверки совпадения слов
	stemm1Map := make(map[string]bool)
	for _, s := range stemmSTR_1 {
		stemm1Map[s] = true
	}

	// Ищем вектор отображающий смысл
	cos_vect := CosineVector(stemmSTR_2)

	// Проверяем совпадение слов
	for _, stemm2 := range stemmSTR_2 {
		if stemm1Map[stemm2] {
			weight++
		}
	}

	// Ищем слово отображающее смысл
	if stemm1Map[cos_vect] {
		weight++
	}

	if float64(len(stemmSTR_1)) < weight {
		return 1.0
	}

	tfidf := float64(float64(weight) / float64(len(stemmSTR_1)))
	return tfidf
}
