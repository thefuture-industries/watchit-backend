package database

import (
	"go-user-service/internal/lib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnectDB(dsn string) {
	loggerApp := lib.NewLogger()

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		loggerApp.Error(err.Error())
		return
	}

	err = db.AutoMigrate(&Users{}, &Payments{})
	if err != nil {
		loggerApp.Error(err.Error())
		return
	}

	loggerApp.Info("Migration completed successfully.")
	loggerApp.Info("Successfully connected to database!")
}

func GetDB() *gorm.DB {
	return db
}
