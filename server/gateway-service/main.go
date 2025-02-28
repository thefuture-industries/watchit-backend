package main

import (
	"bytes"
	"fmt"
	"go-gateway-service/internal/config"
	"go-gateway-service/internal/networking"
	"go-gateway-service/internal/security"
	"io/ioutil"
	"net/http"
)

func main() {
	config := config.NewConfig()

	http.HandleFunc("/api/v1/user/", func(w http.ResponseWriter, r *http.Request) {
		newPath := "/micro/user" + r.URL.Path[len("/api/v1/user"):]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewReader(body))

		// Проверка на наличие запрещённых слов в запросе
		if security.IsRequestSuspicious(string(body)) == 1 {
			http.Error(w, "Suspicious content detected in request", http.StatusForbidden)
			return
		}

		if security.IsRequestSuspicious(string(r.URL.Path)) == 1 || security.IsRequestSuspicious(string(r.URL.RawQuery)) == 1 {
			http.Error(w, "Suspicious content detected in request", http.StatusForbidden)
			return
		}

		// Прокси запрос на микросервис
		targetURL := fmt.Sprintf("%s://%s%s%s", config.Security, config.UserAddr, config.UserPort, newPath)
		networking.ForwardToBackend(w, r, targetURL)
	})

	// Запуск сервера
	fmt.Printf("the multithreaded server listens on the port %s...\n", config.ServerPort)
	if err := http.ListenAndServe(config.ServerPort, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
