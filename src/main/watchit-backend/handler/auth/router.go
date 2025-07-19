package auth

import (
	"net/http"
	"watchit/httpx/pkg/httpx"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/create", httpx.ErrorHandler(h.CreateHandler)).Methods(http.MethodPost)
	authRouter.HandleFunc("/out", httpx.ErrorHandler(h.OutHandler)).Methods(http.MethodPost)
}
