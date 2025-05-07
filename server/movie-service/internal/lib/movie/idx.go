package movie

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-movie-service/internal/common/constants"
	"go-movie-service/internal/lib"
	"go-movie-service/internal/types"
	"os"
)

func MovieIDX() {
	logger := lib.NewLogger()

	file, err := os.Open(constants.MOVIE_JSON_PATH)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(bufio.NewReader(file))

	token, err := decoder.Token()
	if err != nil || token != json.Delim('[') {
		logger.Error("Not an array")
		return
	}

	index := make(map[int]uint64)
	for decoder.More() {
		offset, _ := file.Seek(0, os.SEEK_CUR)
		var page types.MoviePage
		if err := decoder.Decode(&page); err != nil {
			logger.Error(err.Error())
			return
		}

		index[int(page.Page)] = uint64(offset)
	}

	idxFile, _ := os.Create(constants.MOVIE_IDX_PATH)
	defer idxFile.Close()
	for page, offset := range index {
		fmt.Fprintf(idxFile, "%d %d\n", page, offset)
	}
}
