package logger

import (
	"bytes"
	"context"
	env "flicksfi/internal/config"
	"fmt"
	"io"
	"net/http"
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
	}
}

const (
	bucketName = "flicksfi-logs"
	fileName   = "flicksfi-monitoring.log"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int64
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

		if err := l.Logging(entry); err != nil {
			l.logger.Error("failed to write logs",
				zap.String("bucket", bucketName),
				zap.String("file", fileName),
				zap.Error(err),
			)
		}

		// логирование успешного добавления данных
		l.logger.Info("successfully appended log",
			zap.String("bucket", bucketName),
			zap.String("file", fileName),
			zap.Int("size", len(entry)),
		)
	})
}

// Logging добавляет данные в log (s3 storage)
// ------------------------------------------
func (l *Logger) Logging(entry string) error {
	// Получаем текущие логи
	existing_logs, err := l.readDataFromS3()
	if err != nil {
		return err
	}

	// Добавляем новую запись
	updated_logs := append(existing_logs, []byte(entry)...)

	// Запись данных в s3 storage
	return l.writeLogsToS3(updated_logs)
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
func (l *Logger) writeLogsToS3(logs []byte) error {
	_, err := l.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(logs),
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
		err = l.createNewLogFile()
		if err != nil {
			return nil, fmt.Errorf("failed to create new log file: %w", err)
		}
		return []byte{}, nil
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// createNewLogFile создает новый файл в s3 storage
// ------------------------------------------------
func (l *Logger) createNewLogFile() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Создаем начальное содержимое файла
	initialContent := []byte(fmt.Sprintf("# Log file created at %s\n",
		time.Now().Format("02/Jan/2006:15:04:05 -0700")))

	// Загружаем пустой файл в S3
	_, err := l.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(initialContent),
	})
	if err != nil {
		return fmt.Errorf("failed to create initial log file: %w", err)
	}

	return nil
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
