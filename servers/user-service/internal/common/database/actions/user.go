// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package actions

import (
	"fmt"
	"go-user-service/internal/common/database"
)

func CreateUser(user *database.User) error {
	create := database.GetDB().Create(user)

	if create.Error != nil {
		return fmt.Errorf("failed to create user: %w", create.Error)
	}

	return nil
}

func GetUserByID(id uint) (*database.User, error) {
	var user database.User
	result := database.GetDB().First(&user, id)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", result.Error)
	}
	return &user, nil
}
