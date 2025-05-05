// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package actions

import (
	"errors"
	"fmt"
	"go-user-service/internal/common/database"
	"go-user-service/internal/lib"

	"gorm.io/gorm"
)

func CreateUser(user *database.Users) error {
	logger := lib.NewLogger()
	db := database.GetDB()

	if err := db.Create(user).Error; err != nil {
		logger.Error(fmt.Sprintf("Database -> 'Failed to create user: %v'", err))
		return fmt.Errorf("Failed to create user.")
	}

	return nil
}

func GetUserByID(id uint) (*database.Users, error) {
	logger := lib.NewLogger()
	db := database.GetDB()

	var user database.Users
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		logger.Error(fmt.Sprintf("Database -> 'Error while fetching user: %v'", err))
		return nil, fmt.Errorf("Error while fetching user")
	}

	return &user, nil
}

func GetUserByEmail(email string) (*database.Users, error) {
	logger := lib.NewLogger()
	db := database.GetDB()

	var user database.Users
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		logger.Error(fmt.Sprintf("Database -> 'Error while fetching user: %v'", err))
		return nil, fmt.Errorf("Error while fetching user")
	}

	return &user, nil
}

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

func GetUserByUsername(username string) (*database.Users, error) {
	logger := lib.NewLogger()
	db := database.GetDB()

	var user database.Users
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		logger.Error(fmt.Sprintf("Database -> 'Error while fetching user: %v'", err))
		return nil, fmt.Errorf("Error while fetching user")
	}

	return &user, nil
}
