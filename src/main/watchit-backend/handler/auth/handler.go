package auth

import (
	jsoniter "github.com/json-iterator/go"
	"watchit/httpx/handler"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Handler struct {
	*handler.BaseHandler
}

func NewHandler(base *handler.BaseHandler) *Handler {
	return &Handler{BaseHandler: base}
}
