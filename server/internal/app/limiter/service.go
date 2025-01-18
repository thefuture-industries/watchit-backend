package limiter

import (
	"context"
	"database/sql"
	"flicksfi/cmd/configuration"
	"flicksfi/internal/types"
	"fmt"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	db      *sql.DB
	logger  *zap.Logger
	monitor *configuration.Track
}

func NewService(db *sql.DB, logger *zap.Logger, monitor *configuration.Track) *Service {
	return &Service{
		db:      db,
		logger:  logger,
		monitor: monitor,
	}
}

// Получение лимитов из БД по UserID
// ---------------------------------
func (s *Service) GetLimits(uuid string) (*types.Limiter, error) {
	// мониторинг времени запроса
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// запрос к БД
	queryStart := time.Now()
	row := s.db.QueryRowContext(ctx, "select * from limiter where uuid = ? limit 1", uuid)

	// читаем из результата
	limiter := new(types.Limiter)
	err := row.Scan(&limiter.ID, &limiter.UUID, &limiter.TextLimiter, &limiter.YoutubeLimit, &limiter.UpdateAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("limiter not found")
		}

		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", uuid),
			zap.Error(err))
		return nil, fmt.Errorf("failed to get limiter: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	return limiter, nil
}

// Уменьшение лимита текста из БД по UserID
// ----------------------------------------
func (s *Service) ReducingLimitText(uuid string) error {
	// мониторинг времени запроса
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// запрос к БД
	queryStart := time.Now()
	_, err := s.db.ExecContext(ctx, "update limiter set text_limit = text_limit - 1 set update_at = ? where uuid = ?", time.Now().Format("2006-01-02 15:04:05"), uuid)
	if err != nil {
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", uuid),
			zap.Error(err))
		return fmt.Errorf("reduce limit text: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	return nil
}

// Обновление лимитов после 24 часов
// ---------------------------------
func (s *Service) UpdateLimits(uuid string) error {
	// мониторинг времени запроса
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Получение лимитов
	var status string
	queryStart := time.Now()
	err := s.db.QueryRowContext(ctx, "select case when timestampdiff(hour, update_at, now()) >= 24 then '1' else '0' end as status from limiter where uuid = ?", uuid).Scan(&status)
	if err != nil {
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", uuid),
			zap.Error(err))
		return fmt.Errorf("get limiter: %w", err)
	}

	status_number, err := strconv.Atoi(status)
	if err != nil {
		// Логирование ошибки
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", uuid),
			zap.Error(err))
		return fmt.Errorf("convert to int: %w", err)
	}

	if status_number == 0 {
		return nil
	}

	_, err = s.db.ExecContext(ctx, "update limiter set text_limit = 2, youtube_limit = 3, update_at = ? where uuid = ?", time.Now().Format("2006-01-02 15:04:05"), uuid)
	if err != nil {
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", uuid),
			zap.Error(err))
		return fmt.Errorf("update limiter: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	return nil
}
