package main

import (
	"github.com/gorilla/handlers"
	"net/http"
)

func (s *httpServer) cors(handler http.Handler) http.Handler {
	origins := handlers.AllowedOrigins([]string{"http://localhost:5173", "http://0.0.0.0:80", "https://0.0.0.0:443"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"})

	allowCredentials := handlers.AllowCredentials()

	return handlers.CORS(origins, methods, headers, allowCredentials)(handler)
}
