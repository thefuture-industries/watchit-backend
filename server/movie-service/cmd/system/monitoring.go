package system

import (
	"fmt"
	"go-movie-service/internal/lib"
	"time"

	"gorm.io/gorm"
)

type System struct {
	db     *gorm.DB
	logger *lib.Logger
}

func NewSystem(db *gorm.DB) *System {
	return &System{
		db:     db,
		logger: lib.NewLogger(),
	}
}

func (s *System) StartDBMonitoring() error {
	done := make(chan struct{})

	dbSql, err := s.db.DB()
	if err != nil {
		return err
	}

	go func() {
		ticker := time.NewTicker(60 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				stats := dbSql.Stats()

				location, err := time.LoadLocation("Asia/Yekaterinburg")
				if err != nil {
					return
				}

				s.logger.System("================================")
				now := time.Now().In(location)
				s.logger.System(fmt.Sprintf("\033[94m%02d.%02d.%d %02d:%02d:%02d\033[0m",
					now.Day(), now.Month(), now.Year(),
					now.Hour(), now.Minute(), now.Second()))
				s.logger.System("================================")
				s.logger.System(fmt.Sprintf("'Database' -> OpenConnection: %d", stats.OpenConnections))
				s.logger.System(fmt.Sprintf("'Database' -> InUse: %d", stats.InUse))
				s.logger.System(fmt.Sprintf("'Database' -> Idle: %d", stats.Idle))
				s.logger.System(fmt.Sprintf("'Database' -> WaitCount: %d", stats.WaitCount))
				s.logger.System(fmt.Sprintf("'Database' -> WaitDuration: %d", stats.WaitDuration))
				s.logger.System(fmt.Sprintf("'Database' -> WaitDuration: %d", stats.WaitDuration))
				s.logger.System(fmt.Sprintf("'Database' -> MaxLifetimeClosed: %d", stats.MaxLifetimeClosed))

			case <-done:
				return
			}
		}
	}()

	return nil
}
