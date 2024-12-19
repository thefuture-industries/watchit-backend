package interfaces

import "flicksfi/internal/types"

// Запись рекомендаций пользователя
type IRecommendation interface {
	// Получение рекомендаций пользователя
	GetRecommendation(uuid string) ([]types.Recommendations, error)

	// Запись рекомендаций пользователя
	AddRecommendation(recommendation types.RecommendationAddPayload) error

	// Проверка на существования рекомендаций пользователя
	IsRecommendation(uuid, title string) (bool, error)

	// Получение данных из данных таблицы
	GetMovieRecommendations(recoms []types.Recommendations) ([]types.Movie, error)

	// Парсинг жанров массива в [int]
	ArrayGenreIDS(str string) ([]int, error)
}
