// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package main

import (
	"fmt"
	"go-movie-service/cmd/api"
	"go-movie-service/internal/common/database"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	godotenv.Load()

	// Инициализация логгера
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	// ----------------------------
	// Найстрока и подключение к бд
	// ----------------------------
	database.ConnectDB(os.Getenv("DSN"))
	db := database.GetDB()

	// Создаем каналы для сигналов и ошибок
	signals := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Создаем и запускаем основной сервер
	server := api.NewAPIServer(":8011", db)
	go func() {
		fmt.Println(`    __  ___           _         __  ____                                      _
   /  |/  /___ _   __(_)__     /  |/  (_)_____________  ________  ______   __(_)_______
  / /|_/ / __ \ | / / / _ \   / /|_/ / / ___/ ___/ __ \/ ___/ _ \/ ___/ | / / / ___/ _ \
 / /  / / /_/ / |/ / /  __/  / /  / / / /__/ /  / /_/ (__  )  __/ /   | |/ / / /__/  __/
/_/  /_/\____/|___/_/\___/  /_/  /_/_/\___/_/   \____/____/\___/_/    |___/_/\___/\___/
                                                                                        `)
		if err := server.Run(); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		logger.Error("server error", zap.Error(err))
		os.Exit(1)
	case sig := <-signals:
		logger.Info("received shutdown signal", zap.String("signal", sig.String()))

		logger.Info("shutdown completed")
	}
}
