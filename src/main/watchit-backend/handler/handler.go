package handler

import (
	"database/sql"
	"watchit/httpx/infra/logger"
)

type BaseHandler struct {
	Db     *sql.DB
	Logger *logger.Logger
	Store  store.Storage
}
