package machinelearning

import (
	"math/rand"
	"time"
)

func ShuffleArray[T any](slice []T) error {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	return nil
}
