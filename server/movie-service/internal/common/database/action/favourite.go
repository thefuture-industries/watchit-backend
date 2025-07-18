package action

import (
	"go-movie-service/internal/common/database"
	"go-movie-service/internal/lib"
)

func CreateFavourite(uuid string, genreID uint16) error {
	logger := lib.NewLogger()
	db := database.GetDB()

	return nil
}
