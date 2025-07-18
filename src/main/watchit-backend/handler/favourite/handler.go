package favourite

import "watchit/httpx/handler"

type Handler struct {
	*handler.BaseHandler
}

func NewHandler(base *handler.BaseHandler) *Handler {
	return &Handler{BaseHandler: base}
}
