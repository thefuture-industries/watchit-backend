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

	for _, word := range stemmSTR_2 {
		if strings.Contains(word, "<?>") {
			// Извлекаем текст между специальными символами и сравниваем с title
			clean_text := extractText(word, markers)
			if strings.ToLower(title) == strings.ToLower(clean_text) {
				weight += 10000.0
			}
		}
	}

	tone1 := AnalyzeTone(str1)
	tone2 := AnalyzeTone(str2)
	similarity := CalculateToneSimilarity(tone1, tone2)

	if similarity >= 0.73 {
		weight += 1.0
	}

	if float64(len(stemmSTR_1)) < weight {
		return 1.0
	}

	tfidf := float64(float64(weight) / float64(len(stemmSTR_1)))
	return tfidf
}
