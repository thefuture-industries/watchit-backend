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

	file, err := os.Open(constants.MOVIE_PIDX_PATH_WRITE)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	defer file.Close()
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
