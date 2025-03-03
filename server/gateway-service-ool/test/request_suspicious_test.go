package test

import (
	"go-gateway-service/internal/security"
	"testing"
)

func TestIsRequestSuspicious(t *testing.T) {
	tests := []struct {
		name           string
		request        string
		expectedResult int
	}{
		{
			name:           "Suspicious URL with script tag",
			request:        "http://localhost/api/v1/user/sync?<script>alert('XSS');</script>",
			expectedResult: 1, // Запрещённое слово в URL
		},
		{
			name:           "Suspicious URL with script in path",
			request:        "http://localhost/api/v1/user/sync/<script>",
			expectedResult: 1, // Запрещённое слово в пути URL
		},
		{
			name:           "Suspicious query parameter with script",
			request:        "http://localhost/api/v1/user/sync?param=<script>",
			expectedResult: 1, // Запрещённое слово в параметре запроса
		},
		{
			name:           "Suspicious body with script tag",
			request:        "POST /api/v1/user/create HTTP/1.1\r\n\r\n<script>alert('XSS');</script>",
			expectedResult: 1, // Запрещённое слово в теле запроса
		},
		{
			name:           "Clean request",
			request:        "http://localhost/api/v1/user/sync",
			expectedResult: 0, // Запрос без запрещённых слов
		},
		{
			name:           "Clean POST request",
			request:        "POST /api/v1/user/create HTTP/1.1\r\n\r\nvalid data here",
			expectedResult: 0, // Запрос без запрещённых слов
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Проверяем запрос с помощью функции IsRequestSuspicious
			result := security.IsRequestSuspicious(tt.request)

			// Сравниваем результат с ожидаемым значением
			if result != tt.expectedResult {
				t.Errorf("expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
