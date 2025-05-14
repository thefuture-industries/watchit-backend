package packages_test

import (
	"encoding/json"
	"go-movie-service/internal/packages"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCacheJSON(t *testing.T) {
	// Тестовые данные
	testData := map[string]string{
		"message": "Test message",
		"status":  "success",
	}

	// Тест кейсы
	tests := []struct {
		name            string
		cacheSeconds    int
		status          int
		wantStatusCode  int
		wantCacheHeader string
	}{
		{
			name:            "Cache for 3600 seconds",
			cacheSeconds:    3600,
			status:          http.StatusOK,
			wantStatusCode:  http.StatusOK,
			wantCacheHeader: "public, max-age=3600",
		},
		{
			name:            "Cache for 60 seconds",
			cacheSeconds:    60,
			status:          http.StatusOK,
			wantStatusCode:  http.StatusOK,
			wantCacheHeader: "public, max-age=60",
		},
		{
			name:            "Cache with non-200 status",
			cacheSeconds:    3600,
			status:          http.StatusCreated,
			wantStatusCode:  http.StatusCreated,
			wantCacheHeader: "public, max-age=3600",
		},
	}

	// Запускаем тесты
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем ResponseRecorder для записи ответа
			rr := httptest.NewRecorder()

			// Вызываем функцию CacheJSON
			err := packages.CacheJSON(rr, tt.cacheSeconds, tt.status, testData)
			if err != nil {
				t.Errorf("CacheJSON() error = %v", err)
				return
			}

			// Проверяем код статуса
			if rr.Code != tt.wantStatusCode {
				t.Errorf("CacheJSON() returned wrong status code: got %v want %v",
					rr.Code, tt.wantStatusCode)
			}

			// Проверяем заголовок Cache-Control
			cacheControl := rr.Header().Get("Cache-Control")
			if cacheControl != tt.wantCacheHeader {
				t.Errorf("CacheJSON() returned wrong Cache-Control header: got %v want %v",
					cacheControl, tt.wantCacheHeader)
			}

			// Проверяем заголовок Expires
			expires := rr.Header().Get("Expires")
			if expires == "" {
				t.Errorf("CacheJSON() did not set Expires header")
			}

			// Проверяем, что заголовок Expires установлен на правильное время в будущем
			expiresTime, err := time.Parse(http.TimeFormat, expires)
			if err != nil {
				t.Errorf("CacheJSON() set invalid Expires header: %v", err)
			}

			// Проверяем только, что заголовок Expires установлен и находится в будущем
			now := time.Now()
			if expiresTime.Before(now) {
				t.Errorf("CacheJSON() set Expires time in the past: got %v, current time %v",
					expiresTime, now)
			}

			// Проверяем тип контента
			contentType := rr.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("CacheJSON() set wrong Content-Type: got %v want %v",
					contentType, "application/json")
			}

			// Проверяем тело ответа
			var responseData map[string]string
			err = json.NewDecoder(rr.Body).Decode(&responseData)
			if err != nil {
				t.Errorf("Error decoding response body: %v", err)
			}

			// Проверяем, что данные в ответе совпадают с тестовыми данными
			for key, expectedValue := range testData {
				actualValue, exists := responseData[key]
				if !exists {
					t.Errorf("Response missing expected key: %s", key)
				}
				if actualValue != expectedValue {
					t.Errorf("Response has wrong value for key %s: got %v want %v",
						key, actualValue, expectedValue)
				}
			}
		})
	}
}

func TestWriteError(t *testing.T) {
	// Тест кейсы
	tests := []struct {
		name           string
		status         int
		err            error
		wantStatusCode int
		wantErrorMsg   string
	}{
		{
			name:           "Bad Request Error",
			status:         http.StatusBadRequest,
			err:            &customError{"Bad request error"},
			wantStatusCode: http.StatusBadRequest,
			wantErrorMsg:   "Bad request error",
		},
		{
			name:           "Internal Server Error",
			status:         http.StatusInternalServerError,
			err:            &customError{"Internal server error"},
			wantStatusCode: http.StatusInternalServerError,
			wantErrorMsg:   "Internal server error",
		},
	}

	// Запускаем тесты
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем ResponseRecorder для записи ответа
			rr := httptest.NewRecorder()

			// Вызываем функцию WriteError
			packages.WriteError(rr, tt.status, tt.err)

			// Проверяем код статуса
			if rr.Code != tt.wantStatusCode {
				t.Errorf("WriteError() returned wrong status code: got %v want %v",
					rr.Code, tt.wantStatusCode)
			}

			// Проверяем тип контента
			contentType := rr.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("WriteError() set wrong Content-Type: got %v want %v",
					contentType, "application/json")
			}

			// Проверяем тело ответа
			expected := `{"error":"` + tt.wantErrorMsg + `"}`
			if !strings.Contains(rr.Body.String(), tt.wantErrorMsg) {
				t.Errorf("WriteError() returned wrong body: got %v want to contain %v",
					rr.Body.String(), expected)
			}
		})
	}
}

// Вспомогательный тип для тестирования
type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}
