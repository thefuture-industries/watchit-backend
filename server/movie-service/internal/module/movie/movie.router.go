package movie

import (
	"go-movie-service/internal/common/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) RegisterRoutes(router *mux.Router) {
	movieRouter := router.PathPrefix("/movie").Subrouter()

	movieRouter.Use(middleware.AuthMiddleware)

	movieRouter.HandleFunc("", h.MovieGetHandler).Methods(http.MethodGet)
}
