package machinelearning

import (
	"go-movie-service/internal/lib"
	"math"
)

type TFIDFBuilder struct {
	logger *lib.Logger
}

func NewTFIDFBuilder() *TFIDFBuilder {
	return &TFIDFBuilder{
		logger: lib.NewLogger(),
	}
}

func (this *TFIDFBuilder) TF(df int, N int) float32 {
	return float32(df) / float32(N)
}

func (this *TFIDFBuilder) IDF(N int, df int) float64 {
	return math.Log(float64(N) / (1 + float64(df)))
}

func (this *TFIDFBuilder) TFIDF(tf float32, idf float64) float64 {
	return float64(tf) * idf
}
