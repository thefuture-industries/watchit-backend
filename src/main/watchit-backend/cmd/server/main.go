package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"watchit/httpx/infra/logger"
	"watchit/httpx/infra/store/postgres"
	"watchit/httpx/infra/store/postgres/store"
)

type httpServer struct {
	db     *sql.DB
	logger *logger.Logger
	store  store.Storage
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	log := logger.NewLogger()

	// connection db
	db, err := postgres.New(os.Getenv("DB_ADDR"), 50, 10, "3m")
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer db.Close()
	fmt.Printf("[INFO] Successfully connected to database\n")

	store := store.NewStorage(db, log)

	server := &httpServer{
		db:     db,
		logger: log,
		store:  store,
	}

	// start http server
	if err := server.httpStart(); err != nil {
		log.Error(err.Error())
	}
}
