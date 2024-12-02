package main

import (
	"flick_finder/internal/config"
	"log"
	"os"

	"flick_finder/internal/db"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"

	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// --------------------------------
// Функция создания миграции для БД
// --------------------------------
func main() {
	// ----------------------
	// Найстрока и подключение к бд
	// ----------------------
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	// Вывод ошибки с БД
	if err != nil {
		log.Fatal(err)
	}

	// Драйвер для миграции с MySQL
	driver, err := mysql.WithInstance(db, &mysql.Config{})
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
