package pkg

import (
	"flicksfi/internal/types"
	"math/rand"
	"time"
)

// Удаление дубликатов из массива
// ------------------------------
func DeleteDublicateMovies(arr []types.Movie) []types.Movie {
	seen := make(map[int]bool)
	var result []types.Movie

	for _, movie := range arr {
		if _, ok := seen[movie.Id]; !ok {
			seen[movie.Id] = true
			result = append(result, movie)
		}
	}

	return result
}

// Массив с random начальным индексом
// ----------------------------------
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

	fillterArray := DeleteDublicateMovies(arr[startIndex-1 : endIndex])
	return fillterArray
}
