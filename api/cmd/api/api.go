package api

import (
	"database/sql"
	"flicksfi/internal/app/admin"
	"flicksfi/internal/app/apis"
	"flicksfi/internal/app/favourite"
	"flicksfi/internal/app/limiter"
	"flicksfi/internal/app/recommendation"
	"flicksfi/internal/app/user"
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
	limiterService := limiter.NewService(s.db)
	userService := user.NewService(s.db)
	userHandler := user.NewHandler(userService, limiterService)
	userHandler.RegisterRoutes(subrouter)

	// ----------------------------
	// определение recommendations router
	// ----------------------------
	recommendationService := recommendation.NewService(s.db)
	recommendationHandler := recommendation.NewHandler(recommendationService, userService)
	recommendationHandler.RegisterRoutes(subrouter)

	// ----------------------------
	// определение favourite router
	// ----------------------------
	favouriteService := favourite.NewService(s.db)
	favouriteHandler := favourite.NewHandler(favouriteService, userService, recommendationService)
	favouriteHandler.RegisterRoutes(subrouter)

	// ----------------------------
	// определение favourite router
	// ----------------------------
	limiterService = limiter.NewService(s.db)
	limiterHandler := limiter.NewHandler(limiterService, userService)
	limiterHandler.RegisterRoutes(subrouter)

	// -----------------------
	// определение apis router
	// -----------------------
	apisService := apis.NewService(s.db)
	apisHandler := apis.NewHandler(apisService, userService, limiterService)
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
