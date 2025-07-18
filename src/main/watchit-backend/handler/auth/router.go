package auth

import (
	"github.com/gorilla/mux"
	"net/http"
	"watchit/httpx/pkg/httpx"
)

func (h *Handler) RegisterRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/create", httpx.ErrorHandler(h.CreateHandler)).Methods(http.MethodGet)
	authRouter.HandleFunc("/out", httpx.ErrorHandler(h.OutHandler)).Methods(http.MethodPost)
}
