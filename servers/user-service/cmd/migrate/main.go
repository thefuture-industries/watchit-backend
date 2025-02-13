// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package main

import (
	"go-user-service/internal/common/database"
	"log"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// --------------------------------
// Функция создания миграции для БД
// --------------------------------
func main() {
	godotenv.Load()

	database.ConnectDB(os.Getenv("DSN"))
	db := database.GetDB()

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20230101_create_tables",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&database.Users{}, &database.Payments{})
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("users"); err != nil {
					return err
				}

				if err := tx.Migrator().DropTable("payments"); err != nil {
					return err
				}

				return nil
			},
		},
	})
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}
