package movie

import (
	"go-movie-service/internal/types"
	"io"
	"os"
)

func getPageByOffset(filePath string, offset int64) (types.Movies, error) {
	var page types.Movies

	f, err := os.Open(filePath)
	if err != nil {
		return page, err
	}
	defer f.Close()

	_, err = f.Seek(offset, io.SeekStart)
	if err != nil {
		return page, err
	}

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&page); err != nil {
		return page, err
	}

	return page, nil
}
