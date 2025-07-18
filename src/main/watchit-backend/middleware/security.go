package middleware

import (
	"io"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	maxBodySize        = 1 << 20
	maxViolationsPerIP = 3
	ipBlockDuration    = 10 * time.Minute
)

var (
	blockedIPs   = make(map[string]time.Time)
	ipViolations = make(map[string]int)
	mutex        sync.Mutex
)

var sqlInjectionPattern = regexp.MustCompile(`(?i)\b(SELECT|UNION|INSERT|UPDATE|DELETE|DROP|OR\s+1=1)\b`)
var xssPattern = regexp.MustCompile(`(?i)(<script.*?>.*?</script>|<.*?on\w+=['"].*?['"]|javascript:)`)

func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)

		if isBlocked(ip) {
			log.Printf("Blocked IP %s", ip)
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}

		// limit size
		r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

		var bodyContent string
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Request body too large or unreadable", http.StatusRequestEntityTooLarge)
				registerViolation(ip)
				return
			}

			bodyContent = string(data)
			r.Body = io.NopCloser(strings.NewReader(bodyContent))
		}

		// check query params
		for _, values := range r.URL.Query() {
			for _, val := range values {
				if containsSQLInjection(val) || containsXSS(val) {
					http.Error(w, "Malicious query parameter", http.StatusBadRequest)
					registerViolation(ip)
					return
				}
			}
		}

		// check body
		if bodyContent != "" && (containsSQLInjection(bodyContent) || containsXSS(bodyContent)) {
			http.Error(w, "Malicious content in body", http.StatusBadRequest)
			registerViolation(ip)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func containsSQLInjection(input string) bool {
	return sqlInjectionPattern.MatchString(input)
}

func containsXSS(input string) bool {
	return xssPattern.MatchString(input)
}

func getIP(r *http.Request) string {
	ip := r.RemoteAddr
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		ip = strings.Split(forwarded, ",")[0]
	}

	host, _, _ := net.SplitHostPort(ip)
	return host
}

func isBlocked(ip string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	expiry, exists := blockedIPs[ip]
	if !exists {
		return false
	}

	if time.Now().After(expiry) {
		delete(blockedIPs, ip)
		delete(ipViolations, ip)

		return false
	}

	return true
}

func registerViolation(ip string) {
	mutex.Lock()
	defer mutex.Unlock()

	ipViolations[ip]++
	if ipViolations[ip] >= maxViolationsPerIP {
		blockedIPs[ip] = time.Now().Add(ipBlockDuration)
	}
}
