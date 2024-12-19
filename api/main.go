package main

import (
	"context"
	"database/sql"
	"flicksfi/cmd/api"
	"flicksfi/cmd/configuration"
	"flicksfi/internal/config"
	"flicksfi/internal/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

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
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

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
		logger.Info("starting main server on :8080")
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
	// Если есть ошибка с БД, то вывести сообщениеdb
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

// # Искусственно создаем нагрузку
// hey -n 1000 -c 10 http://localhost:8080/api/v1/health (-- golang)
// autocannon -c 10 -d 10 http://localhost:8080/api/v1/health (-- node.js)
