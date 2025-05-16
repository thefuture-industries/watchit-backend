package main

import (
	"context"
	"database/sql"
	"flicksfi/cmd/api"
	"flicksfi/cmd/configuration"
	"flicksfi/internal/config"
	"flicksfi/internal/db"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

// @title flicksfi API
// @version 0.1.2
// @description API service for working with the flicksfi application for fast movie and video search. tg [https://t.me/flicksfi]

// @host localhost:8080
// @BasePath /api/v1

func main() {
	// Инициализация логгера
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	// ----------------------------
	// Найстрока и подключение к бд
	// ----------------------------
	db, err := db.NewPostgreSQLStorage("user=" + config.Envs.DBUser + " password=" + config.Envs.DBPassword + " host=" + config.Envs.DBHost + " dbname=" + config.Envs.DBName + " sslmode=disable")

	// Проверка ошибки БД
	if err != nil {
		log.Fatal(err)
	}

	// Проверка открытия БД
	initStorage(db)

	// Создаем каналы для сигналов и ошибок
	signals := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Конфигурация приложения (метрики и мониторинг)
	service := configuration.NewService(db, logger)
	service.StartDBMonitoring(30 * time.Second)

	// Запуск метрик сервера
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":2121", nil); err != nil {
			logger.Error("metrics server error", zap.Error(err))
		}
	}()

	// Создаем и запускаем основной сервер
	server := api.NewAPIServer(":8080", db)
	go func() {
		fmt.Println(`    _______      __        _____    ___    ____  ____
   / __/ (_)____/ /_______/ __(_)  /   |  / __ \/  _/
  / /_/ / / ___/ //_/ ___/ /_/ /  / /| | / /_/ // /
 / __/ / / /__/ ,< (__  ) __/ /  / ___ |/ ____// /
/_/ /_/_/\___/_/|_/____/_/ /_/  /_/  |_/_/   /___/
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

		// Graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Сначала останавливаем основной сервер
		if err := server.Shutdown(ctx); err != nil {
			logger.Error("error shutting down server", zap.Error(err))
		}

		// Затем останавливаем сервис мониторинга
		if err := service.Shutdown(ctx); err != nil {
			logger.Error("error during service shutdown", zap.Error(err))
			os.Exit(1)
		}

		logger.Info("shutdown completed")
	}
}

// --------------------
// Проверка открытия БД
// --------------------
func initStorage(db *sql.DB) {
	// --------------------------------------------
	// Если есть ошибка с БД, то вывести сообщение
	// ---------------------------------------------
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// --------------------------
	// Если все ок и БД открылась
	// --------------------------
	log.Println("DB: Successfully initialized!")
}
