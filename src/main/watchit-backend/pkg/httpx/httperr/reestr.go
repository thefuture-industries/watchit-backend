package httperr

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/lib/pq"
)

func New(msg string, code int) *ServerError {
	return &ServerError{
		Msg:  msg,
		Code: code,
	}
}

func InternalServerError(msg string) *ServerError {
	return New(msg, http.StatusInternalServerError)
}

func ServiceUnavailable(msg string) *ServerError {
	return New(msg, http.StatusServiceUnavailable)
}

func Forbidden(msg string) *ServerError {
	return New(msg, http.StatusForbidden)
}

func BadRequest(msg string) *ServerError {
	return New(msg, http.StatusBadRequest)
}

func RequestTimeout(msg string) *ServerError {
	return New(msg, http.StatusRequestTimeout)
}

func NotFound(msg string) *ServerError {
	return New(msg, http.StatusNotFound)
}

func Conflict(msg string) *ServerError {
	return New(msg, http.StatusConflict)
}

func Unauthorized(msg string) *ServerError {
	return New(msg, http.StatusUnauthorized)
}

func Db(ctx context.Context, err error) *ServerError {
	if err == nil {
		return nil
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return RequestTimeout(Err_ContextDeadlineExceeded.Error())
	}

	if errors.Is(err, context.Canceled) {
		return Conflict(Err_ContextCanceled.Error())
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return RequestTimeout(Err_DbTimeout.Error())
		}

		return InternalServerError(Err_DbNetwork.Error())
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505": // error unnique
			return Conflict(Err_UniqueViolation.Error())
		case "23503": // not found
			return NotFound(Err_UserNotFound.Error())
		}
	}

	switch {
	case errors.Is(err, Err_NotDeleted):
		return Conflict(Err_NotDeleted.Error())

	case errors.Is(err, Err_NotUpdated):
		return Conflict(Err_NotUpdated.Error())

	}

	return InternalServerError("не удалось выполнить операцию с базой данных")
}
