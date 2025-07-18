package movie

import (
	"watchit/httpx/handler"
	"watchit/httpx/pkg/machinelearning"
)

type Handler struct {
	*handler.BaseHandler
}

func NewHandler(base *handler.BaseHandler) *Handler {
	return &Handler{BaseHandler: base}
}

var lsaBuilder *machinelearning.LSABuilder = machinelearning.NewLSABuilder()
