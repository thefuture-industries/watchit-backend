package interfaces

import "flicksfi/internal/types"

type IAdmin interface {
	// Получение мониторинга приложения
	GetMonitoring() types.MonitoringResponse
}
