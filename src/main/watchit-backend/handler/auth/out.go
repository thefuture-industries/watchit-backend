package auth

import (
	"net/http"
	"watchit/httpx/pkg/httpx"
)

func (h *Handler) OutHandler(w http.ResponseWriter, r *http.Request) error {
	httpx.HttpResponse(w, r, http.StatusOK, "OUT")
	return nil
}
