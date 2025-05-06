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
	UUID          string `gorm:"size:255;not null;unique" json:"uuid"`
	SecretWord    string `gorm:"size:255;not null;unique" json:"secret_word"`
	Username      string `gorm:"size:100;not null" json:"username"`
	UsernameUpper string `gorm:"size:100;not null" json:"username_upper"`
	Email         string `gorm:"size:50;unique" json:"email"`
	EmailUpper    string `gorm:"size:50" json:"email_upper"`
	IPAddress     string `gorm:"size:40;not null" json:"ip_address"`
	Country       string `gorm:"size:70;not null" json:"country"`
	RegionName    string `gorm:"size:70;not null" json:"region_name"`
	ZIP           string `gorm:"size:40;not null" json:"zip"`
}

type Payments struct {
	gorm.Model
	UUID      string  `gorm:"unique;size:255;not null"`
	Email     string  `gorm:"size:100;not null"`
	Card      string  `gorm:"size:50;not null"`
	CardEnd   string  `gorm:"size:3;not null"`
	Total     float64 `gorm:"size:5;not null"`
	PaymentAt string  `gorm:"size:255;not null"`
	EndingAt  string  `gorm:"size:255;not null"`
}
