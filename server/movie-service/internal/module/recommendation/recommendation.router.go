package recommendation

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) RegisterRoutes(router *mux.Router) {
	recommendationRouter := router.PathPrefix("/recommendation").Subrouter()

	recommendationRouter.Use(middleware)

	recommendationRouter.HandleFunc("/{uuid}", h.RecommendationGetHandler).Methods(http.MethodGet)
	recommendationRouter.HandleFunc("", h.RecommendationAddHandler).Methods(http.MethodPost)
}
