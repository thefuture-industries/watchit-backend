package movie

import (
	"go-movie-service/internal/common/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) RegisterRoutes(router *mux.Router) {
	movieRouter := router.PathPrefix("/movie").Subrouter()

	movieRouter.Use(middleware.AuthMiddleware)

	movieRouter.HandleFunc("/{id}", h.MovieDetailsHandler).Methods(http.MethodGet)
	movieRouter.HandleFunc("", h.MovieGetHandler).Methods(http.MethodGet)
	movieRouter.HandleFunc("/t/{uuid}", h.MovieTextHandler).Methods(http.MethodPost)
	movieRouter.HandleFunc("/image/{image}", h.MovieImageHandler).Methods(http.MethodGet)

	movieFREERouter := router.PathPrefix("/movie/f").Subrouter()
	movieFREERouter.HandleFunc("/t/gost", h.MovieTextFREEHandler).Methods(http.MethodPost)
}
