// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package packages

import (
	"fmt"
	"go-user-service/internal/common/utils"
	"net/http"
	"time"

	"github.com/noneandundefined/vision-go"
	"go.uber.org/zap"
)

type Vision struct {
	logger  *zap.Logger
	monitor *vision.Vision
}

func NewVision(monitor *vision.Vision, logger *zap.Logger) *Vision {
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

		next.ServeHTTP(wrapped, r)

		err := utils.GetErrorFromContext(r.Context())
		fmt.Println(err)
		if err != nil {
			fmt.Println("ERROR!!!")
			v.monitor.VisionDBError()
			v.monitor.VisionError(err)
			v.logger.Error("[REQ ERROR]", zap.Error(err))
		}

		// var err error
		// select {
		// case err = <-errorChan:
		// 	// Ошибка получена
		// 	v.monitor.VisionError(err)
		// 	v.logger.Error("request error",
		// 		zap.String("method", r.Method),
		// 		zap.String("path", r.URL.Path),
		// 		zap.Error(err),
		// 	)
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError) // Отправляем клиенту ошибку 500
		// default:
		// }

		duration := time.Since(start)

		v.monitor.VisionRequest(duration)
	})
}
