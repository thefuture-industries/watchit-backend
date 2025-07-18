package handler

import (
	"database/sql"
	"watchit/httpx/infra/logger"
	"watchit/httpx/infra/store/postgres/store"
)

type BaseHandler struct {
	Db     *sql.DB
	Logger *logger.Logger
	Store  store.Storage
}
