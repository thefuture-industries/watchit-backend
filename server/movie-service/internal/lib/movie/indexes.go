package movie

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-movie-service/internal/common/constants"
	"go-movie-service/internal/lib"
	"go-movie-service/internal/types"
	"io"
	"os"
	"strconv"
	"strings"
)

func Pidx(page uint16) {
	logger := lib.NewLogger()

	file, err := os.Open(constants.MOVIE_JSON_PATH_WRITE)
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
		offset, _ := file.Seek(0, io.SeekCurrent)
		var page types.Movies
		if err := decoder.Decode(&page); err != nil {
			logger.Error(err.Error())
			return
		}

		index[int(page.Page)] = uint64(offset)
	}

	idxFile, _ := os.Create(constants.MOVIE_PIDX_PATH_WRITE)
	defer idxFile.Close()
	for page, offset := range index {
		fmt.Fprintf(idxFile, "%d %d\n", page, offset)
	}
}

func LoadPIDX() map[uint32]uint32 {
	logger := lib.NewLogger()
	idx := make(map[uint32]uint32)

	file, err := os.Open(constants.MOVIE_PIDX_PATH_WRITE)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
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
