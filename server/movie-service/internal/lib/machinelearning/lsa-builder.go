package machinelearning

import (
	"go-movie-service/internal/types"
	"sort"
)

type LSABuilder struct {
	nlpBuilder         *NLPBuilder
	avgOverview        uint8
	top                int
	vocabularyIndexMap map[string]int
	vocabulary         []string
	tokenizedDocs      [][]string
	idfCache           map[string]float64
}

func NewLSABuilder() *LSABuilder {
	return &LSABuilder{
		nlpBuilder:         NewNLPBuilder(),
		avgOverview:        39,
		top:                25,
		vocabularyIndexMap: nil,
		vocabulary:         nil,
		tokenizedDocs:      nil,
		idfCache:           nil,
	}
}

func (this *LSABuilder) AddVocabulary(documents []string) {
	if this.vocabulary != nil {
		return
	}

	this.tokenizedDocs = make([][]string, len(documents))
	wCount := make(map[string]int)

	for i, doc := range documents {
		tokens := this.nlpBuilder.Preprocess(doc)
		this.tokenizedDocs[i] = tokens

		for _, token := range tokens {
			wCount[token]++
		}
	}

	wcList := make([]types.WC, 0, len(wCount))
	for w, c := range wCount {
		wcList = append(wcList, types.WC{w, c})
	}

	sort.Slice(wcList, func(i, j int) bool {
		return wcList[i].Count > wcList[j].Count
	})

	limit := int(this.avgOverview) * len(documents)
	if limit > len(wcList) {
		limit = len(wcList)
	}

	this.vocabulary = make([]string, limit)
	this.vocabularyIndexMap = make(map[string]int, limit)

	for i := 0; i < limit; i++ {
		this.vocabulary[i] = wcList[i].Word
		this.vocabularyIndexMap[wcList[i].Word] = i
	}
}

func CalcIDF() {
	N := len(tokenizedDocs)
	idfCache = make(map[string]float64, len(vocabulary))
	dfCount := make(map[string]int)

	for _, doc := range tokenizedDocs {
		seen := make(map[string]bool)
		for _, token := range doc {
			if !seen[token] {
				dfCount[token]++
				seen[token] = true
			}
		}
	}

	for word := range vocabularyIndexMap {
		df := dfCount[word]

		if df == 0 {
			df = 1
		}

		idfCache[word] = IDF(N, df)
	}
}
