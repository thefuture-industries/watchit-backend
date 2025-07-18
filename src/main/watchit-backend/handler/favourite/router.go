package favourite

import (
	"github.com/gorilla/mux"
	"net/http"
	"watchit/httpx/middleware"
	"watchit/httpx/pkg/httpx"
)

func (h *Handler) RegisterRoutes(router *mux.Router) {
	favouriteRouter := router.PathPrefix("/favourites").Subrouter()

	favouriteRouter.Use(middleware.IsAuthenticated)

	favouriteRouter.HandleFunc("/", httpx.ErrorHandler(h.GetFavouritesHandler)).Methods(http.MethodGet)
	favouriteRouter.HandleFunc("/", httpx.ErrorHandler(h.AddFavouriteHandler)).Methods(http.MethodPost)
}
