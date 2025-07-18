package store

import (
	"database/sql"
	"watchit/httpx/infra/logger"
)

type FavouriteStore struct {
	db     *sql.DB
	logger *logger.Logger
}
