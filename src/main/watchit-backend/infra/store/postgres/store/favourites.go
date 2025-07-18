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

func (s *FavouriteStore) Get_FavouritesByUuid(ctx context.Context, uuid string) (*[]models.Favourite, error) {
	favourites := []models.Favourite{}

	query := `
		SELECT id, user_uuid, movie_id, movie_poster FROM favourites WHERE user_uuid = $1
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		favourite := models.Favourite{}
	}
}
