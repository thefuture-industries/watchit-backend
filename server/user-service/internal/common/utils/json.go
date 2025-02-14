package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ---------------------------------------
// Проверка и декодирование данных от user
// ---------------------------------------
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

// ---------------------------
// Функция ответа пользователю
// ---------------------------
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Security-Policy", "script-src 'self';")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(map[string]any{"message": v})
}

// --------------------------------
// Функция обработки ошибок сервера
// --------------------------------
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, err.Error())
}
