package interfaces

import "flicksfi/internal/types"

type ILimiter interface {
	// Получение лимитов из БД по UserID
	GetLimits(uuid string) (*types.Limiter, error)

	// Уменьшение лимита текста из БД по UserID
	ReducingLimitText(uuid string) error

	// Обновление лимитов после 24 часов
	UpdateLimits(uuid string) error
}
