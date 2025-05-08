package configuration

import (
	"flicksfi/internal/types"
	"math"

	"context"
	"database/sql"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	db      *sql.DB
	logger  *zap.Logger
	metrics *Metrics
	done    chan struct{}
}

// NewService создание сервиса мониторинга
// ---------------------------------------
func NewService(db *sql.DB, logger *zap.Logger) *Service {
	return &Service{
		db:      db,
		logger:  logger,
		metrics: NewMetrics(),
	}
}

// StartDBMonitoring запуск мониторинга БД
// ---------------------------------------
func (s *Service) StartDBMonitoring(interval time.Duration) {
	// Создание канала для завершения мониторинга
	s.done = make(chan struct{})

	// Запуск горутины для мониторинга БД
	go func() {
		// Создание тикера для периодического выполнения мониторинга
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		// Бесконечный цикл для постоянного мониторинга
		for {
			// Ожидание сигнала о завершении работы
			select {
			// Ожидание сигнала о завершении работы
			case <-ticker.C:
				// Получение статистики БД
				stats := s.db.Stats()
				// Установка метрики для количества открытых соединений
				s.metrics.DBConnections.Set(float64(stats.OpenConnections))

				// Проверка БД на доступность
				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
				err := s.db.PingContext(ctx)
				cancel()

				if err != nil {
					s.logger.Error("database health check failed", zap.Error(err))
				}

				// Логирование статистики БД
				s.logger.Info("database stats",
					zap.Int("open_connections", stats.OpenConnections),
					zap.Int("in_use", stats.InUse),
					zap.Int("idle", stats.Idle),
					zap.Int("wait_count", int(stats.WaitCount)),
					zap.Int("wait_duration", int(stats.WaitDuration.Seconds())),
				)

			// Ожидание сигнала о завершении работы
			case <-s.done:
				return
			}
		}
	}()
}

type Track struct {
	stats *types.MonitoringStats
}

func NewTrack() *Track {
	return &Track{
		stats: &types.MonitoringStats{
			LastErrors: make([]types.ErrorLog, 0),
		},
	}
}

// TrackRequest отслеживание запроса
// ---------------------------------
func (t *Track) TrackRequest(duration time.Duration) {
	t.stats.Lock()
	defer t.stats.Unlock()

	t.stats.RequestCount++
	t.stats.TotalLatency += duration
}

// TrackError отслеживание ошибки
// ------------------------------
func (t *Track) TrackError(err error) {
	t.stats.Lock()
	defer t.stats.Unlock()
	t.stats.ErrorCount++

	err_log := types.ErrorLog{
		Timestamp: time.Now(),
		Error:     err.Error(),
	}

	// Ограничение количества последних ошибок до 10
	if len(t.stats.LastErrors) >= 10 {
		t.stats.LastErrors = t.stats.LastErrors[1:]
	}

	t.stats.LastErrors = append(t.stats.LastErrors, err_log)
}

// TrackDBQuery отслеживание запроса к БД
// --------------------------------------
func (t *Track) TrackDBQuery(duration time.Duration) {
	t.stats.Lock()
	defer t.stats.Unlock()

	t.stats.DBQueryCount++
	t.stats.DBTotalLatency += duration
}

// TrackDBError отслеживание ошибки БД
// -----------------------------------
func (t *Track) TrackDBError() {
	t.stats.Lock()
	defer t.stats.Unlock()

	t.stats.DBErrorCount++
}

// GetTrackStats получение статистики
// ----------------------------------
func (t *Track) GetTrackStats() types.MonitoringResponse {
	t.stats.Lock()
	defer t.stats.Unlock()

	return types.MonitoringResponse{
		Requests: struct {
			Total        int64   `json:"total"`
			Errors       int64   `json:"errors"`
			SuccessRate  float64 `json:"success_rate"`
			AvgLatencyMs float64 `json:"avg_latency_ms"`
		}{
			Total:  t.stats.RequestCount,
			Errors: t.stats.ErrorCount,
			SuccessRate: func(stats *types.MonitoringStats) float64 {
				if stats.RequestCount == 0 {
					return 0
				}
				successCount := stats.RequestCount - stats.ErrorCount
				successRate := float64(successCount) / float64(stats.RequestCount) * 100
				return math.Round(successRate*100) / 100
			}(t.stats),
			AvgLatencyMs: func(stats *types.MonitoringStats) float64 {
				if stats.RequestCount == 0 {
					return 0
				}
				avgLatency := float64(stats.TotalLatency.Milliseconds()) / float64(stats.RequestCount)
				return math.Round(avgLatency*100) / 100
			}(t.stats),
		},
		Database: struct {
			TotalQueries int64   `json:"total_queries"`
			Errors       int64   `json:"errors"`
			AvgLatencyMs float64 `json:"avg_latency_ms"`
		}{
			TotalQueries: t.stats.DBQueryCount,
			Errors:       t.stats.DBErrorCount,
			AvgLatencyMs: float64(t.stats.DBTotalLatency.Milliseconds()) / float64(t.stats.DBQueryCount),
		},
		LastErrors: t.stats.LastErrors,
	}
}
