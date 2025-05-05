package packages

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
)

type Logger struct {
	logger        *zap.Logger
	currentLogDir string
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{
		logger:        logger,
		currentLogDir: getCurrentLogDir(),
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int64
}

func getCurrentLogDir() string {
	dateFormatFolder := time.Now().Format("02.01")
	dir := filepath.Join("logs", dateFormatFolder)
	if err := os.MkdirAll(dir, 0755); err != nil {
		ErrorLog(err)
	}
	return dir
}

func (l *Logger) updateLogDirIfNeeded() {
	newLogDir := getCurrentLogDir()
	if newLogDir != l.currentLogDir {
		l.currentLogDir = newLogDir
	}
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

		l.updateLogDirIfNeeded()
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

// Запись логов в папку logs/date/app.log
// ---------------------------------------
func (l *Logger) productionLogging(log string) {
	APP_DIR := filepath.Join(l.currentLogDir, "app.log")

	// Открытие или создание файла
	appLogFile, err := os.OpenFile(APP_DIR, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		l.logger.Error("failed to open app log file: %v", zap.Error(err))
		return
	}
	defer appLogFile.Close()

	// Запись данных в файл
	_, err = appLogFile.WriteString(log)
	if err != nil {
		l.logger.Error("error writing to log file: %v", zap.Error(err))
		return // Важно выйти, если не смогли записать
	}
}

// WriteHeader добавляет статус ответа в обертку
// ---------------------------------------------
func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

// Write добавляет размер записи в обертку
// ---------------------------------------
func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += int64(size)
	return size, err
}

func ErrorLog(err error) {
	dateFormatFolder := time.Now().Format("02-01")
	APP_DIR := filepath.Join(getCurrentLogDir(), fmt.Sprintf("errors-%s.log", dateFormatFolder))

	file, fileErr := os.OpenFile(APP_DIR, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "", 0)
	currentTime := time.Now().Format("[02/Jan/2006:15:04:05 -0700]")

	fmt.Printf("%s %s\n", currentTime, err.Error())
	logger.Printf("%s \"%s\"\n", currentTime, err.Error())
}
