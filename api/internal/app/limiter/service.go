package limiter

import (
	"context"
	"database/sql"
	"flicksfi/internal/types"
	"fmt"
	"strconv"
	"time"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

// Получение лимитов из БД по UserID
// ---------------------------------
func (s *Service) GetLimits(uuid string) (types.Limiter, error) {
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

// Уменьшение лимита текста из БД по UserID
// ----------------------------------------
func (s *Service) ReducingLimitText(uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, "update limiter set text_limit = text_limit - 1 set update_at = ? where uuid = ?", time.Now().Format("2006-01-02 15:04:05"), uuid)
	if err != nil {
		return fmt.Errorf("error reducing limit")
	}

	return nil
}

// Обновление лимитов после 24 часов
// ---------------------------------
func (s *Service) UpdateLimits(uuid string) error {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Получение лимитов
	var status string
	err := s.db.QueryRowContext(ctx, "select case when timestampdiff(hour, update_at, now()) >= 24 then '1' else '0' end as status from limiter where uuid = ?", uuid).Scan(&status)
	if err != nil {
		return fmt.Errorf("error database request")
	}

	status_number, err := strconv.Atoi(status)
	if err != nil {
		return fmt.Errorf("error convert to int")
	}

	if status_number == 0 {
		return nil
	}

	_, err = s.db.ExecContext(ctx, "update limiter set text_limit = 2, youtube_limit = 3, update_at = ? where uuid = ?", time.Now().Format("2006-01-02 15:04:05"), uuid)
	if err != nil {
		return fmt.Errorf("error database request")
	}

	return nil
}
