package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	env "flick_finder/internal/config"

	"github.com/go-playground/validator"
	"github.com/juju/ratelimit"
)

// ------------------------------
// ------------------------------
// Переменная для валидации данных
// ------------------------------
var Validate = validator.New()

var key string = env.Envs.SUPER_SECRET_KEY
var iv string = env.Envs.IV

// ------------------------------
// Проверка и декодирование данных от user
// ------------------------------
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

// ------------------------
// ------------------------
// Функция ответа пользователю
// ------------------------
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Security-Policy", "script-src 'self';")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

// ------------------------
// ------------------------
// Функция обработки ошибок
// ------------------------
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

// ------------------------
// ------------------------
// Функция избежания DDos
// ------------------------
func DDosPropperty() *ratelimit.Bucket {
	return ratelimit.NewBucket(10, int64(time.Second))
}

// -----------------------
// -----------------------
// Шифрование (encrypt)
// -----------------------
func Encrypt(plaintext string) (string, error) {
	// Создание блока шифрования AES
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("error creating the encryption block")
	}

	// Создание блочного режима шифрования CBC
	gsm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creating block encryption mode")
	}

	// Шифрование данных
	ciphertext := gsm.Seal(nil, []byte(iv), []byte(plaintext), nil)

	// Кодирование зашифрованных данных в base64 для удобства
	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
	return encodedCiphertext, nil
}

// -----------------------
// -----------------------
// Де-шифрование (decrypt)
// -----------------------
func Decrypt(encodedCiphertext string) (string, error) {
	// Декодирование зашифрованных данных из base64
	ciphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return "", fmt.Errorf("error decoding encrypted data")
	}

	// Создание блока шифрования AES
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("error creating the encryption block")
	}

	// Создание блочного режима шифрования CBC
	gsm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creating block encryption mode")
	}

	// Расшифровка данных
	plaintext, err := gsm.Open(nil, []byte(iv), ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("data decryption error")
	}

	return string(plaintext), nil
}
