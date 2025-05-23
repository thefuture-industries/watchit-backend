package machinelearning

import (
	"go-movie-service/internal/types"
	"math"
	"sort"
	"sync"

	"gonum.org/v1/gonum/mat"
)

type LSABuilder struct {
	nlpBuilder         *NLPBuilder
	tfidfBuilder       *TFIDFBuilder
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
		tfidfBuilder:       NewTFIDFBuilder(),
		avgOverview:        39,
		top:                25,
		vocabularyIndexMap: nil,
		vocabulary:         nil,
		tokenizedDocs:      nil,
		idfCache:           nil,
	}
}

func (this *LSABuilder) addVocabulary(documents []string) {
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

func (this *LSABuilder) calcIDF() {
	N := len(this.tokenizedDocs)
	this.idfCache = make(map[string]float64, len(this.vocabulary))
	dfCount := make(map[string]int)

	for _, doc := range this.tokenizedDocs {
		seen := make(map[string]bool)
		for _, token := range doc {
			if !seen[token] {
				dfCount[token]++
				seen[token] = true
			}
		}
	}

	for word := range this.vocabularyIndexMap {
		df := dfCount[word]

		if df == 0 {
			df = 1
		}

		this.idfCache[word] = this.tfidfBuilder.IDF(N, df)
	}
}

func (this *LSABuilder) AnalyzeByMovie(documents []types.Movie, inputText string) (*mat.Dense, []types.Movie) {
	documentsTake := documents
	if len(documents) > 0 {
		documentsTake = documents[:len(documents)/5]
	}

	dOverview := make([]string, len(documentsTake)+1)
	for i, movie := range documentsTake {
		dOverview[i] = movie.Overview
	}
	dOverview[len(dOverview)-1] = inputText

	this.addVocabulary(dOverview)
	this.calcIDF()

	nDocs := len(dOverview)
	nTerms := len(this.vocabulary)

	data := make([]float64, nDocs*nTerms)

	var wg sync.WaitGroup
	sem := make(chan struct{}, 100)

	for i := 0; i < nDocs; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()

			tokens := this.tokenizedDocs[i]

			termFreq := make(map[string]int)
			for _, token := range tokens {
				termFreq[token]++
			}

			total := len(tokens)
			for token, count := range termFreq {
				idf, okIDF := this.idfCache[token]
				idx, okIDX := this.vocabularyIndexMap[token]

				if okIDF && okIDX && idx < nTerms && i < nDocs {
					tf := this.tfidfBuilder.TF(count, total)
					tfidf := this.tfidfBuilderTFIDF(tf, idf)

					data[i*nTerms+idx] = tfidf
				}
			}
		}(i)
	}
	wg.Wait()

	matrix := mat.NewDense(nDocs, nTerms, data)

	var svd mat.SVD
	ok := svd.Factorize(matrix, mat.SVDThin)
	if !ok {
		return nil, nil
	}

	U := mat.NewDense(nDocs, nDocs, nil)
	svd.UTo(U)

	return U, documentsTake
}

func (this *LSABuilder) cosineSimilarity(a, b []float64) float64 {
	var dot, normA, normB float64
	for i := range a {
		dot += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	if normA == 0 || normB == 0 {
		return 0
	}

	return dot / (math.Sqrt(normA) * math.Sqrt(normB))
}
