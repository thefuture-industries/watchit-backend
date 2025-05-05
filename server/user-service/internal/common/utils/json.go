// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package utils

import (
	"encoding/json"
	"fmt"
	"go-user-service/internal/packages"
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
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Security-Policy", "script-src 'self';")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(map[string]any{"message": v}); err != nil {
		packages.ErrorLog(err)
		http.Error(w, "Неизвестная ошибка от сервера", http.StatusBadGateway)
	}
}

// --------------------------------
// Функция обработки ошибок сервера
// --------------------------------
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, err.Error())
}
