package store

import (
	"database/sql"
	"watchit/httpx/infra/logger"
)

type Storage struct {
	Users interface {
	}
}

func NewStorage(db *sql.DB, logger *logger.Logger) Storage {
	return Storage{}
}
