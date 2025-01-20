package main

import (
	"flicksfi/internal/config"
	"log"
	"os"

	"flicksfi/internal/db"

	"github.com/golang-migrate/migrate/v4"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// --------------------------------
// Функция создания миграции для БД
// --------------------------------
func main() {
	// ----------------------
	// Найстрока и подключение к бд
	// ----------------------
	db, err := db.NewPostgreSQLStorage("user=" + config.Envs.DBUser + " password=" + config.Envs.DBPassword + " host=" + config.Envs.DBHost + " dbname=" + config.Envs.DBName + " sslmode=disable")

	// Вывод ошибки с БД
	if err != nil {
		log.Fatal(err)
	}

	// Драйвер для миграции с MySQL
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Тут происходит определение куда сохранять файлы миграции
	// cmd/migrate/migrations
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	// Обработка ошибки миграции
	if err != nil {
		log.Fatal(err)
	}

	// Создание и обработка миграции файла
	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
