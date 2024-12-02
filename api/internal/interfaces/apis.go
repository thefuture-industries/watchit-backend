package interfaces

import "flick_finder/internal/types"

type IApis interface {
	// Из string массива жанров в массив цифр
	ArrayGenreIDS(str string) ([]int, error)

	// Получение лимитов из БД по UserID
	GetLimitText(uuid string) (types.Limiter, error)

	// Уменьшение лимита текста из БД по UserID
	ReducingLimitText(uuid string) error
}
