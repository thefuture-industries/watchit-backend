package main

import (
	"database/sql"
	"flick_finder/cmd/api"
	"flick_finder/internal/config"
	"flick_finder/internal/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// ----------------------------
	// Найстрока и подключение к бд
	// ----------------------------
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	// Проверка ошибки БД
	if err != nil {
		log.Fatal(err)
	}

	// Проверка открытия БД
	initStorage(db)

	// Слушаем сервер на порте 8080
	server := api.NewAPIServer(":8080", db)

	// Обработчик ошибок server api
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// --------------------
// Проверка открытия БД
// --------------------
func initStorage(db *sql.DB) {
	// --------------------------------------------
	// Если есть ошибка с БД, то вывести сообщениеdb
	// ---------------------------------------------
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// --------------------------
	// Если все ок и БД открылась
	// --------------------------
	log.Println("DB: Successfully initialized!")
}
