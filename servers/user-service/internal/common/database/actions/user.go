// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package actions

import (
	"errors"
	"fmt"
	"go-user-service/internal/common/database"

	"gorm.io/gorm"
)

func CreateUser(user *database.Users) error {
	create := database.GetDB().Create(user)

	if create.Error != nil {
		return fmt.Errorf("failed to create user: %w", create.Error)
	}

	return nil
}

func GetUserByID(id uint) (*database.Users, error) {
	var user database.Users
	result := database.GetDB().Where("id = ?", id).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf("an unexpected error occurred while getting the user")
	}

	return &user, nil
}

func GetUserByEmail(email string) (*database.Users, error) {
	var user database.Users
	result := database.GetDB().Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf("an unexpected error occurred while getting the user")
	}

	return &user, nil
}

func GetUserByUUID(uuid string) (*database.Users, error) {
	var user database.Users
	result := database.GetDB().Where("uuid = ?", uuid).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf("an unexpected error occurred while getting the user")
	}

	return &user, nil
}

func GetUserByUsername(username string) (*database.Users, error) {
	var user database.Users
	result := database.GetDB().Where("username = ?", username).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf("an unexpected error occurred while getting the user")
	}

	return &user, nil
}
