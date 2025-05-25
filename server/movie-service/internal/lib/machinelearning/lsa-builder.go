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

func (lsa *LSABuilder) addVocabulary(documents []string) {
	if lsa.vocabulary != nil {
		return
	}

	lsa.tokenizedDocs = make([][]string, len(documents))
	wCount := mke(map[string]int)a

	for i, doc := range documents {
		tokens := lsa.nlpBuilder.Preprocess(doc)
		lsa.tokenizedDocs[i] = tokens

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

	limit := int(lsa.avgOverview) * len(documents)
	if limit > len(wcList) {
		limit = len(wcList)
	}

	lsa.vocabulary = make([]string, limit)
	lsa.vocabularyIndexMap = make(map[string]int, limit)

	for i := 0; i < limit; i++ {
		lsa.vocabulary[i] = wcList[i].Word
		lsa.vocabularyIndexMap[wcList[i].Word] = i
	}
}

func (lsa *LSABuilder) calcIDF() {
	N := len(lsa.tokenizedDocs)
	lsa.idfCache = make(map[string]float64, len(lsa.vocabulary))

	dfCount := make(map[string]int)
	var mu sync.Mutex

	var wg sync.WaitGroup
	for _, doc := range lsa.tokenizedDocs {
		wg.Add(1)
		go func(doc []string) {
			defer wg.Done()

			seen := make(map[string]bool)
			for _, token := range doc {
				seen[token] = true
			}

			mu.Lock()
			for token := range seen {
				dfCount[token]++
			}
			mu.Unlock()
		}(doc)
	}
	wg.Wait()

	for word := range lsa.vocabularyIndexMap {
		df := dfCount[word]

		if df == 0 {
			df = 1
		}

		lsa.idfCache[word] = lsa.tfidfBuilder.IDF(N, df)
	}
}

func (lsa *LSABuilder) AnalyzeByMovie(documents []types.Movie, inputText string) (*mat.Dense, []types.Movie) {
	documentsTake := documents
	if len(documents) > 0 {
		documentsTake = documents[:len(documents)/5]
	}

	overviews := make([]string, len(documentsTake)+1)
	for i, movie := range documentsTake {
		overviews[i] = movie.Overview
	}
	overviews[len(overviews)-1] = inputText

	lsa.addVocabulary(overviews)
	lsa.calcIDF()

	nDocs := len(overviews)
	nTerms := len(lsa.vocabulary)

	data := make([]float64, nDocs*nTerms)

	var wg sync.WaitGroup
	// sem := make(chan struct{}, 100)

	for i := 0; i < nDocs; i++ {
		wg.Add(1)
		// sem <- struct{}{}

		go func(i int) {
			defer wg.Done()
			// defer func() { <-sem }()

			tokens := lsa.tokenizedDocs[i]

			termFreq := make(map[string]int)
			for _, token := range tokens {
				termFreq[token]++
			}

			total := len(tokens)
			for token, count := range termFreq {
				idf, okIDF := lsa.idfCache[token]
				idx, okIDX := lsa.vocabularyIndexMap[token]

				if okIDF && okIDX && idx < nTerms && i < nDocs {
					tf := lsa.tfidfBuilder.TF(count, total)
					tfidf := lsa.tfidfBuilder.TFIDF(tf, idf)

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

func (lsa *LSABuilder) CosineSimilarity(a, b []float64) float64 {
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
