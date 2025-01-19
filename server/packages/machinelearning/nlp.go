package machinelearning

import (
	"regexp"
	"strings"
)

var (
	// Предварительно скомпилированные регулярные выражения
	splitRegex   = regexp.MustCompile(`[\s.,!?;:]+`)
	stop_numbers = regexp.MustCompile(`\d`)
)

var stop_words = map[string]bool{
	// Местоимения
	"i":      true,
	"me":     true,
	"my":     true,
	"myself": true,
	"we":     true,
	"our":    true,
	"ours":   true,
	"us":     true,
	"this":   true,
	"these":  true,
	"those":  true,
	"he":     true,
	"she":    true,
	"it":     true,
	"its":    true,
	"itself": true,
	"them":   true,
	"their":  true,
	// Предлоги
	"by":     true,
	"from":   true,
	"into":   true,
	"during": true,
	"after":  true,
	"before": true,
	"above":  true,
	"below":  true,
	// Союзы
	"but":  true,
	"or":   true,
	"as":   true,
	"if":   true,
	"then": true,
	"else": true,
	"so":   true,
	// Вспомогательные глаголы
	"will":   true,
	"would":  true,
	"shall":  true,
	"should": true,
	"can":    true,
	"could":  true,
	"may":    true,
	"might":  true,
	"must":   true,
	// Наречия
	"here":  true,
	"there": true,
	"now":   true,
	"just":  true,
	"very":  true,
	"too":   true,
	// Другие частые слова
	"about": true,
	"again": true,
	"once":  true,
	"been":  true,
	"being": true,
	"both":  true,
	"each":  true,
	"few":   true,
	"more":  true,
	"most":  true,
	"other": true,
	"some":  true,
	"such":  true,
	"what":  true,
	"which": true,
	"why":   true,
	"how":   true,
}

// NLP обработка естественного языка
// ---------------------------------
func Stemming(str string) []string {
	// Если пустая строка
	if len(str) == 0 {
		return nil
	}

	// Функция для извлечения фраз в специальных скобках
	var extractPhrases = func(text string) ([]string, string) {
		var phrases []string

		// Поиск фраз в специальных символах
		patterns := []struct {
			start, end string
		}{
			{"<?>", "<?>"},
			{"<>", "<>"},
			{"=>", "<="},
		}

		modifiedText := text
		for _, pattern := range patterns {
			startIdx := strings.Index(modifiedText, pattern.start)
			for startIdx != -1 {
				// Находим конец фразы
				endIdx := strings.Index(modifiedText[startIdx+len(pattern.start):], pattern.end)
				if endIdx == -1 {
					break
				}

				// Корректируем endIdx, чтобы он указывал на правильную позицию
				endIdx = startIdx + len(pattern.start) + endIdx

				// Извлекаем полную фразу вместе с содержимым
				fullPhrase := modifiedText[startIdx : endIdx+len(pattern.end)]
				phrases = append(phrases, fullPhrase)

				// Заменяем найденную фразу пробелами
				placeholder := strings.Repeat(" ", len(fullPhrase))
				modifiedText = modifiedText[:startIdx] + placeholder + modifiedText[endIdx+len(pattern.end):]

				startIdx = strings.Index(modifiedText, pattern.start)
			}
		}

		return phrases, modifiedText
	}

	// Извлекаем фразы и получаем модифицированный текст
	phrases, modifiedText := extractPhrases(str)

	// Вложенная функция для обработки окончаний
	var suffix = func(word string) string {
		switch {
		case strings.HasSuffix(word, "fulness"):
			return word[:len(word)-7]
		case strings.HasSuffix(word, "ousness"):
			return word[:len(word)-7]
		case strings.HasSuffix(word, "ization"):
			return word[:len(word)-7] + "ize"
		case strings.HasSuffix(word, "ational"):
			return word[:len(word)-7] + "ate"
		case strings.HasSuffix(word, "tional"):
			return word[:len(word)-6] + "tion"
		case strings.HasSuffix(word, "alize"):
			return word[:len(word)-5] + "al"
		case strings.HasSuffix(word, "icate"):
			return word[:len(word)-5] + "ic"
		case strings.HasSuffix(word, "ative"):
			return word[:len(word)-5]
		case strings.HasSuffix(word, "ement"):
			return word[:len(word)-5]
		case strings.HasSuffix(word, "ingly"):
			return word[:len(word)-5]
		case strings.HasSuffix(word, "fully"):
			return word[:len(word)-5]
		case strings.HasSuffix(word, "ably"):
			return word[:len(word)-4]
		case strings.HasSuffix(word, "ibly"):
			return word[:len(word)-4]
		case strings.HasSuffix(word, "ing"):
			stem := word[:len(word)-3]
			if len(stem) > 0 && func(c byte) bool {
				return !strings.ContainsRune("aeiou", rune(c))
			}(stem[len(stem)-1]) {
				return stem + "e"
			}
			return stem
		case strings.HasSuffix(word, "ies"):
			return word[:len(word)-3] + "y"
		case strings.HasSuffix(word, "ive"):
			return word[:len(word)-3]
		case strings.HasSuffix(word, "es"):
			return word[:len(word)-2]
		case strings.HasSuffix(word, "ly"):
			return word[:len(word)-2]
		case strings.HasSuffix(word, "ed"):
			stem := word[:len(word)-2]
			if len(stem) > 0 && func(c byte) bool {
				return !strings.ContainsRune("aeiou", rune(c))
			}(stem[len(stem)-1]) {
				return stem + "e"
			}
			return stem
		case strings.HasSuffix(word, "'s"):
			return word[:len(word)-2]
		case strings.HasSuffix(word, "s"):
			return word[:len(word)-1]
		default:
			return word
		}
	}

	// Вложенная функция для лемматизации
	var lemmatize = func(word string) string {
		switch word {
		// Глаголы
		case "am", "is", "are", "was", "were":
			return "be"
		case "has", "have", "had":
			return "have"
		case "does", "did":
			return "do"
		case "going", "goes", "went":
			return "go"
		case "made", "making":
			return "make"
		case "saw", "seen", "seeing":
			return "see"
		case "came", "coming":
			return "come"
		case "took", "takes", "taken", "taking":
			return "take"

		// Прилагательные
		case "better", "best":
			return "good"
		case "worse", "worst":
			return "bad"
		case "bigger", "biggest":
			return "big"
		case "smaller", "smallest":
			return "small"
		case "larger", "largest":
			return "large"

		// Наречия и другие части речи
		case "more", "most":
			return "many"
		case "less", "least":
			return "little"
		case "further", "furthest":
			return "far"

		// Существительные
		case "children":
			return "child"
		case "people":
			return "person"
		case "lives":
			return "life"
		case "wives":
			return "wife"

		// Прилагательные в сравнительной и превосходной степени
		case "fewer", "fewest":
			return "few"

		default:
			return word
		}
	}

	// Разделение предложение на массив слов
	text := splitRegex.ReplaceAllString(modifiedText, " ")
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text)

	var str_result = make([]string, 0, len(words)+len(phrases))

	// Работы с предложением
	for _, word := range words {
		// Пропускаем короткие слова и стоп-слова
		if len(word) <= 2 || stop_numbers.MatchString(word) {
			continue
		}

		// Обработка слов составных (-)
		if strings.Contains(word, "-") {
			parts := strings.Split(word, "-")
			for _, part := range parts {
				if !stop_words[part] && len(part) > 2 {
					processed := lemmatize(suffix(part))
					str_result = append(str_result, processed)
				}
			}

			continue
		}

		// Обработка слов
		if !stop_words[word] && !stop_numbers.MatchString(word) {
			processed := lemmatize(suffix(word))
			if processed != "" && len(processed) > 2 {
				str_result = append(str_result, processed)
			}
		}
	}

	str_result = append(str_result, phrases...)

	return str_result
}
