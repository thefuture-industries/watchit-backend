package api

import (
	"database/sql"
	"flick_finder/internal/app/admin"
	"flick_finder/internal/app/apis"
	"flick_finder/internal/app/user"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	// ---------------------------------
	// Создания router и префикс /api/v1
	// ---------------------------------
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// -----------------------
	// определение user router
	// -----------------------
	userService := user.NewService(s.db)
	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(subrouter)

	// -----------------------
	// определение apis router
	// -----------------------
	apisService := apis.NewService(s.db)
	apisHandler := apis.NewHandler(apisService, userService)
	apisHandler.RegisterRoutes(subrouter)

	// -----------------------
	// определение admin router
	// -----------------------
	adminService := admin.NewService(s.db)
	adminHandler := admin.NewHandler(adminService, userService)
	adminHandler.RegisterRoutes(subrouter)

	// ---------------
	// подключаем CORS
	// ---------------
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"})

	// --------------------
	// Возращяем http ответ
	// --------------------
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, handlers.CORS(origins, methods, headers)(router))
}
