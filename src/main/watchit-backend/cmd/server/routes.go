package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"watchit/httpx/handler"
	"watchit/httpx/handler/auth"
	"watchit/httpx/handler/favourite"
	"watchit/httpx/handler/movie"
	"watchit/httpx/middleware"
)

func (s *httpServer) routes() http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.NewLogger(s.logger).LoggerMiddleware)
	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.SecurityMiddleware)

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	baseHandler := &handler.BaseHandler{
		Db:     s.db,
		Logger: s.logger,
		Store:  s.store,
	}

	// routes path
	auth.NewHandler(baseHandler).RegisterRoutes(subrouter)
	movie.NewHandler(baseHandler).RegisterRoutes(subrouter)
	favourite.NewHandler(baseHandler).RegisterRoutes(subrouter)

	// docs
	s.docs(subrouter)

	return s.cors(router)
}
