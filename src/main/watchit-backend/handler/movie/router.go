package movie

import (
	"github.com/gorilla/mux"
	"net/http"
	"watchit/httpx/middleware"
	"watchit/httpx/pkg/httpx"
)

func (h *Handler) RegisterRoutes(router *mux.Router) {
	movieRouter := router.PathPrefix("/movies").Subrouter()

	movieRouter.Use(middleware.IsAuthenticated)

	movieRouter.HandleFunc("/suggest", httpx.ErrorHandler(h.GetMoviesSuggestHandler)).Methods(http.MethodPost)
	movieRouter.HandleFunc("/search", httpx.ErrorHandler(h.GetMoviesBySearchHandler)).Methods(http.MethodPost)
	movieRouter.HandleFunc("/{id}", httpx.ErrorHandler(h.GetDetailsMovieHandler)).Methods(http.MethodGet)
}
