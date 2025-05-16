package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Настройка и создание БД
// -----------------------
func NewPostgreSQLStorage(cfg string) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Настройка пула соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)

	return db, nil
}
