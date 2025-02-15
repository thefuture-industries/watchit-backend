// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package packages

import (
	"context"
	"go-user-service/cmd/conf"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Vision struct {
	logger  *zap.Logger
	monitor *conf.Vision
}

func NewVision(monitor *conf.Vision, logger *zap.Logger) *Vision {
	return &Vision{
		monitor: monitor,
		logger:  logger,
	}
}

type ContextKey string

const (
	// DBQueryContextKey ключ для монитора БД в контексте
	DBQueryContextKey ContextKey = "dbQueryMonitor"
	// ErrorContextKey ключ для канала ошибок в контексте
	ErrorContextKey ContextKey = "errorChan"
)

func (v *Vision) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter{
			ResponseWriter: w,
			status:         200,
		}

		errorChan := make(chan error, 1)
		defer close(errorChan)

		ctx := context.WithValue(r.Context(), ErrorContextKey, errorChan)
		r = r.WithContext(ctx)

		next.ServeHTTP(wrapped, r)

		var err error
		select {
		case err = <-errorChan:
			// Ошибка получена
			v.monitor.VisionError(err)
			v.logger.Error("request error",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Error(err),
			)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError) // Отправляем клиенту ошибку 500
		default:
		}

		duration := time.Since(start)

		v.monitor.VisionRequest(duration)
	})
}
