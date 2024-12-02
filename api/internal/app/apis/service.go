package apis

import (
	"context"
	"database/sql"
	"flick_finder/internal/types"
	"fmt"
	"strconv"
	"strings"
	"time"
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

// ---------------------------------
// Получение лимитов из БД по UserID
// ---------------------------------
func (s *Service) GetLimitText(uuid string) (types.Limiter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, "select * from limiter where uuid = ?", uuid)
	if err != nil {
		return types.Limiter{}, fmt.Errorf("error execute query to db")
	}

	var limiter types.Limiter
	for rows.Next() {
		err := rows.Scan(&limiter.ID, &limiter.UUID, &limiter.TextLimiter, &limiter.YoutubeLimit, &limiter.UpdateAt)
		if err != nil {
			return types.Limiter{}, fmt.Errorf("error scaning data")
		}
	}

	return limiter, nil
}

// ---------------------------------
// Уменьшение лимита текста из БД по UserID
// ---------------------------------
func (s *Service) ReducingLimitText(uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, "update limiter set text_limit = text_limit - 1 where uuid = ?", uuid)
	if err != nil {
		return fmt.Errorf("error reducing limit")
	}

	return nil
}
