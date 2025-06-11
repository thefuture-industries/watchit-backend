package favourite

import (
	"go-movie-service/internal/common/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) RegisterRoutes(router *mux.Router) {
	favouriteRouter := router.PathPrefix("/favourite").Subrouter()

	favouriteRouter.Use(middleware.AuthMiddleware)

	favouriteRouter.HandleFunc("", h.GetFavouritesHandler).Methods(http.MethodGet)
	favouriteRouter.HandleFunc("", h.AddFavouriteHandler).Methods(http.MethodPost)
}
