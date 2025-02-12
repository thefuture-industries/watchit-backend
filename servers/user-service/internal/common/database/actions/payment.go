// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package actions

import (
	"fmt"
	"go-user-service/internal/common/database"
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
	result := database.GetDB().Find(&payments, uuid)

	if result.Error != nil {
		return nil, fmt.Errorf("the payment was not found")
	}

	return payments, nil
}
