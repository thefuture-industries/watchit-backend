package admin

import (
	"database/sql"
	"flicksfi/cmd/configuration"
	"flicksfi/internal/types"
)

type Service struct {
	db    *sql.DB
	track *configuration.Track
}

func NewService(db *sql.DB, track *configuration.Track) *Service {
	return &Service{
		db:    db,
		track: track,
	}
}

// Получение мониторинга приложения
// --------------------------------
func (s *Service) GetMonitoring() types.MonitoringResponse {
	return s.track.GetTrackStats()
}
