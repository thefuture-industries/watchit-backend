// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package database

import (
	"fmt"
	"go-user-service/internal/common/packages"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB(dsn string) {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		packages.ErrorLog(err)
		return
	}

	err = db.AutoMigrate(&Users{}, &Payments{})
	if err != nil {
		log.Fatalf("Failed to migrate to database: %v", err)
		packages.ErrorLog(err)
		return
	}

	fmt.Println("Migration completed successfully.")
	fmt.Println("Successfully connected to database!")
}

func GetDB() *gorm.DB {
	return db
}
