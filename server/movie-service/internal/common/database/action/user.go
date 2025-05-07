package action

import (
	"errors"
	"fmt"
	"go-movie-service/internal/common/database"
	"go-movie-service/internal/common/database/schema"
	"go-movie-service/internal/lib"

	"gorm.io/gorm"
)

func GetUserByUUID(uuid string) (*schema.Users, error) {
	logger := lib.NewLogger()
	db := database.GetDB()

	var user schema.Users
	if err := db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		logger.Error(fmt.Sprintf("Database -> 'Error while fetching user: %v'", err))
		return nil, fmt.Errorf("Error while fetching user")
	}

	return &user, nil
}

func UpdateRecommendation(uuid string, genreID uint) (bool, error) {
	logger := lib.NewLogger()
	db := database.GetDB()

	var recommendation schema.Recommendations
	if err := db.Where("uuid = ? and genre_id = ?", uuid, genreID).First(&recommendation).Error; err != nil {
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
