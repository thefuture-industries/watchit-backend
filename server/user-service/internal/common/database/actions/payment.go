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


func CreatePayment(payment *database.Payments) error {
	create := database.GetDB().Create(payment)

	if create.Error != nil {
		return fmt.Errorf("failed to create payment: %w", create.Error)
	}

	return nil
}

func GetPaymentsByUUID(uuid string) ([]*database.Payments, error) {
	var payments []*database.Payments
	result := database.GetDB().Where("uuid = ?", uuid).First(&payments)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf("an unexpected error occurred while getting the user")
	}

	return payments, nil
}
