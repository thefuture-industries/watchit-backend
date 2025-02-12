// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package database

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	UUID      string `gorm:"unique;size:255;not null"`
	Username  string `gorm:"unique;size:20;not null"`
	Email     string `gorm:"unique;size:100;not null"`
	IPAddress string `gorm:"size:15;not null"`
	Country   string `gorm:"size:50;not null"`
	CreatedAt string `gorm:"size:255;not null"`
}

type Payments struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	UUID      string  `gorm:"unique;size:255;not null"`
	Email     string  `gorm:"size:100;not null"`
	Card      string  `gorm:"size:50;not null"`
	CardEnd   string  `gorm:"size:3;not null"`
	Total     float64 `gorm:"size:5;not null"`
	PaymentAt string  `gorm:"size:255;not null"`
	EndingAt  string  `gorm:"size:255;not null"`
}
