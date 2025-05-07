package action

import (
	"errors"
	"fmt"
	"go-movie-service/internal/common/database"
	"go-movie-service/internal/common/database/schema"
	"go-movie-service/internal/lib"

	"gorm.io/gorm"
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

func UpdateRecommendation(uuid string, genreID uint) (bool, error) {
	logger := lib.NewLogger()
	db := database.GetDB()

	var genres schema.Genres
	if err := db.Where("genre_id = ?", genreID).First(&genres).Error; err != nil {
		logger.Error(fmt.Sprintf("Database -> 'Failed to create recommendation: %v'", err))
		return false, fmt.Errorf("Failed to create recommendation.")
	}

	var recommendation schema.Recommendations
	if err := db.Where("uuid = ? and genre_id = ?", uuid, genres.ID).First(&recommendation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		logger.Error(fmt.Sprintf("Database -> 'Failed to query recommendation: %v'", err))
		return false, fmt.Errorf("Failed to query recommendation")
	}

	recommendation.Count++
	if err := db.Save(&recommendation).Error; err != nil {
		logger.Error(fmt.Sprintf("Database -> 'Failed to query recommendation: %v'", err))
		return false, fmt.Errorf("Failed to query recommendation")
	}

	return true, nil
}
