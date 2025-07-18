package machinelearning

import (
	"regexp"
	"strings"
	"unicode"
)

type NLPBuilder struct {
	stopWords map[string]struct{}
}

func NewNLPBuilder() *NLPBuilder {
	words := []string{
		"i", "me", "my", "myself", "we", "our", "ours", "us", "this", "these", "those",
		"he", "she", "it", "its", "itself", "them", "their", "by", "from", "into", "during",
		"after", "before", "above", "below", "but", "or", "as", "if", "then", "else", "so",
		"will", "would", "shall", "should", "can", "could", "may", "might", "must",
		"here", "there", "now", "just", "very", "too", "about", "again", "once", "been",
		"being", "both", "each", "few", "more", "most", "other", "some", "such", "what",
		"which", "why", "how", "the",
	}

	stopWords := make(map[string]struct{})
	for _, w := range words {
		stopWords[w] = struct{}{}
	}

	return &NLPBuilder{
		stopWords: stopWords,
	}
}

func (nlp *NLPBuilder) Preprocess(input string) []string {
	if input == "" {
		return []string{}
	}

	re := regexp.MustCompile(`[()<>""{}\[\]]`)
	input = re.ReplaceAllString(input, " ")

	words := strings.FieldsFunc(strings.ToLower(input), func(r rune) bool {
		return unicode.IsSpace(r) || strings.ContainsRune(",.!?;:\t\n\r", r)
	})

	var output []string
	for _, word := range words {
		for _, part := range strings.Split(word, "-") {
			if len(part) <= 2 || nlp.isStopWord(part) || hasDigit(part) {
				continue
			}
			stemmed := nlp.stemming(part)
			lemmatized := nlp.lemmatize(stemmed)
			if len(lemmatized) > 2 {
				output = append(output, lemmatized)
			}
		}
	}

	return output
}

// nolint
func (nlp *NLPBuilder) lemmatize(word string) string {
	switch word {
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
	case "more", "most":
		return "many"
	case "less", "least":
		return "little"
	case "further", "furthest":
		return "far"
	case "children":
		return "child"
	case "people":
		return "person"
	case "lives":
		return "life"
	case "wives":
		return "wife"
	case "fewer", "fewest":
		return "few"
	default:
		return word
	}
}

// nolint
func (nlp *NLPBuilder) stemming(word string) string {
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
		if len(stem) > 0 && !isVowel(rune(stem[len(stem)-1])) {
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
		if len(stem) > 0 && !isVowel(rune(stem[len(stem)-1])) {
			return stem + "e"
		}
		return stem
	default:
		return word
	}
}

func (nlp *NLPBuilder) isStopWord(word string) bool {
	_, exists := nlp.stopWords[word]
	return exists
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiou", r)
}

func hasDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
