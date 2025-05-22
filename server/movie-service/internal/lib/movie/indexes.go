package movie

import (
	"bufio"
	"go-movie-service/internal/common/utils"
	"go-movie-service/internal/types"
	"io"
	"os"
	"strconv"
	"strings"
)

func loadIndex(path string) (map[int]int64, error) {
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

		page, err1 := strconv.Atoi(parts[0])
		offset, err2 := strconv.ParseInt(parts[1], 10, 64)
		if err1 != nil || err2 != nil {
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
