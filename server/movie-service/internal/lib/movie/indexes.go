package movie

import (
	"bufio"
	"fmt"
	"go-movie-service/internal/common/constants"
	"go-movie-service/internal/common/utils"
	"go-movie-service/internal/lib"
	"go-movie-service/internal/types"
	"io"
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
		logger.Error(fmt.Sprintf("page %s not found in index", page))
		return 3
	}

	return uint64(offset)
}

func loadIndex(path string) (map[int]uint16, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	index := make(map[int]int64)
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

		index[page] = offset
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return index, nil
}

func getPageByOffset(filePath string, offset int64) (types.Movies, error) {
	var page types.Movies

	file, err := os.Open(filePath)
	if err != nil {
		return page, err
	}
	defer file.Close()

	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return page, err
	}

	decoder := utils.JSON.NewDecoder(file)
	if err := decoder.Decode(&page); err != nil {
		return page, err
	}

	return page, nil
}
