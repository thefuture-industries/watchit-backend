package recommendation

import (
	"go-movie-service/internal/common/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) RegisterRoutes(router *mux.Router) {
	recommendationRouter := router.PathPrefix("/recommendation").Subrouter()

	recommendationRouter.Use(middleware.AuthMiddleware)

	recommendationRouter.HandleFunc("/{uuid}", h.RecommendationGetHandler).Methods(http.MethodGet)
	recommendationRouter.HandleFunc("", h.RecommendationAddHandler).Methods(http.MethodPost)
}
