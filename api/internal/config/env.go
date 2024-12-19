package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// database
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string

	// jwt
	JWTExpirationInSeconds int64
	JWTSecret              string

	// hashing
	SUPER_SECRET_KEY string
	IV               string

	// S3 storage
	ACCESS_KEY string
	SECRET_KEY string

	// apis
	YOUTUBE_KEY_API    string
	TMDB_KEY_API       string
	GIGA_CHAT_ID       string
	GIGA_CHAT_SECRET   string
	GIGA_CHAT_AUTH_KEY string
}

var Envs = initConfig()

// ---------------------
// Функция для переменных
// ---------------------
func initConfig() Config {
	godotenv.Load()

	// Данные config
	return Config{
		// database
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DBU_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "flick_finder"),

		// jwt
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "_logify_-secret-_token_-!2024!-envs."),

		// hashing
		SUPER_SECRET_KEY: getEnv("SUPER_SECRET_KEY", "abc&1*~#^2^#s0^=)^^7%b34"),
		IV:               getEnv("IV", "123456789012"),

		// S3 storage
		ACCESS_KEY: getEnv("ACCESS_KEY", "1c4700bdc1b24df4a432afc62f350800"),
		SECRET_KEY: getEnv("SECRET_KEY", "99f4d1b4ccf04975941e138e8b4e21ee"),

		// apis
		YOUTUBE_KEY_API:    getEnv("YOUTUBE_KEY_API", "AIzaSyDBDGaVTs3rUgYtKXeBkaQY6veyqWp8PKg"),
		TMDB_KEY_API:       getEnv("TMDB_KEY_API", "ecfe8540ac63325e0c50686c0be8848d"),
		GIGA_CHAT_ID:       getEnv("GIGA_CHAT_ID", "aaeaef98-1937-4790-b7e0-fad35de06a9b"),
		GIGA_CHAT_SECRET:   getEnv("GIGA_CHAT_PERSONAL", "ab15a8cc-8903-4fb3-8ffb-b2f7dc158fd8"),
		GIGA_CHAT_AUTH_KEY: getEnv("GIGA_CHAT_AUTH_KEY", "YWFlYWVmOTgtMTkzNy00NzkwLWI3ZTAtZmFkMzVkZTA2YTliOmFiMTVhOGNjLTg5MDMtNGZiMy04ZmZiLWIyZjdkYzE1OGZkOA=="),
	}
}

// Читаем Config
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

// Читаем Config для INT
func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
