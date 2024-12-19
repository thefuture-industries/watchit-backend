package favourite

import (
	"context"
	"database/sql"
	"flicksfi/cmd/configuration"
	"flicksfi/internal/types"
	"fmt"
	"time"

	"go.uber.org/zap"
)

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

// Проверка что избранное не существует в БД у uuid
// ------------------------------------------------
func (s *Service) CheckFavourites(uuid string, movieID int) error {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД
	queryStart := time.Now()
	row := s.db.QueryRowContext(ctx, "select * from favourites where uuid = ? and movieId = ?", uuid, movieID)
	favourite := new(types.Favourites)

	// читаем из результата
	err := row.Scan(&favourite.ID, &favourite.UUID, &favourite.MovieID, &favourite.MoviePoster, &favourite.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("user_id", uuid),
			zap.Int("movie_id", movieID),
			zap.Error(err),
		)
		return fmt.Errorf("get favourite: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	return nil
}

// Запись избранного в БД
// ----------------------
func (s *Service) AddFavourite(favourite types.FavouriteAddPayload) error {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на создания пользователя
	queryStart := time.Now()
	_, err := s.db.ExecContext(ctx, "insert into favourites (uuid, movieId, moviePoster, createdAt) values (?, ?, ?, ?)", favourite.UUID, favourite.MovieID, favourite.MoviePoster, time.Now().Format("2006-01-02 15:04:05"))

	// Обработка ошибки
	if err != nil {
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", favourite.UUID),
			zap.Error(err))
		return fmt.Errorf("add favourite: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	return nil
}

// Получение списка избранных фильмов
// ----------------------------------
func (s *Service) Favourites(uuid string) ([]types.Favourites, error) {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД
	queryStart := time.Now()
	rows, err := s.db.QueryContext(ctx, "select * from favourites where uuid = ?", uuid)
	if err != nil {
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", uuid),
			zap.Error(err))
		return nil, fmt.Errorf("query favourites: %w", err)
	}
	defer rows.Close()

	// Сканирование данных
	var favourites []types.Favourites
	for rows.Next() {
		var f types.Favourites
		if err := rows.Scan(&f.ID, &f.UUID, &f.MovieID, &f.MoviePoster, &f.CreatedAt); err != nil {
			// Логирование ошибки
			s.monitor.TrackDBError()
			s.monitor.TrackError(err)
			s.logger.Error("database error",
				zap.String("uuid", uuid),
				zap.Error(err))
			return nil, fmt.Errorf("scan favourite: %w", err)
		}

		favourites = append(favourites, f)
	}

	if len(favourites) == 0 {
		return []types.Favourites{}, nil
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	return favourites, nil
}

// Удаление избранного фильма
// --------------------------
func (s *Service) DeleteFavourite(payload types.FavouriteDeletePayload) error {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД
	queryStart := time.Now()
	_, err := s.db.ExecContext(ctx, "delete from favourites where uuid = ? and movieId = ?", payload.UUID, payload.MovieID)
	if err != nil {
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", payload.UUID),
			zap.Error(err))
		return fmt.Errorf("delete favourite: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	return nil
}
