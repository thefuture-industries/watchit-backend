package apis

import (
	"database/sql"
	"strconv"
	"strings"
	"unicode"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

// // Парсинг жанров массива в [int]
// // ------------------------------
func (s *Service) ArrayGenreIDS(str string) ([]int, error) {
	parts := strings.FieldsFunc(str, func(r rune) bool {
		return unicode.IsSpace(r) || r == ','
	})

	result := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}

		result[i] = num
	}

	return result, nil
}
