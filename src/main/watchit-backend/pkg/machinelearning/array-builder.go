package machinelearning

import (
	"math/rand"
	"time"
)

func ShuffleArray[T any](slice []T) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	return nil
}
