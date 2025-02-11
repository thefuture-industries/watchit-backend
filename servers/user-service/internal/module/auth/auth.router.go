package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/signin", SigninHandler).Methods(http.MethodPost)
}
