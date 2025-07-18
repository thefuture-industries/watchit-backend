package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	pkgstr "watchit/httpx/pkg/strings"
)

type Logger struct {
	logDir string
	dev    bool
}

var (
	timeYkb *time.Location
)

func NewLogger() *Logger {
	now := time.Now()
	day := fmt.Sprintf("%02d", now.Day())
	month := fmt.Sprintf("%02d", now.Month())

	var errLoad error
	timeYkb, errLoad = time.LoadLocation("Asia/Yekaterinburg")
	if errLoad != nil {
		panic(errLoad.Error())
	}

	logPath := os.Getenv("LOG_DIR")
	logDir := filepath.Join(logPath, fmt.Sprintf("%s.%s", day, month))

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Println("Error creating log directory:", err)
	}

	go_env := os.Getenv("GO_ENV") == "DEV"

	return &Logger{logDir: logDir, dev: go_env}
}

func (l *Logger) writeDateTime(log *strings.Builder, t time.Time) {
	log.WriteString(l.twoDigit(t.Day()))
	log.WriteByte('.')
	log.WriteString(l.twoDigit(int(t.Month())))
	log.WriteByte('.')
	log.WriteString(strconv.Itoa(t.Year()))
	log.WriteByte(' ')
	log.WriteString(l.twoDigit(t.Hour()))
	log.WriteByte(':')
	log.WriteString(l.twoDigit(t.Minute()))
	log.WriteByte(':')
	log.WriteString(l.twoDigit(t.Second()))
}

func (l *Logger) twoDigit(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}

	return strconv.Itoa(n)
}

func (l *Logger) logFile(fileName string, message string) {
	filePath := filepath.Join(l.logDir, fileName)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(message + "\n"); err != nil {
		fmt.Println("Error writing log to file:", err)
	}
}

func (l *Logger) Info(format string, args ...any) {
	now := time.Now().In(timeYkb)

	var log strings.Builder

	log.WriteString("[")
	l.writeDateTime(&log, now)
	log.WriteString("] [INFO] ")
	pkgstr.Formatted(&log, format, args...)

	logMessage := log.String()

	// if this develops version
	if l.dev {
		println(logMessage)
	}

	l.logFile("server.log", logMessage)
}

func (l *Logger) Warning(format string, args ...any) {
	now := time.Now().In(timeYkb)

	var log strings.Builder

	log.WriteString("[")
	l.writeDateTime(&log, now)
	log.WriteString("] [WARN] ")
	pkgstr.Formatted(&log, format, args...)

	logMessage := log.String()

	// if this develops version
	if l.dev {
		println(logMessage)
	}

	l.logFile("warning.log", logMessage)
}

func (l *Logger) Session(format string, args ...any) {
	now := time.Now().In(timeYkb)

	var log strings.Builder

	log.WriteString("[")
	l.writeDateTime(&log, now)
	log.WriteString("] [SESS] ")
	pkgstr.Formatted(&log, format, args...)

	logMessage := log.String()

	// if this develops version
	if l.dev {
		println(logMessage)
	}

	l.logFile("session.log", logMessage)
}

func (l *Logger) Error(format string, args ...any) {
	now := time.Now().In(timeYkb)

	var log strings.Builder

	log.WriteString("[")
	l.writeDateTime(&log, now)
	log.WriteString("] [ERROR] ")
	pkgstr.Formatted(&log, format, args...)

	logMessage := log.String()

	println(logMessage)
	l.logFile("errors.log", logMessage)
}
