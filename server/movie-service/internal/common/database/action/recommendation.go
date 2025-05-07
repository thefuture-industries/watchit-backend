package action

import (
	"fmt"
	"go-movie-service/internal/common/database"
	"go-movie-service/internal/common/database/schema"
	"go-movie-service/internal/lib"
)

func CreateRecommendation(uuid string, genreID uint) error {
	logger := lib.NewLogger()
	db := database.GetDB()

	var genres schema.Genres
	if err := db.Where("genre_id = ?", genreID).First(&genres).Error; err != nil {
		logger.Error(fmt.Sprintf("Database -> 'Failed to create recommendation: %v'", err))
		return fmt.Errorf("Failed to create recommendation.")
	}

	recomSchemaPrepare := schema.Recommendations{
		UUID:    uuid,
		GenreID: genres.ID,
		Count:   1,
	}

	if err := db.Create(&recomSchemaPrepare).Error; err != nil {
		logger.Error(fmt.Sprintf("Database -> 'Failed to create recommendation: %v'", err))
		return fmt.Errorf("Failed to create recommendation.")
	}

	return nil
}
