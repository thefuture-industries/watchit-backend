package movie

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-movie-service/internal/common/constants"
	"go-movie-service/internal/lib"
	"go-movie-service/internal/types"
	"os"
	"strconv"
	"strings"
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
		var page types.Movies
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

func LoadIDX() map[uint32]uint32 {
	idx := make(map[uint32]uint32)

	file, _ := os.Open(constants.MOVIE_IDX_PATH)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) == 2 {
			page, _ := strconv.ParseUint(parts[0], 10, 32)
			offset, _ := strconv.ParseUint(parts[1], 10, 32)
			idx[uint32(page)] = uint32(offset)
		}
	}

	return idx
}
