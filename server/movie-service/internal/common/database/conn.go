package database

import (
	"go-movie-service/internal/common/constants"
	"go-movie-service/internal/lib"

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

	err = db.AutoMigrate(&Recommendations{}, &Genres{})
	if err != nil {
		loggerApp.Error(err.Error())
		return
	}

	for name, index := range constants.GENRES {
		genre := Genres{
			GenreID:   index,
			GenreName: name,
		}

		db.FirstOrCreate(&genre, Genres{GenreID: index})
	}

	loggerApp.Info("Migration completed successfully.")
	loggerApp.Info("Successfully connected to database!")
}

func GetDB() *gorm.DB {
	return db
}
