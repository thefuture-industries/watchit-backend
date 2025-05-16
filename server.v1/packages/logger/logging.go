package logger

import (
	"context"
	env "flicksfi/internal/config"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
	client *s3.Client
	mu     sync.Mutex
}

func NewLogger(logger *zap.Logger) *Logger {
	client, err := createS3Client()
	if err != nil {
		logger.Error("Failed to create s3 client", zap.Error(err))
		return nil
	}

	return &Logger{
		logger: logger,
		client: client,
		mu:     sync.Mutex{},
	}
}

const (
	bucketName = "flicksfi-logs"
	fileName   = "flicksfi-monitoring.log"
	bufferSize = 173
)

var buffer = make([]string, 0, bufferSize)

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int64
}

// flush отправляет накопленные логи в S3
// --------------------------------------
func (l *Logger) FlushBuffer(logs []string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// копируем данные из буфера в новый массив
	quotedLogs := make([]string, len(logs))
	copy(quotedLogs, logs)

	// Объединение всех строк с разделителем новой строки
	finalStr := strings.Join(quotedLogs, "")

	// получаем старые данные из s3 storage
	old_data_logs, err := l.readDataFromS3()
	if err != nil {
		l.logger.Error("failed to read old data logs", zap.Error(err))
		return
	}

	// если есть старые данные, то добавляем их в конец
	if len(old_data_logs) > 0 {
		finalStr = string(old_data_logs) + "\n" + finalStr
	}

	// удаляем пробелы в начале и конце
	finalStr = strings.TrimSpace(finalStr)

	// записываем данные в s3 storage
	if err := l.writeLogsToS3(finalStr); err != nil {
		l.logger.Error("failed to upload log batch",
			zap.String("bucket", bucketName),
			zap.String("file", fileName),
			zap.Error(err),
		)

		return
	}

	l.logger.Info("successfully flushed buffer .log",
		zap.String("bucket", bucketName),
		zap.String("file", fileName),
		zap.Int("size", len(finalStr)),
	)
}

// LoggerMiddleware добавляет данные в log (s3 storage)
// ----------------------------------------------------
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

		if err := l.AddLogBuffer(entry); err != nil {
			// логирование ошибки при добавлении данных
			l.logger.Warn("buffer is full", zap.String("status", "sending to write logs"))
		}
	})
}

// Logging добавляет данные в log (s3 storage)
// ------------------------------------------
func (l *Logger) AddLogBuffer(log string) error {
	buffer = append(buffer, log)

	// Записываем лог в файл
	// Для продакшен или разработки
	l.productionLogging(log)

	if len(buffer) >= bufferSize {
		l.FlushBuffer(buffer)
		buffer = buffer[:0]
		return fmt.Errorf("buffer is full")
	}

	return nil
}

// getHTTPVersion добавляет версию http в строку
// ---------------------------------------------
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

// createS3Client создает клиент для взаимодействия с s3 storage
// -------------------------------------------------------------
func createS3Client() (*s3.Client, error) {
	// устанавливаем кредиты или данные api user
	creds := credentials.NewStaticCredentialsProvider(env.Envs.ACCESS_KEY, env.Envs.SECRET_KEY, "")

	// настройка конфига для взаимодействия с s3 storage
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("ru-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           "https://s3.ru-1.storage.selcloud.ru",
				SigningRegion: "ru-1",
			}, nil
		})),
		config.WithCredentialsProvider(creds),
	)

	// обработка ошибки
	if err != nil {
		return nil, fmt.Errorf("error create s3 client")
	}

	// возвращяем client
	return s3.NewFromConfig(cfg), nil
}

// writeLogsToS3 записывает лог в S3
// ---------------------------------
func (l *Logger) writeLogsToS3(logs string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := l.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   strings.NewReader(logs),
	})

	return err
}

// readDataFromS3 читает данные из s3 storage
// ------------------------------------------
func (l *Logger) readDataFromS3() ([]byte, error) {
	// создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// получение объекта из s3 storage
	resp, err := l.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return []byte{}, nil
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
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
