// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package main

import (
	"fmt"
	"go-user-service/cmd/api"
	"go-user-service/internal/common/database"
	"go-user-service/internal/packages"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Test struct {
	test string
}

func main() {
	if err := godotenv.Load(); err != nil {
		packages.ErrorLog(err)
	}

	// Инициализация логгера
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	if err := logger.Sync(); err != nil {
		packages.ErrorLog(err)
	}

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
	server := api.NewAPIServer(":8001", db)
	go func() {
		fmt.Println("\n" + `██╗   ██╗███████╗███████╗██████╗       ███╗   ███╗██╗ ██████╗██████╗  ██████╗      ██████╗  ██████╗
██║   ██║██╔════╝██╔════╝██╔══██╗      ████╗ ████║██║██╔════╝██╔══██╗██╔═══██╗    ██╔════╝ ██╔═══██╗
██║   ██║███████╗█████╗  ██████╔╝█████╗██╔████╔██║██║██║     ██████╔╝██║   ██║    ██║  ███╗██║   ██║
██║   ██║╚════██║██╔══╝  ██╔══██╗╚════╝██║╚██╔╝██║██║██║     ██╔══██╗██║   ██║    ██║   ██║██║   ██║
╚██████╔╝███████║███████╗██║  ██║      ██║ ╚═╝ ██║██║╚██████╗██║  ██║╚██████╔╝    ╚██████╔╝╚██████╔╝
 ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝      ╚═╝     ╚═╝╚═╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝      ╚═════╝  ╚═════╝
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
