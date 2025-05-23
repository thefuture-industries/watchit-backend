package machinelearning

import "sort"

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

	wcList := make([]wc, 0, len(wCount))
	for w, c := range wCount {
		wcList = append(wcList, wc{w, c})
	}

	sort.Slice(wcList, func(i, j int) bool {
		return wcList[i].count > wcList[j].count
	})

	limit := int(this.avgOverview) * len(documents)
	if limit > len(wcList) {
		limit = len(wcList)
	}

	this.vocabulary = make([]string, limit)
	this.vocabularyIndexMap = make(map[string]int, limit)

	for i := 0; i < limit; i++ {
		this.vocabulary[i] = wcList[i].word
		this.vocabularyIndexMap[wcList[i].word] = i
	}
}
