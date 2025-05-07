package main

import (
	"fmt"
	"go-movie-service/cmd/api"
	"go-movie-service/cmd/system"
	"go-movie-service/internal/common/database"
	"go-movie-service/internal/lib"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	loggerApp := lib.NewLogger()

	if err := godotenv.Load(); err != nil {
		loggerApp.Error(err.Error())
	}

	// Инициализация логгера
	logger, err := zap.NewProduction()
	if err != nil {
		loggerApp.Error(err.Error())
	}
	if err := logger.Sync(); err != nil {
		loggerApp.Error(err.Error())
	}

	// Найстрока и подключение к бд
	database.ConnectDB(os.Getenv("DSN"))
	db := database.GetDB()

	// Конфигурация приложения (метрики и мониторинг)
	system := system.NewSystem(db)
	system.StartDBMonitoring()

	// Создаем каналы для сигналов и ошибок
	signals := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Создаем и запускаем основной сервер
	server := api.NewAPIServer(":8011", db)
	go func() {
		loggerApp.Info("Swagger listen :8011/micro/movie/adm/doc")
		fmt.Println("\n" + `███╗   ███╗ ██████╗ ██╗   ██╗██╗███████╗    ███╗   ███╗██╗ ██████╗██████╗  ██████╗      ██████╗  ██████╗
████╗ ████║██╔═══██╗██║   ██║██║██╔════╝    ████╗ ████║██║██╔════╝██╔══██╗██╔═══██╗    ██╔════╝ ██╔═══██╗
██╔████╔██║██║   ██║██║   ██║██║█████╗█████╗██╔████╔██║██║██║     ██████╔╝██║   ██║    ██║  ███╗██║   ██║
██║╚██╔╝██║██║   ██║╚██╗ ██╔╝██║██╔══╝╚════╝██║╚██╔╝██║██║██║     ██╔══██╗██║   ██║    ██║   ██║██║   ██║
██║ ╚═╝ ██║╚██████╔╝ ╚████╔╝ ██║███████╗    ██║ ╚═╝ ██║██║╚██████╗██║  ██║╚██████╔╝    ╚██████╔╝╚██████╔╝
╚═╝     ╚═╝ ╚═════╝   ╚═══╝  ╚═╝╚══════╝    ╚═╝     ╚═╝╚═╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝      ╚═════╝  ╚═════╝
                                                                                                         `)
		if err := server.Run(); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		loggerApp.Error(fmt.Sprintf("server error: %s", err.Error()))
		os.Exit(1)
	case sig := <-signals:
		loggerApp.Warning(fmt.Sprintf("received shutdown signal: %s", sig.String()))
		loggerApp.Warning("shutdown completed")
	}
}
