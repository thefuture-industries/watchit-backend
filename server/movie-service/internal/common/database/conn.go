package database

import (
	"go-movie-service/internal/lib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB(dsn string) {
	loggerApp := lib.NewLogger()

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		loggerApp.Error(err.Error())
		return
	}

	// err = db.AutoMigrate(&Users{}, &Payments{})
	// if err != nil {
	// 	loggerApp.Error(err.Error())
	// 	return
	// }

	// loggerApp.Info("Migration completed successfully.")
	loggerApp.Info("Successfully connected to database!")
}

func GetDB() *gorm.DB {
	return db
}
