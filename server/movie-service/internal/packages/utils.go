package packages

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Security-Policy", "script-src 'self';")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func CacheJSON(w http.ResponseWriter, seconds int, status int, v any) error {
	w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(seconds))
	w.Header().Set("Expires", time.Now().Add(time.Duration(seconds)*time.Second).Format(http.TimeFormat))

	return WriteJSON(w, status, v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	if jsonErr := WriteJSON(w, status, map[string]string{"error": err.Error()}); jsonErr != nil {
		log.Printf("Failed to write error JSON: %v", jsonErr)
	}
}
