// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package actions

import (
	"fmt"
	"go-user-service/internal/common/database"
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
	result := database.GetDB().First(&user, id)

	if result.Error != nil {
		return nil, fmt.Errorf("the user was not found")
	}

	return &user, nil
}

func GetUserByEmail(email string) (*database.Users, error) {
	var user database.Users
	result := database.GetDB().First(&user, email)

	if result.Error != nil {
		return nil, fmt.Errorf("the user was not found")
	}

	return &user, nil
}

func GetUserByUUID(uuid string) (*database.Users, error) {
	var user database.Users
	result := database.GetDB().First(&user, uuid)

	if result.Error != nil {
		return nil, fmt.Errorf("the user was not found")
	}

	return &user, nil
}

func GetUserByUsername(username string) (*database.Users, error) {
	var user database.Users
	result := database.GetDB().First(&user, username)

	if result.Error != nil {
		return nil, fmt.Errorf("the user was not found")
	}

	return &user, nil
}
