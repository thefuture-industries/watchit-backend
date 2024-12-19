package configuration

import (
	"context"

	"go.uber.org/zap"
)

// Shutdown gracefully закрытие сервисов
// -------------------------------------
func (s *Service) Shutdown(ctx context.Context) error {
	// Логирование начала завершения работы
	s.logger.Info("starting graceful shutdown")

	// Закрытие каналов, если они были открыты
	if s.done != nil {
		close(s.done)
	}

	// Ожидание завершения активных операций
	done := make(chan struct{})
	go func() {
		s.logger.Info("waiting for active operations to complete")
		close(done)
	}()

	// Ожидание завершения работы
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	}

	// Закрытие базы данных
	if err := s.db.Close(); err != nil {
		s.logger.Error("error closing database connection", zap.Error(err))
		return err
	}

	// Логирование завершения завершения работы
	s.logger.Info("graceful shutdown completed")
	return nil
}
