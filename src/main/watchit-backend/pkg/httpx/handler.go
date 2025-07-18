package httpx

import (
	"net/http"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

// ErrorHandler обёртка для HTTP-хендлеров, которая позволяет
// обрабатывать ошибки централизованно
func ErrorHandler(fn AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			HttpResponseError(w, r, err)
		}
	}
}
