package recommendation

import (
	"context"
	"database/sql"
	"flicksfi/cmd/configuration"
	"flicksfi/internal/types"
	"flicksfi/pkg/movie"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"go.uber.org/zap"
)

type KeyValueGenre struct {
	Key   int
	Value int
}

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

// Получение рекомендаций пользователя
// -----------------------------------
func (s *Service) GetRecommendation(uuid string) ([]types.Recommendations, error) {
	// Мониторинг времени запроса
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на получения рекомендаций
	queryStart := time.Now()
	rows, err := s.db.QueryContext(ctx, "select * from recommendations where uuid = ? limit ?", uuid, 100)
	// Обработка ошибки
	if err != nil {
		// Мониторинг ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)

		// Логирование ошибки
		s.logger.Error("database error",
			zap.String("uuid", uuid),
			zap.Error(err))
		return nil, fmt.Errorf("query recommendations: %w", err)
	}
	defer rows.Close()

	// Мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	// Инициализация данных из БД
	var recoms []types.Recommendations
	for rows.Next() {
		var recom types.Recommendations

		if err := rows.Scan(&recom.ID, &recom.UUID, &recom.Title, &recom.Genre); err != nil {
			// Логирование ошибки
			s.monitor.TrackDBError()
			s.monitor.TrackError(err)
			s.logger.Error("scan recommendation",
				zap.String("uuid", uuid),
				zap.Error(err))
			return nil, fmt.Errorf("scan recommendation: %w", err)
		}

		recoms = append(recoms, recom)
	}

	if err = rows.Err(); err != nil {
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("iterate recommendations",
			zap.String("uuid", uuid),
			zap.Error(err))
		return nil, fmt.Errorf("iterate recommendations: %w", err)
	}

	return recoms, nil
}

// Запись рекомендаций пользователя
// --------------------------------
func (s *Service) AddRecommendation(recommendation types.RecommendationAddPayload) error {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на создания пользователя
	queryStart := time.Now()
	_, err := s.db.ExecContext(ctx, "insert into recommendations (uuid, title, genre) values (?, ?, ?)", recommendation.UUID, recommendation.Title, recommendation.Genre)

	// Обработка ошибки
	if err != nil {
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", recommendation.UUID),
			zap.Error(err))
		return fmt.Errorf("insert recommendation: %w", err)
	}

	// Мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	return nil
}

// Проверка на существования рекомендаций пользователя
func (s *Service) IsRecommendation(uuid, title string) (bool, error) {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на проверку существования рекомендаций
	queryStart := time.Now()
	row := s.db.QueryRowContext(ctx, "select * from recommendations where uuid = ? and title = ?", uuid, title)
	var recom types.Recommendations

	err := row.Scan(&recom.ID, &recom.UUID, &recom.Title, &recom.Genre)
	if err != nil {
		if err == sql.ErrNoRows {
			s.monitor.TrackDBQuery(time.Since(queryStart))
			return false, nil
		}

		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", uuid),
			zap.String("title", title),
			zap.Error(err))
		return false, fmt.Errorf("query recommendation: %w", err)
	}

	if recom.ID != 0 {
		return true, nil
	}

	// Мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	return false, nil
}

// Получение данных из данных таблицы
// ----------------------------------
func (s *Service) GetMovieRecommendations(recoms []types.Recommendations) ([]types.Movie, error) {
	genre_count := make(map[int]int)

	// Подсчет количества повторяющихся жанров
	for _, recom := range recoms {
		genreIDs, err := s.ArrayGenreIDS(recom.Genre)
		if err != nil {
			// Логирование ошибки
			s.logger.Error("array genre ids",
				zap.String("uuid", recom.UUID),
				zap.Error(err))
			return nil, fmt.Errorf("array genre ids: %w", err)
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
			// Логирование ошибки
			s.logger.Error("get movie by genre",
				zap.String("uuid", recoms[0].UUID),
				zap.String("genres", recoms[0].Title),
				zap.Error(err))
			return nil, fmt.Errorf("get movie by genre: %w", err)
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
			return nil, fmt.Errorf("array genre ids: %w", err)
		}

		result[i] = num
	}

	return result, nil
}
