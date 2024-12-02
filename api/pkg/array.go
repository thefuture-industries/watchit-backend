package pkg

import (
	"flick_finder/internal/types"
	"math/rand"
	"time"
)

func TruncateArrayMovies(arr []types.Movie) []types.Movie {
	var arrayLength int = len(arr)
	var startIndex int = 1

	if arrayLength >= 100 {
		rand.Seed(time.Now().UnixNano())
		startIndex = rand.Intn(arrayLength-99) + 1
	}

	var endIndex int = startIndex + 99
	if endIndex > arrayLength {
		endIndex = arrayLength
	}

	return arr[startIndex-1 : endIndex]
}
