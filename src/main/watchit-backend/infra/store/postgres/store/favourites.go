package store

import (
	"context"
	"database/sql"
	"watchit/httpx/infra/logger"
)

type FavouriteStore struct {
	db     *sql.DB
	logger *logger.Logger
}

func (s *FavouriteStore) Get_FavouritesByUuid(ctx context.Context, uuid string) ()
