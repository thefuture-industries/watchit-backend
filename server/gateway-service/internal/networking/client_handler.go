package networking

import (
	"bytes"
	"fmt"
	"go-gateway-service/internal/utils"
	"io/ioutil"
	"log"
	"net/http"
)

func ForwardToBackend(w http.ResponseWriter, r *http.Request, targetURL string) {
	log.Println("Request: " + GetClientIP(r))
	log.Println("We forward the request to the microservice: " + targetURL)
	utils.LogRequestToFile(GetClientIP(r), r.Method, r.URL.Path, "200 (Forwarded)", fmt.Sprintf("We forward the request to the microservice: %s", targetURL))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	r.Body = ioutil.NopCloser(bytes.NewReader(body))

	// Создаём новый запрос к микросервису
	req, err := http.NewRequest(r.Method, targetURL, bytes.NewReader(body))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header = r.Header

	// Отправляем запрос на микросервис
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request to microservice", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ от микросервиса
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(response)
}
