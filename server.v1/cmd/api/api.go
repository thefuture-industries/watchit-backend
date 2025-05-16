package api

import (
	"context"
	"database/sql"
	"flicksfi/cmd/configuration"
	"flicksfi/internal/app/admin"
	"flicksfi/internal/app/apis"
	"flicksfi/internal/app/favourite"
	"flicksfi/internal/app/limiter"
	"flicksfi/internal/app/recommendation"
	"flicksfi/internal/app/user"
	logging "flicksfi/packages/logger"
	"log"
	"net/http"

	_ "flicksfi/docs"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type APIServer struct {
	addr   string
	db     *sql.DB
	server *http.Server
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

	// Swagger UI
	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))
	router.HandleFunc("/adm/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.html")
	})

	// ---------------
	// Создание logger
	// ---------------
	track := configuration.NewTrack()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// ------------------
	// Логирование server
	// ------------------
	router.Use(logging.NewLogger(logger).LoggerMiddleware)

	// --------------
	// Services route
	// --------------
	limiterService := limiter.NewService(s.db, logger, track)
	userService := user.NewService(s.db, logger, track)
	recommendationService := recommendation.NewService(s.db, logger, track)
	favouriteService := favourite.NewService(s.db, logger, track)
	apisService := apis.NewService(s.db)
	adminService := admin.NewService(s.db, track)

	// -----------------------
	// определение user router
	// -----------------------
	userHandler := user.NewHandler(userService, limiterService)
	userHandler.RegisterRoutes(subrouter)

	// ----------------------------
	// определение recommendations router
	// ----------------------------
	recommendationHandler := recommendation.NewHandler(recommendationService, userService)
	recommendationHandler.RegisterRoutes(subrouter)

	// ----------------------------
	// определение favourite router
	// ----------------------------
	favouriteHandler := favourite.NewHandler(favouriteService, userService, recommendationService)
	favouriteHandler.RegisterRoutes(subrouter)

	// ----------------------------
	// определение favourite router
	// ----------------------------
	limiterHandler := limiter.NewHandler(limiterService, userService)
	limiterHandler.RegisterRoutes(subrouter)

	// -----------------------
	// определение apis router
	// -----------------------
	apisHandler := apis.NewHandler(apisService, userService, limiterService)
	apisHandler.RegisterRoutes(subrouter)

	// ------------------------
	// определение admin router
	// ------------------------
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

func (s *APIServer) Shutdown(ctx context.Context) error {
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}

	return nil
}
