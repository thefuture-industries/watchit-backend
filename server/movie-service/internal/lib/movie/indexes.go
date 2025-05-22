package movie

import (
	"bufio"
	"fmt"
	"go-movie-service/internal/common/constants"
	"go-movie-service/internal/lib"
	"os"
	"strconv"
	"strings"
)

func PIDX(page uint16) uint64 {
	logger := lib.NewLogger()

	index, err := loadIndex(constants.MOVIE_PIDX_PATH_READ)
	if err != nil {
		logger.Error(fmt.Sprintf("error loading indexes: %s", err.Error()))
		return 3
	}

	offset, ok := index[page]
	if !ok {
		logger.Error(fmt.Sprintf("page %d not found in index", page))
		return 3
	}

	return uint64(offset)
}

func loadIndex(path string) (map[uint16]int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	index := make(map[uint16]int64)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		page, pageErr := strconv.Atoi(parts[0])
		offset, offsetErr := strconv.ParseInt(parts[1], 10, 64)
		if pageErr != nil || offsetErr != nil {
			continue
		}

		index[uint16(page)] = offset
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return index, nil
}
