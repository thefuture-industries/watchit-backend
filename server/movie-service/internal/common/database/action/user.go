package action

import (
	"errors"
	"fmt"
	"go-movie-service/internal/common/database"
	"go-movie-service/internal/lib"

	"gorm.io/gorm"
)

func GetUserByUUID(uuid string) (*database.Users, error) {
	logger := lib.NewLogger()
	db := database.GetDB()

	var user database.Users
	if err := db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		logger.Error(fmt.Sprintf("Database -> 'Error while fetching user: %v'", err))
		return nil, fmt.Errorf("Error while fetching user")
	}

	return &user, nil
}
