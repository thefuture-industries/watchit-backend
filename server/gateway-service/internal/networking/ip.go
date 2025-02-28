package networking

import (
	"net"
	"net/http"
	"strings"
)

func GetClientIP(r *http.Request) string {
	// Сначала проверяем X-Forwarded-For (если сервер за прокси)
	// Этот заголовок содержит список IP-адресов, через которые прошёл запрос
	// Первый IP — это реальный IP клиента
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// IP-адрес может быть в формате "X.X.X.X, Y.Y.Y.Y", нам нужен первый
		parts := strings.Split(xForwardedFor, ",")
		return strings.TrimSpace(parts[0])
	}

	// Если нет прокси, то используем RemoteAddr
	// Формат: "IP:PORT"
	ip := r.RemoteAddr
	// Убираем порт
	host, _, err := net.SplitHostPort(ip)
	if err != nil {
		return ip // если не удалось разделить, возвращаем всё
	}

	if host == "::1" {
		return "127.0.0.1"
	}

	return host
}
