package movie

import (
	"go-movie-service/internal/common/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) MovieRoutes(router *mux.Router) {
	movieRouter := router.PathPrefix("/movie").Subrouter()

	movieRouter.Use(middleware.AuthMiddleware)

	movieRouter.HandleFunc("/{id}", h.MovieDetailsHandler).Methods(http.MethodGet)
}
