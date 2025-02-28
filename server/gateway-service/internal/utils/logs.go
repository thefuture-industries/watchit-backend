package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogRequestToFile(clientIP, method, url, status, message string) {
	err := os.MkdirAll("logs/app", os.ModePerm)
	if err != nil {
		// Если не удалось создать директорию, выводим ошибку и завершаем программу
		log.Println("Error creating logs directory:", err)
		return
	}

	// Открываем или создаём файл для записи
	logFile, err := os.OpenFile("logs/app/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// Логгер, настроенный на запись в файл
	logger := log.New(logFile, "", log.LstdFlags)

	// Записываем лог с текущим временем
	logMessage := fmt.Sprintf(
		"[%-20s] Device: %-15s | Method: %-5s | URL: %-30s | Status: %-5s | Message: %-50s",
		time.Now().Format("2006-01-02 15:04:05"), clientIP, method, url, status, message,
	)
	logger.Println(logMessage)
}
