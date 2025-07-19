package movie

import (
	"net/http"
	"watchit/httpx/middleware"
	"watchit/httpx/pkg/httpx"

	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(router *mux.Router) {
	movieRouter := router.PathPrefix("/movies").Subrouter()

	movieRouter.Use(middleware.IsAuthenticated)

	movieRouter.HandleFunc("/suggest", httpx.ErrorHandler(h.GetMoviesSuggestHandler)).Methods(http.MethodPost)
	movieRouter.HandleFunc("/search", httpx.ErrorHandler(h.GetMoviesBySearchHandler)).Methods(http.MethodPost)
	movieRouter.HandleFunc("/image/{image}", httpx.ErrorHandler(h.GetMovieImageHandler)).Methods(http.MethodGet)
	movieRouter.HandleFunc("/details/{id}", httpx.ErrorHandler(h.GetDetailsMovieHandler)).Methods(http.MethodGet)
}
