package machinelearning

import (
	"math"
	"watchit/httpx/infra/logger"
)

type TFIDFBuilder struct {
	logger *logger.Logger
}

func NewTFIDFBuilder() *TFIDFBuilder {
	return &TFIDFBuilder{
		logger: logger.NewLogger(),
	}
}

func (tfidf *TFIDFBuilder) TF(df int, N int) float32 {
	return float32(df) / float32(N)
}

func (tfidf *TFIDFBuilder) IDF(N int, df int) float64 {
	return math.Log(float64(N) / (1 + float64(df)))
}

func (tfidf *TFIDFBuilder) TFIDF(tf float32, idf float64) float64 {
	return float64(tf) * idf
}
