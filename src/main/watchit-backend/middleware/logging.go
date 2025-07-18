package middleware

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"watchit/httpx/infra/logger"
)

// Logger - структура конструктора
// logger - логирования данных
// currentLogDir - получение/создание папки с датой example: (15.02/)
type Logger struct {
	logger        *logger.Logger
	currentLogDir string
}

// NewLogger - конструктор
func NewLogger(logger *logger.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int64
}

// Получение/создание папки с датой example: (15.02/)
func (l *Logger) getCurrentLogDir() string {
	dateFormatFolder := time.Now().Format("02.01")
	dir := filepath.Join(os.Getenv("LOG_DIR"), "logs", dateFormatFolder)
	if err := os.MkdirAll(dir, 0755); err != nil {
		l.logger.Error(err.Error())
	}
	return dir
}

// Обновление даты в папке
func (l *Logger) updateLogDirIfNeeded() {
	newLogDir := l.getCurrentLogDir()
	if newLogDir != l.currentLogDir {
		l.currentLogDir = newLogDir
	}
}

// Логирование в server.log данные о запросах к серверу
func (l *Logger) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter{
			ResponseWriter: w,
			status:         200,
		}

		next.ServeHTTP(wrapped, r)

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

		l.updateLogDirIfNeeded()
		l.productionLogging(entry)
	})
}

// Получение HTTP версии протокола
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

// Запись логов в папку logs/date/server.log
func (l *Logger) productionLogging(log string) {
	APP_DIR := filepath.Join(l.currentLogDir, "server.log")

	appLogFile, err := os.OpenFile(APP_DIR, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		l.logger.Error("failed to open app log file: %v", err)
		return
	}
	defer appLogFile.Close()

	_, err = appLogFile.WriteString(log)
	if err != nil {
		l.logger.Error("error writing to log file: %v", err)
		return
	}
}

// Добавляет статус ответа в обертку
func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

// Добавляет размер записи в обертку
func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += int64(size)
	return size, err
}
