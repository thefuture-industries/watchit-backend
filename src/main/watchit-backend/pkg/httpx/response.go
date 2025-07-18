package httpx

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
	"watchit/httpx/infra/logger"
	"watchit/httpx/pkg/httpx/httperr"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func HttpParse(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("отсутствует текст запроса")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func HttpResponse(w http.ResponseWriter, r *http.Request, status int, v any) {
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(writer).Encode(map[string]any{"message": v}); err != nil {
		http.Error(w, "неизвестная ошибка от сервера", http.StatusBadGateway)
	}
}

func HttpResponseError(w http.ResponseWriter, r *http.Request, err error) {
	log := logger.NewLogger()
	log.Error("%v", err)

	if httpErr, ok := err.(httperr.HTTPError); ok {
		HttpResponse(w, r, httpErr.StatusCode(), httpErr.Error())
		return
	}

	HttpResponse(w, r, http.StatusInternalServerError, err.Error())
}

func HttpCache(w http.ResponseWriter, limit int) {
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", limit))
}
