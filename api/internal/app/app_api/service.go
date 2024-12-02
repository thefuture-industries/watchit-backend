package app_api

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"time"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

// Генерация api_key
func generateAPIKEY() string {
	// Создаем массив байтов нужной длины
	b := make([]byte, 10)

	// Генерируем случайные байты
	if _, err := rand.Read(b); err != nil {
		log.Fatal("error create api_key")
	}

	return base64.StdEncoding.EncodeToString(b)
}

// Создание api_key
func (s *Service) CreateAPIKEY() string {
	api_key := fmt.Sprintf("%s-%s-%s", generateAPIKEY(), generateAPIKEY(), generateAPIKEY())
	return api_key
}

// Запись в БД api_key
func (s *Service) InsertAPIKEY(api_key string, uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, "insert into api_keys (uuid, api_key, createdAt) values(?, ?, ?)", uuid, api_key, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return fmt.Errorf("database insert error")
	}

	return nil
}
