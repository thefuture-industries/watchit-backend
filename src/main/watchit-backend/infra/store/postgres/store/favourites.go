package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"watchit/httpx/infra/logger"
	"watchit/httpx/infra/store/postgres/models"
)

type FavouriteStore struct {
	db     *sql.DB
	logger *logger.Logger
}

func (s *FavouriteStore) Create_Favourite(ctx context.Context, favourite *models.Favourite) error {
	query := `
		INSERT INTO favourites (user_uuid, movie_id, movie_poster) VALUES ($1, $2, $3)
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, favourite.UserUUID, favourite.MovieId, favourite.MoviePoster)
	if err != nil {
		return err
	}

	return nil
}

func (s *FavouriteStore) Get_FavouritesByUuid(ctx context.Context, uuid string) (*[]models.Favourite, error) {
	favourites := []models.Favourite{}

	query := `
		SELECT id, user_uuid, movie_id, movie_poster FROM favourites WHERE user_uuid = $1
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		favourite := models.Favourite{}

		err := rows.Scan(
			&favourite.ID,
			&favourite.UserUUID,
			&favourite.MovieId,
			&favourite.MoviePoster,
		)
		if err != nil {
			return nil, err
		}

		favourites = append(favourites, favourite)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &favourites, nil
}

func (s *FavouriteStore) Get_FavouriteByUuidByMovieId(ctx context.Context, uuid string, movieId int) (*models.Favourite, error) {
	favourite := models.Favourite{}

	query := `
		SELECT id, user_uuid, movie_id, movie_poster FROM favourites WHERE movie_id = $1 AND user_uuid = $2
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, movieId, uuid).Scan(
		&favourite.ID,
		&favourite.UserUUID,
		&favourite.MovieId,
		&favourite.MoviePoster,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &favourite, nil
}
