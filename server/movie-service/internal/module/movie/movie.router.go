package movie

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) MovieRoutes(router *mux.Router) {
	movieRouter := router.PathPrefix("/movie").Subrouter()

	movieRouter.HandleFunc("/{id}", h.MovieDetailsHandler).Methods(http.MethodGet)
}
