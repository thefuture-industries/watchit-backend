package favourite

import (
	"context"
	"database/sql"
	"flicksfi/internal/types"
	"fmt"
	"time"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

// Проверка что избранное не существует в БД у uuid
// ------------------------------------------------
func (s *Service) CheckFavourites(uuid string, movieID int) error {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД
	rows, err := s.db.QueryContext(ctx, "select * from favourites where uuid = ? and movieId = ?", uuid, movieID)
	if err != nil {
		return fmt.Errorf("database request error")
	}

	// Сканирование данных
	favourite := new(types.Favourites)
	for rows.Next() {
		rows.Scan(&favourite.ID, &favourite.UUID, &favourite.MovieID, &favourite.MoviePoster, &favourite.CreatedAt)
	}
	
	if favourite.ID != 0 {
		return fmt.Errorf("such a movie is already in favorites")
	}

	return nil
}

// Запись избранного в БД
// ----------------------
func (s *Service) AddFavourite(favourite types.FavouriteAddPayload) error {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на создания пользователя
	_, err := s.db.ExecContext(ctx, "insert into favourites (uuid, movieId, moviePoster, createdAt) values (?, ?, ?, ?)", favourite.UUID, favourite.MovieID, favourite.MoviePoster, time.Now().Format("2006-01-02 15:04:05"))

	// Обработка ошибки
	if err != nil {
		return fmt.Errorf("database insert error")
	}

	return nil
}

// Получение списка избранных фильмов
// ----------------------------------
func (s *Service) Favourites(uuid string) ([]types.Favourites, error) {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, "select * from favourites where uuid = ?", uuid)
	if err != nil {
		return nil, fmt.Errorf("error execute query to db")
	}

	var favourites []types.Favourites
	for rows.Next() {
		var f types.Favourites
		err := rows.Scan(&f.ID, &f.UUID, &f.MovieID, &f.MoviePoster, &f.CreatedAt)
		if err != nil {
			return nil, err
		}

		favourites = append(favourites, f)
	}

	if len(favourites) == 0 {
		return nil, fmt.Errorf("you don't have any favourites!")
	}

	return favourites, nil
}

// Удаление избранного фильма
// --------------------------
func (s *Service) DeleteFavourite(payload types.FavouriteDeletePayload) error {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, "delete from favourites where uuid = ? and movieId = ?", payload.UUID, payload.MovieID)
	if err != nil {
		return fmt.Errorf("database delete error")
	}

	return nil
}
