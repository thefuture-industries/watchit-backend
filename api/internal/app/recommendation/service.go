package recommendation

import (
	"context"
	"database/sql"
	"flicksfi/internal/types"
	"flicksfi/pkg/movie"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type KeyValueGenre struct {
	Key   int
	Value int
}

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

// Получение рекомендаций пользователя
// -----------------------------------
func (s *Service) GetRecommendation(uuid string) ([]types.Recommendations, error) {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на создания пользователя
	rows, err := s.db.QueryContext(ctx, "select * from recommendations where uuid = ?", uuid)

	// Обработка ошибки
	if err != nil {
		return nil, fmt.Errorf("database insert error")
	}

	// Инициализация данных из БД
	var recoms []types.Recommendations
	for rows.Next() {
		var recom types.Recommendations

		err := rows.Scan(&recom.ID, &recom.UUID, &recom.Title, &recom.Genre)
		if err != nil {
			return nil, err
		}

		recoms = append(recoms, recom)
	}

	return recoms, nil
}

// Запись рекомендаций пользователя
// --------------------------------
func (s *Service) AddRecommendation(recommendation types.RecommendationAddPayload) error {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на создания пользователя
	_, err := s.db.ExecContext(ctx, "insert into recommendations (uuid, title, genre) values (?, ?, ?)", recommendation.UUID, recommendation.Title, recommendation.Genre)

	// Обработка ошибки
	if err != nil {
		return fmt.Errorf("database insert error")
	}

	return nil
}

// Получение данных из данных таблицы
// ----------------------------------
func (s *Service) GetMovieRecommendations(recoms []types.Recommendations) ([]types.Movie, error) {
	genre_count := make(map[int]int)

	// Подсчет количества повторяющихся жанров
	for _, recom := range recoms {
		genreIDs, err := s.ArrayGenreIDS(recom.Genre)
		if err != nil {
			return nil, err
		}

		for _, genreID := range genreIDs {
			genre_count[genreID]++
		}
	}

	// Создаем слайс для сортировки
	var sorted []KeyValueGenre
	for k, v := range genre_count {
		sorted = append(sorted, KeyValueGenre{k, v})
	}

	// Сортируем по значениям (по убыванию)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Value == sorted[j].Value {
			return sorted[i].Key < sorted[j].Key
		}
		return sorted[i].Value > sorted[j].Value
	})

	// Распределение количества жанров
	allocations := allocateWeight(sorted, 50)

	// Получение фильмов по жанрам
	var movieResponse []types.Movie
	for genre, count := range allocations {
		movies, err := movie.GetMovieByGenre(genre, count)
		if err != nil {
			return nil, err
		}

		movieResponse = append(movieResponse, movies...)
	}

	return movieResponse, nil
}

// Высчитывание размера и веса массива
// -----------------------------------
func allocateWeight(genre_count []KeyValueGenre, limit int) map[int]int {
	total_count := 0

	// Сначала подсчитаем общую сумму значений
	for _, item := range genre_count {
		total_count += item.Value
	}

	// Вычисление весов и начальное распределение
	allocations := make(map[int]int)
	remaining := limit

	// распределяем пропорционально
	for _, item := range genre_count {
		share := float64(item.Value) / float64(total_count) * float64(limit)
		allocation := int(share)

		// Гарантируем минимум 1 элемент для каждого жанра
		if allocation < 1 {
			allocation = 1
		}

		allocations[item.Key] = allocation
		remaining -= allocation
	}

	// распределяем оставшиеся элементы
	i := 0
	for remaining != 0 {
		if i >= len(genre_count) {
			i = 0
		}

		if remaining > 0 {
			allocations[genre_count[i].Key]++
			remaining--
		} else if remaining < 0 {
			if allocations[genre_count[i].Key] > 1 {
				allocations[genre_count[i].Key]--
				remaining++
			}
		}
		i++
	}

	return allocations
}

// Парсинг жанров массива в [int]
// ------------------------------
func (s *Service) ArrayGenreIDS(str string) ([]int, error) {
	parts := strings.FieldsFunc(str, func(r rune) bool {
		return unicode.IsSpace(r) || r == ','
	})

	result := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}

		result[i] = num
	}

	return result, nil
}
