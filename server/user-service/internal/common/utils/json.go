package utils

import (
	"compress/gzip"
	"fmt"
	"go-user-service/internal/lib"
	"io"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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
func WriteJSON(w http.ResponseWriter, r *http.Request, status int, v any) {
	logger := lib.NewLogger()

	accept := r.Header.Get("Accept-Encoding")
	shouldGzip := strings.Contains(accept, "gzip")

	var writer io.Writer = w
	var gzipWriter *gzip.Writer

	if shouldGzip {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Vary", "Accept-Encoding")

		gzipWriter = gzip.NewWriter(w)
		defer gzipWriter.Close()

		writer = gzipWriter
	}

	w.Header().Set("Content-Security-Policy", "script-src 'self';")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(writer).Encode(map[string]any{"message": v}); err != nil {
		logger.Error(err.Error())
		http.Error(w, "Unknown error from the server", http.StatusBadGateway)
	}
}

func CacheJSON(w http.ResponseWriter, limit int) {
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", limit))
}
