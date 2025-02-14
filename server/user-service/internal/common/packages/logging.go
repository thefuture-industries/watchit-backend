// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package packages

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int64
}

func (l *Logger) LoggerMiddleware(next http.Handler) http.Handler {
	// возвращаем обработчик
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// получаем время начала запроса
		start := time.Now()

		// создаем обертку для ответа
		wrapped := &responseWriter{
			ResponseWriter: w,
			status:         200,
		}

		// вызываем следующий обработчик
		next.ServeHTTP(wrapped, r)

		// создаем строку с данными для записи в log
		entry := fmt.Sprintf("%s - - [%s] \"%s %s %s\" %d %d \"%s\" %v\n",
			r.RemoteAddr,
			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
			getHTTPVersion(r),
			r.Method,
			r.URL.Path,
			wrapped.status,
			wrapped.size,
			r.UserAgent(),
			time.Since(start),
		)

		l.productionLogging(entry)
	})
}

func getHTTPVersion(r *http.Request) string {
	switch r.ProtoMajor {
	case 1:
		return "HTTP/1.1"
	case 2:
		return "HTTP/2.0"
	case 3:
		return "HTTP/3.0"
	default:
		return fmt.Sprintf("HTTP/%d.%d", r.ProtoMajor, r.ProtoMinor)
	}
}

// Запись логов в папку logs/app.log
// ---------------------------------------
func (l *Logger) productionLogging(log string) {
	APP_DIR := filepath.Join("logs", "app")

	// Проверка на существование и создание папки
	err := os.MkdirAll(APP_DIR, 0755)
	if err != nil {
		l.logger.Error("error creating app directory: %v", zap.Error(err))
		return // Важно выйти, если папка не создана
	}

	// Открытие или создание файла
	appLogFile, err := os.OpenFile(filepath.Join(APP_DIR, "app.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		l.logger.Error("failed to open app log file: %v", zap.Error(err))
		return // Важно выйти, если файл не открыт
	}
	defer appLogFile.Close()

	// Запись данных в файл
	_, err = appLogFile.WriteString(log)
	if err != nil {
		l.logger.Error("error writing to log file: %v", zap.Error(err))
		return // Важно выйти, если не смогли записать
	}
}
