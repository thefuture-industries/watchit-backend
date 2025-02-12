// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID      string `gorm:"unique"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	IPAddress string
	Country   string
	CreatedAt string
}

type Payments struct {
	gorm.Model
	UUID      string
	Email     string
	Card      string
	CardEnd   string
	Total     float64
	PaymentAt string
	EndingAt  string
}
